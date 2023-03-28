package client_go

import (
	"context"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/client-go/watch"
	"k8s.io/client-go/kubernetes"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/typed/apis/v1beta1"
)

type ManagerClientGo struct {
	clientSet     *kubernetes.Clientset
	ctx           context.Context
	mw            watch.ManagerWatch
	gatewayClient *gateway.GatewayV1beta1Client
}

func CreateClientGoManager(mw watch.ManagerWatch, clientSet *kubernetes.Clientset, gatewayClient *gateway.GatewayV1beta1Client, ctx context.Context) *ManagerClientGo {
	return &ManagerClientGo{
		clientSet:     clientSet,
		ctx:           ctx,
		mw:            mw,
		gatewayClient: gatewayClient,
	}
}
