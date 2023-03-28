package client_go

import (
	client_go "github.com/ZhangSetSail/GoStudy/gin-demo/controller/client-go"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/:kind/name-list", client_go.GetClientSetManage().GetResourcesName)
	r.POST("/gateway")
	WatchRouter(r.Group("/watch"))
}
