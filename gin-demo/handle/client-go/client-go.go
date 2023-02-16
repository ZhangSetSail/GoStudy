package client_go

import (
	"context"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/client-go/watch"
	"k8s.io/client-go/kubernetes"
)

type ManagerClientGo struct {
	clientSet *kubernetes.Clientset
	ctx       context.Context
	mw        watch.ManagerWatch
}

func CreateClientGoManager(mw watch.ManagerWatch, clientSet *kubernetes.Clientset, ctx context.Context) *ManagerClientGo {
	return &ManagerClientGo{
		clientSet: clientSet,
		ctx:       ctx,
		mw:        mw,
	}
}
