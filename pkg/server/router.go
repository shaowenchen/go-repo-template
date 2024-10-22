package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shaowenchen/go-repo-template/swagger"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupV1Router(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("", Get)
		v1.GET("/version", Version)
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

func SetupV1AuthRouter(r *gin.Engine) {
	v1auth := r.Group("api/v1").Use(AuthMiddleware())
	{
		v1auth.GET("/auth", GetAuth)
	}
}
