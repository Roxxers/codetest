package server

import (
	"github.com/gin-gonic/gin"
	"thirdlight.com/watcher-node/lib"
)

// SetupRouter creates the http server and defines all routes with methods.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/files", func(c *gin.Context) {
		c.JSON(200, map[string]string{"hello": "world"})
	})

	r.POST("/hello", func(c *gin.Context) {
		var introduction lib.HelloOperation
		c.BindJSON(&introduction)
		c.JSON(200, introduction)
	})

	return r
}
