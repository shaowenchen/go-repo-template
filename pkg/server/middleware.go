package server

import (
	"github.com/gin-gonic/gin"
)

func Islogin(c *gin.Context) bool {
	return false
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Islogin(c) {
			c.Next()
		} else {
			println("not auth")
			c.Abort()
		}
		return
	}
}
