/*
 * Copyright (c) 2022.
 *
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * All rights reserved
 *
 * This program contains research , trial - errors. So this program can't guarantee your system will work as intended.
 */

package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"notans/backend/common"
	"notans/backend/security"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}

type UserAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func GetAllUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var users []User
		db.Find(&users)
		common.RespondJSON(context, http.StatusOK, users, "SUCCESS_FETCHED")
	}
}

func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var users *User
		id := context.Param("id")
		users = FindUser(&id, &id)
		common.RespondJSON(context, http.StatusOK, users, "SUCCESS_FETCHED")
	}

}

func NewUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user User
		err := json.NewDecoder(context.Request.Body).Decode(&user)
		if err != nil {
			common.LogPrintln("Service::User:JsonDecoder", err.Error())
			return
		}
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if err != nil {
			common.LogPrintln("Service::User:GeneratedPassword", err.Error())
			return
		}

		user.Password = string(password)

		tx := db.Create(&user)
		if tx.Error != nil {
			common.RespondJSON(context, http.StatusBadRequest, nil, "USER_EXIST")
			return
		}
		user.Password = ""
		common.RespondJSON(context, http.StatusCreated, user, "SUCCESS_FETCHED")
	}

}

func DeleteUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		user := FindUser(&id, nil)
		if user == nil {
			common.RespondJSON(context, http.StatusNotFound, nil, "NOT_FOUND")
			return
		}
		db.Where("id = ?", id).Delete(&User{})
		common.RespondJSON(context, http.StatusCreated, nil, "SUCCESS_DELETED")
	}

}

func FindUser(id *string, username *string) *User {
	var user User
	var tx *gorm.DB

	tx = db.Where("id = ?", &id).Or("Username = ?", &username).First(&user)

	if tx.Error != nil {
		common.LogPrintln("SERVICE::USER::FIND", tx.Error.Error())
		return nil
	}

	return &user
}

func SignIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user UserAuthRequest
		err := json.NewDecoder(context.Request.Body).Decode(&user)
		if err != nil {
			common.LogPrintln("Service::User:JsonDecoder", err.Error())
			return
		}

		usr := FindUser(nil, &user.Username)
		if usr == nil {
			common.RespondJSON(context, http.StatusNotFound, nil, "NOT_FOUND")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(user.Password))
		if err != nil {
			common.LogPrintln("SERVICE::USER::SIGNIN", err.Error())
			common.RespondJSON(context, http.StatusUnauthorized, nil, "USERNAME_PASSWORD_INVALID")
			return
		}

		jwt := security.Jwt{
			Config: config,
		}

		token := jwt.Create(usr.Username)

		common.RespondJSON(context, http.StatusOK, UserAuthResponse{
			Username: usr.Username,
			Token:    token,
		}, "SUCCESS_FETCHED")

	}

}
