package server

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	endpoints "thirdlight.com/aggregation-server/lib"
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

// TODO: Maybe add more to this?????
func getFiles(c *gin.Context) {
	c.JSON(200, nodes.FetchAllFiles())
}

// registerNode is the endpoint for registering a watcher node with this server
func registerNode(c *gin.Context) {
	var intro lib.HelloOperation
	c.BindJSON(&intro)

	// Error here means we don't have instance stored, so add it to the list
	if _, err := nodes.Find(intro.Instance); err != nil {
		watcherAddr := parseRemoteAddr(c.Request.RemoteAddr)
		if err := nodes.New(intro.Instance, watcherAddr, intro.Port); err != nil {
			c.Status(500)
			return
		}
		// No error, therefore added successfully.
		c.Status(200)
		return
	}
	// 200 here as there is nothing to do but not returning 200 creates a debug output in watcher node
	// Really should be using something like 204
	c.Status(200)
	log.Debugf("Seen already registered node: %s", intro.Instance)
	return
}

// deregisterNode is the endpoint for removing a watcher node, usually on shutdown of that node, from the server
func deregisterNode(c *gin.Context) {
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
}

// SetupRouter creates the http server and defines all routes with methods.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET(endpoints.FilesEndpoint, getFiles)
	r.POST(endpoints.HelloEndpoint, registerNode)
	r.POST(endpoints.ByeEndpoint, deregisterNode)

	return r
}
