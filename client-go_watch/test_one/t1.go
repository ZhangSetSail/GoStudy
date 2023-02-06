package main

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Fatal(err)
	}

	// 创建 client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	factory := informers.NewSharedInformerFactoryWithOptions(clientSet, 0, informers.WithNamespace("default"))
	informer := factory.Core().V1().Pods().Informer()
	informer.AddEventHandler(NewEventHandler())

	stopper := make(chan struct{}, 2)
	go informer.Run(stopper)
	log.Println("watch pod started...")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go test(sigs)
	<-sigs
	stopper <- struct{}{}
	close(stopper)
	log.Println("watch pod stopped...")
}
func test(sigs chan<- os.Signal) {
	time.Sleep(10 * time.Second)
	close(sigs)
}

type EventHandler struct {
}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (e *EventHandler) OnAdd(obj interface{}) {
	event := obj.(*corev1.Pod)
	log.Printf("OnAdd: %s", event.ObjectMeta.Name)
}

func (e *EventHandler) OnUpdate(oldObj, newObj interface{}) {
	event := newObj.(*corev1.Pod)
	log.Printf("OnUpdate: %s", event.ObjectMeta.Name)

}

func (e *EventHandler) OnDelete(obj interface{}) {
	event := obj.(*corev1.Pod)
	log.Printf("OnDelete: %s", event.ObjectMeta.Name)

}
