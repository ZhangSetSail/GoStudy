package router

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func PodRouter(p *gin.Engine) {
	podsGroup := p.Group("/pods")
	podsGroup.GET("/get", controller.GetClientSetManage().GetPod)
	podsGroup.GET("/list", controller.GetClientSetManage().GetPods)
}
