package watch

import (
	"github.com/gin-gonic/gin"
)

func WatchRouter(r *gin.RouterGroup) {
	PodsWatchRouter(r.Group("/pods"))
}
