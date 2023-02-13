package server

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/controller"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle"
	"github.com/ZhangSetSail/GoStudy/gin-demo/router"
	"github.com/sirupsen/logrus"
)

//Run 主体调用执行
func Run() error {
	//设置日志输出格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//初始化k8s客户端
	clientSet := InitK8SClient()
	controller.CreateClientSetManage()
	handle.InitHandle(clientSet)
	return router.InitRouter()
}