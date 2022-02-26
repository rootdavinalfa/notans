/*
 * Copyright (c) 2021-2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package lang

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type LangParams struct {
	LangID string
	Key    string
}

type Model struct {
	LangID      string `json:"lang-id"`
	Description string `json:"description"`
	Data        []Data `json:"data"`
}

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func readJson(langID string) (*Model, error) {
	path := "./resources/lang/" + langID + ".json"
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	lang := Model{}
	err = json.Unmarshal(jsonFile, &lang)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &lang, err
}

func GetLangValue(param LangParams) string {

	var lang, err = readJson(param.LangID)
	if err != nil {
		log.Println("Failed to parse language!")
		return err.Error()
	}

	if lang.LangID != param.LangID {
		log.Println("Returned language not match")
		return "n/a"
	}
	println(param.LangID)
	for _, value := range lang.Data {
		if value.Key == param.Key {
			return value.Value
		}
	}
	return "n/a"
}
