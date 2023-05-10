package main

import (
	"context"
	"fmt"
	"github.com/openkruise/kruise-api/apps/v1alpha1"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	v1beta1
	v1alpha1.DeploymentTemplateSpec{}
	rollout := &v1alpha1.Rollout{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-rollout",
		},
		Spec: v1alpha1.RolloutSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "my-app",
				},
			},
			Template: v1alpha1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "my-app",
					},
				},
				Spec: v1alpha1.PodSpec{
					Containers: []v1alpha1.Container{
						{
							Name:  "my-container",
							Image: "nginx",
						},
					},
				},
			},
			Canary: &v1alpha1.CanaryStrategy{
				Steps: []v1alpha1.CanaryStep{
					{
						SetWeight: 10,
						Pause: &v1alpha1.Duration{
							Duration: metav1.Duration{
								Seconds: 10,
							},
						},
					},
					{
						SetWeight: 50,
						Pause: &v1alpha1.Duration{
							Duration: metav1.Duration{
								Seconds: 20,
							},
						},
					},
					{
						SetWeight: 100,
					},
				},
			},
		},
	}

	rolloutClient := clientset.AppsV1alpha1().Rollouts("default")
	result, err := rolloutClient.Create(context.Background(), rollout, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Created Rollout %q.\n", result.GetObjectMeta().GetName())
}

func int32Ptr(i int32) *int32 { return &i }
