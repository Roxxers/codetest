package server

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"thirdlight.com/aggregation-server/watcher"
	"thirdlight.com/watcher-node/lib"
)

var nodes = watcher.CreateNodesList()

// parseRemoteAddr removes random port info, adds correct port info, and returns address to watcher node
func parseRemoteAddr(addr string) string {
	parts := strings.Split(addr, ":")
	if len(parts) > 1 {
		parts = parts[:len(parts)-1]
	}
	fixedStr := strings.Join(parts, "")
	fixedAddr := fmt.Sprintf("%s", fixedStr)
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

		// Error here means we don't have instance stored, so add it to the list
		if _, err := nodes.Find(intro.Instance); err != nil {
			watcherAddr := parseRemoteAddr(c.Request.RemoteAddr)
			node, err := nodes.New(intro.Instance, watcherAddr, intro.Port)
			if err != nil {
				log.Error(err)
			}
			log.Infof("Created new node: %s", node.Instance)
			c.Status(200)
			return
		}
		log.Debugf("Seen already registered node: %s", intro.Instance)
		c.Status(200)
		return
	})

	r.POST("/bye", func(c *gin.Context) {
		var bye lib.ByeOperation
		c.BindJSON(&bye)
		if err := nodes.Remove(bye.Instance); err != nil {
			// Error means node doesn't exist
			log.Debugln(err)
			c.Status(404)
			return
		}
		log.Infof("Successfully removed node from registered nodes: %s", bye.Instance)
		c.Status(200)
		return
	})

	return r
}
