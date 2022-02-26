package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"notans/backend"
	"notans/backend/common"
	config2 "notans/backend/config"
	"strconv"
)

func main() {
	common.Banner()
	config := config2.GetConfig()
	app := backend.IApp{}
	app.Initialize(config)
	fmt.Println("Notans - Version Alpha")

	app.Router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Router.Use(gin.Recovery())

	port := strconv.Itoa(config.Port)
	fmt.Println("Listened on port: " + port)
	log.Fatalln(app.Router.Run(":" + port))
}
