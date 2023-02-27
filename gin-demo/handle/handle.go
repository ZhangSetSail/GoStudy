package handle

import (
	"context"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/client-go"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/client-go/watch"
	"k8s.io/client-go/kubernetes"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/typed/apis/v1beta1"
)

func InitHandle(clientSet *kubernetes.Clientset, gatewayClient *gateway.GatewayV1beta1Client) {
	//初始化k8s客户端
	mw := watch.CreateResourceWatch(clientSet)
	mw.Start()
	defaultManagerClientGo = client_go.CreateClientGoManager(mw, clientSet, gatewayClient, context.Background())
}

var defaultManagerClientGo *client_go.ManagerClientGo

func GetManagerClientGo() *client_go.ManagerClientGo {
	return defaultManagerClientGo
}
