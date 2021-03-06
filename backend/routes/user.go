/*
 * Copyright (c) 2022.
 *
 * Davin Alfarizky Putra Basudewa <dbasudewa@gmail.com>
 * All rights reserved
 *
 * This program contains research , trial - errors. So this program can't guarantee your system will work as intended.
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"notans/backend/common"
	"notans/backend/service"
)

func UserRoute(route *gin.RouterGroup) {
	common.Get(route, "/user/list", service.GetAllUser())
	common.Get(route, "/user/get/:id", service.GetUser())
	common.Post(route, "/user", service.NewUser())
	common.Delete(route, "/user/:id", service.DeleteUser())
}
