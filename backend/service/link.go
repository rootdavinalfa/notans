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
	"gorm.io/gorm"
	"net/http"
	"notans/backend/common"
	"strconv"
)

type Link struct {
	gorm.Model
	OLink    string
	SLink    string `gorm:"uniqueIndex"`
	Enabled  string `gorm:"type:varchar(1) default '1'"`
	Username string `json:"username"`
	User     User   `json:"-" gorm:"foreignKey:username;references:username"`
}

func CreateNewLink() gin.HandlerFunc {
	return func(context *gin.Context) {
		var link Link
		err := json.NewDecoder(context.Request.Body).Decode(&link)
		if err != nil {
			common.LogPrintln("Service::Link:JsonDecoder", err.Error())
			return
		}

		if common.IsStringEmpty(link.OLink) {
			common.RespondJSON(context, http.StatusNotAcceptable, nil, "OLINK is empty")
			return
		}

		link.SLink = common.RandAlphaNumeric(8)
		user, isExist := context.Get("user")
		if isExist {
			link.Username = user.(*User).Username
		}

		link.Enabled = "1"

		tx := db.Create(&link)
		if tx.Error != nil {
			common.RespondJSON(context, http.StatusBadRequest, nil, "Link exist")
			return
		}

		common.RespondJSON(context, http.StatusCreated, link, "SUCCESS_FETCHED")
	}
}

func GetLinks() gin.HandlerFunc {
	return func(context *gin.Context) {
		var links []Link
		db.Find(&links)
		common.RespondJSON(context, http.StatusOK, links, "SUCCESS_FETCHED")
	}
}

func GetLink() gin.HandlerFunc {
	return func(context *gin.Context) {
		var link *Link
		id, _ := strconv.Atoi(context.Param("id"))
		sLinkParam := context.Param("slinkParam")
		oLink := context.Query("oLink")
		sLink := context.Query("sLink")
		if common.IsStringEmpty(sLink) {
			sLink = sLinkParam
		}
		link = findLink(&id, &oLink, &sLink)
		if link == nil {
			common.RespondJSON(context, http.StatusNotFound, nil, "NOT_FOUND")
			return
		}
		common.RespondJSON(context, http.StatusOK, link, "SUCCESS_FETCHED")
	}
}

func RedirectLink() gin.HandlerFunc {
	return func(context *gin.Context) {
		var link *Link
		sLinkParam := context.Param("slinkParam")
		link = findLink(nil, nil, &sLinkParam)
		if link == nil {
			common.RespondJSON(context, http.StatusNotFound, nil, "NOT_FOUND")
			return
		}
		context.Redirect(http.StatusPermanentRedirect, link.OLink)
	}
}

func findLink(id *int, oLink *string, sLink *string) *Link {
	var link Link
	tx := db.Where("id = ?", &id).Or("o_link = ?", &oLink).Or("s_link = ?", &sLink).First(&link)
	if tx.Error != nil {
		common.LogPrintln("SERVICE::USER::FIND", tx.Error.Error())
		return nil
	}

	return &link
}

func DeleteLink() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		link := findLink(&id, nil, nil)
		if link == nil {
			common.RespondJSON(context, http.StatusNotFound, nil, "NOT_FOUND")
			return
		}
		db.Where("id = ?", id).Delete(&link)
		common.RespondJSON(context, http.StatusCreated, nil, "SUCCESS_DELETED")
	}
}
