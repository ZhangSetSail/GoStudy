package main

import (
	"context"
	"google.golang.org/protobuf/types/known/structpb"
	"istio.io/api/meta/v1alpha1"
	"istio.io/api/networking/v1alpha3"
	istio "istio.io/client-go/pkg/apis/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

func main() {
	// 加载Kubeconfig文件
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}
	istioClient, err := istioClient.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Istio clientset: %v", err)
	}
	envoy := istio.EnvoyFilter{
		TypeMeta: metav1.TypeMeta{
			Kind:       "EnvoyFilter",
			APIVersion: "networking.istio.io/v1alpha3",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        "",
			Labels:      nil,
			Annotations: nil,
		},
		Spec: v1alpha3.EnvoyFilter{
			WorkloadSelector: nil,
			ConfigPatches: []*v1alpha3.EnvoyFilter_EnvoyConfigObjectPatch{{
				ApplyTo: v1alpha3.EnvoyFilter_ApplyTo(),
				Match: &v1alpha3.EnvoyFilter_EnvoyConfigObjectMatch{
					Context:     0,
					Proxy:       nil,
					ObjectTypes: &v1alpha3.EnvoyFilter_EnvoyConfigObjectMatch_Listener{},
				},
				Patch: &v1alpha3.EnvoyFilter_Patch{
					Value: &structpb.Struct{},
				},
			}},
			Priority: 0,
		},
		Status: v1alpha1.IstioStatus{},
	}
	istioClient.NetworkingV1alpha3().EnvoyFilters("zhangqh").Create(context.Background(), &envoy, metav1.CreateOptions{})
}
