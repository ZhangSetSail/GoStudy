package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
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
	clientSet, err := kubernetes.NewForConfig(config)
	gr, err := restmapper.GetAPIGroupResources(clientSet)
	if err != nil {
		fmt.Println(err)
	}
	mapper := restmapper.NewDiscoveryRESTMapper(gr)
	RESTMapping(gvk.GroupKind(), gvk.Version)
	err := buildResource.Dri.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		logrus.Errorf("delete k8s resource error %v", err)
	}
}
