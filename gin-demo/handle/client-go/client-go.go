package client_go

import (
	"context"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/client-go/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ManagerClientGo struct {
	clientSet *kubernetes.Clientset
	ctx       context.Context
	mw        watch.ManagerWatch
	config    *rest.Config
}

func CreateClientGoManager(mw watch.ManagerWatch, clientSet *kubernetes.Clientset, config *rest.Config, ctx context.Context) *ManagerClientGo {
	return &ManagerClientGo{
		clientSet: clientSet,
		ctx:       ctx,
		mw:        mw,
		config:    config,
	}
}
