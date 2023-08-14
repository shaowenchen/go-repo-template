package server

import (
	"github.com/gin-gonic/gin"
)

func SetupV1Router(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("", Get)
		v1.GET("/version", Version)
	}
}

func SetupV1AuthRouter(r *gin.Engine) {
	v1auth := r.Group("api/v1").Use(AuthMiddleware())
	{
		v1auth.GET("/auth", GetAuth)
	}
}
