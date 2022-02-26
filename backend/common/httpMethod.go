/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package common

import (
	"github.com/gin-gonic/gin"
)

// Wrap the router for GET method
func Get(router *gin.RouterGroup, path string, handlerFunc gin.HandlerFunc) {
	router.GET(path, handlerFunc)
}

// Wrap the router for POST method
func Post(router *gin.RouterGroup, path string, handlerFunc gin.HandlerFunc) {
	router.POST(path, handlerFunc)
}

// Wrap the router for PUT method
func Put(router *gin.RouterGroup, path string, handlerFunc gin.HandlerFunc) {
	router.PUT(path, handlerFunc)
}

// Wrap the router for DELETE method
func Delete(router *gin.RouterGroup, path string, handlerFunc gin.HandlerFunc) {
	router.DELETE(path, handlerFunc)
}
