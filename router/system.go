package router

import (
	v1 "gin-tmp/api/v1"

	"github.com/gin-gonic/gin"
)

func SysRouter(router *gin.RouterGroup) {
	route := router.Group("")
	route.GET("/ping", v1.Pong)
}
