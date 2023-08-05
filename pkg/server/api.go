package server

import (

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	ShowData(c, map[string]string{
		"GitCommit": GitCommit,
		"BuildTime": BuildTime,
	})
	return
}

func Get(c *gin.Context) {
	ShowSuccess(c, "get, but not auth")
}

func GetAuth(c *gin.Context) {
	ShowSuccess(c, "auth")
}
