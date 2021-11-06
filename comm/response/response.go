package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一处理响应信息

const (
	SUCCESS int64 = 10011
	ERROR   int64 = 10010
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    SUCCESS,
		"message": "ok",
	})
}
