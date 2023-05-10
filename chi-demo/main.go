package main

import (
	kruiseclientset "github.com/openkruise/kruise-api/client/clientset/versioned"
	"github.com/openkruise/kruise-api/rollouts/v1alpha1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Fatal(err)
	}
	kruiseClient := kruiseclientset.NewForConfigOrDie(config)
	kruiseClient.RolloutsV1alpha1().Rollouts().Create()
	v1alpha1.Rollout{}
}
