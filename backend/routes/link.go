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

func LinkRoute(route *gin.RouterGroup) {
	common.Get(route, "/link/list", service.GetLinks())
	common.Get(route, "/link/get/:id", service.GetLink())
	common.Get(route, "/link/s-get/:slinkParam", service.GetLink())
	common.Post(route, "/link", service.CreateNewLink())
	common.Delete(route, "/link/:id", service.DeleteLink())
}
