package handle

import (
	"context"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/client-go"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/client-go/watch"
	"k8s.io/client-go/kubernetes"
)

func InitHandle(clientSet *kubernetes.Clientset) {
	//初始化k8s客户端
	mw := watch.CreateResourceWatch(clientSet)
	mw.Start()
	defaultManagerClientGo = client_go.CreateClientGoManager(mw, clientSet, context.Background())
}

var defaultManagerClientGo *client_go.ManagerClientGo

func GetManagerClientGo() *client_go.ManagerClientGo {
	return defaultManagerClientGo
}
