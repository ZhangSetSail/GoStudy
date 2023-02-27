package server

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/typed/apis/v1beta1"
)

func InitK8SClient() (*kubernetes.Clientset, *gateway.GatewayV1beta1Client) {
	logrus.Infof("begin init k8s client")
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Fatal(err)
	}
	// 创建 client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	gatewayClient, err := gateway.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return clientSet, gatewayClient
}
