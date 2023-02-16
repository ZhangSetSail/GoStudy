package client_go

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/controller/client-go"
	"github.com/gin-gonic/gin"
)

func PodRouter(r *gin.RouterGroup) {
	r.GET("/status", client_go.GetClientSetManage().GetPodDetails)
	r.GET("/name", client_go.GetClientSetManage().GetPodsName)
}

func PodsRouter(r *gin.RouterGroup) {
	r.GET("/status", client_go.GetClientSetManage().GetPodDetails)
	r.GET("/name", client_go.GetClientSetManage().GetPodsName)
}
