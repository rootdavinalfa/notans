package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"notans/backend/config"
	"notans/backend/security/middleware"
	service2 "notans/backend/service"
)

type IRoutes struct {
	Config *config.Config
	DB     *gorm.DB
}

func (rtr *IRoutes) Bind(router *gin.Engine) {
	router.POST("/auth/signin", service2.SignIn())

	service := router.Group("/service")
	middle := middleware.Middleware{
		Config: rtr.Config,
		DB:     rtr.DB,
	}

	service.Use(middle.AuthMiddle())

	UserRoute(service)
}
