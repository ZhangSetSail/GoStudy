package client_go

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/router/client-go/watch"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	PodRouter(r.Group("/pod"))
	PodsRouter(r.Group("/pods"))
	watch.WatchRouter(r.Group("/watch"))
}
