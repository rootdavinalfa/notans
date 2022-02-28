/*
 * Copyright (c) 2022.
 *
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * All rights reserved
 *
 * This program contains research , trial - errors. So this program can't guarantee your system will work as intended.
 */

package config

import (
	"encoding/json"
	"io/ioutil"
	"notans/backend/common"
)

type Config struct {
	DB              *DBConfig `json:"DB"`
	Port            int       `json:"HttpServingPort"`
	AppKey          string    `json:"JWTAppKey"`
	JwtExpireInHour int
}

type DBConfig struct {
	Dsn    string `json:"Dsn"`
	Driver string `json:"Driver"`
}

func GetConfig() *Config {

	path := "./resources/config/config.json"
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		common.LogFatal("CONFIG::CANT_READ_FILE", err.Error())
		return nil
	}

	lang := Config{}
	err = json.Unmarshal(jsonFile, &lang)
	if err != nil {
		common.LogFatal("CONFIG::PARSING_FILE", err.Error())
		return nil
	}

	return &lang
}
