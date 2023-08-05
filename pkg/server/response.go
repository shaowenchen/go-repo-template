package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowNotLogin(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "未登录",
		"code":    0,
	})
}

func ShowError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": message,
	})
}
func ShowValidatorError(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": message,
	})
}

func ShowErrorParams(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": message,
	})
}

func ShowSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
	})
}

func ShowData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}
