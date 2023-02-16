package watch

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/controller/client-go"
	"github.com/gin-gonic/gin"
)

func PodsWatchRouter(r *gin.RouterGroup) {
	r.GET("/name", client_go.GetClientSetManage().GetWatchPodsName)
}
