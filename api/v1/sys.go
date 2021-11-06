package v1

import (
	"gin-tmp/comm/response"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	response.Ok(c)
}
