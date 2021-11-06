package server

import (
	"gin-tmp/constant/global"
	"gin-tmp/router"

	"github.com/gin-gonic/gin"
)

type Serve struct {
}

func (s *Serve) RunServe() {
	e := gin.Default()
	gin.SetMode(global.GtConfig.System.Mode)
	routerGroup := e.Group("")
	{
		router.SysRouter(routerGroup)
	}
	e.Run(":" + global.GtConfig.System.Port)
}
