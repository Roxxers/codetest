package server

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"thirdlight.com/watcher-node/lib"
)

// parseRemoteAddr removes random port info, adds correct port info, and returns address to watcher node
func parseRemoteAddr(addr string, port uint) string {
	parts := strings.Split(addr, ":")
	if len(parts) > 1 {
		parts = parts[:len(parts)-1]
	}
	fixedStr := strings.Join(parts, ":")
	fixedAddr := fmt.Sprintf("%s:%d", fixedStr, port)
	return fixedAddr
}

// SetupRouter creates the http server and defines all routes with methods.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/files", func(c *gin.Context) {
		c.JSON(200, map[string]string{"hello": "world"})
	})

	r.POST("/hello", func(c *gin.Context) {
		var intro lib.HelloOperation
		c.BindJSON(&intro)

		fmt.Println(parseRemoteAddr(c.Request.RemoteAddr, intro.Port))

		c.JSON(200, intro)
	})

	return r
}
