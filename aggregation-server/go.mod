module thirdlight.com/aggregation-server

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/prometheus/common v0.10.0
	github.com/sirupsen/logrus v1.4.2
	thirdlight.com/watcher-node v0.0.0-00010101000000-000000000000
)

replace thirdlight.com/watcher-node => github.com/third-light/backend-coding-challenge/watcher-node v0.0.0-20200721122029-83dd056c0276
