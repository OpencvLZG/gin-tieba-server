package util

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ResponseStatusOk(c *gin.Context, code int, msg string, result interface{}) {
	t := time.Now().Format("2006-1-2 15:04:05")
	c.JSON(200, gin.H{
		"code":   code,
		"msg":    msg,
		"result": result,
		"time":   t,
	})
}

func ResponseStatusFail(c *gin.Context, code int, msg string, result interface{}) {
	t := time.Now().Format("2006-1-2 15:04:05")
	c.JSON(200, gin.H{
		"code":   code,
		"msg":    msg,
		"result": result,
		"time":   t,
	})
}
