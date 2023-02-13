package handle

import (
	"context"
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle/watch"
	"k8s.io/client-go/kubernetes"
)

func InitHandle(clientSet *kubernetes.Clientset) {
	mw := watch.CreateResourceWatch(clientSet)
	mw.Start()
	defaultManagerPod = CreatePodManager(clientSet, context.Background())
}

var defaultManagerPod *ManagerPod

func GetManagerPod() *ManagerPod {
	return defaultManagerPod
}
