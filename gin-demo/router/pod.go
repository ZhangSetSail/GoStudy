package router

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func PodRouter(p *gin.Engine) {
	podGroup := p.Group("/pod")
	podGroup.GET("/status", controller.GetClientSetManage().GetPod)

	podsGroup := p.Group("/pods")
	podsGroup.GET("/name", controller.GetClientSetManage().GetPodsName)
}
