package server

import (
	"github.com/gin-gonic/gin"
)

func SetupV1Routes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("", Get)
		v1.GET("/version", Version)
	}
}

func SetupV1AuthRoutes(r *gin.Engine) {
	v1auth := r.Group("api/v1").Use(AuthMiddleware())
	{
		v1auth.GET("/auth", GetAuth)
	}
}
