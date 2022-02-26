package service

import (
	"gorm.io/gorm"
	"notans/backend/common"
	config2 "notans/backend/config"
)

var db *gorm.DB
var config *config2.Config

type IService struct {
	Db     *gorm.DB
	Config *config2.Config
}

func (srv *IService) InitializeService() {
	db = srv.Db
	config = srv.Config
	dbMigrate()
}

func dbMigrate() {
	err := db.AutoMigrate(&User{})
	if err != nil {
		common.LogPrintln("SERVICE::DBMIGRATE:USER", err.Error())
	}

}
