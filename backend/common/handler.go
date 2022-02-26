/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package common

import (
	"github.com/gin-gonic/gin"
	"notans/backend/lang"
	response2 "notans/backend/response"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(c *gin.Context, status int, payload interface{}, keyLang string) {

	langID := c.Request.URL.Query().Get("lang")
	if len(langID) == 0 {
		langID = "en-US"
	}

	message := lang.GetLangValue(lang.LangParams{
		LangID: langID,
		Key:    keyLang,
	})

	var errs = false
	if FindfirstDigit(status) == 4 || FindfirstDigit(status) == 5 {
		errs = true
	}
	payload = response2.Response{
		Error:   errs,
		Message: message,
		Data:    payload,
	}

	c.JSON(status, payload)
}
