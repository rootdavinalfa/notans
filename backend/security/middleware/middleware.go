package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"notans/backend/common"
	config2 "notans/backend/config"
	"notans/backend/security"
	service2 "notans/backend/service"
)

type Middleware struct {
	Config *config2.Config
	DB     *gorm.DB
}

func (middle *Middleware) AuthMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := security.Jwt{Config: middle.Config}

		common.LogPrintln("Middleware::Auth", c.Request.RequestURI)
		var jwtt = c.GetHeader("Authorization")
		username := jwt.Validate(jwtt)
		if username == "" {
			common.RespondJSON(c, http.StatusUnauthorized, nil, "UNAUTHORIZED")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user := service2.FindUser(nil, &username)
		if user == nil {
			common.LogPrintln("MIDDLEWARE::AuthMiddle", "User Not Found")
			common.RespondJSON(c, http.StatusUnauthorized, nil, "UNAUTHORIZED")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
