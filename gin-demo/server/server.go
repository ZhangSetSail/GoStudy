package server

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/controller/client-go"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle"
	"github.com/ZhangSetSail/GoStudy/gin-demo/router"
	"github.com/sirupsen/logrus"
)

// Run 主体调用执行
func Run() error {
	//设置日志输出格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//初始化 client-go 相关路由
	client_go.CreateClientSetManage()
	clientSet, gatewayClient := InitK8SClient()
	//初始化实现操错
	handle.InitHandle(clientSet, gatewayClient)
	return router.InitRouter()
}
