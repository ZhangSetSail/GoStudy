package client_go

import (
	"fmt"
	"github.com/ZhangSetSail/GoStudy/gin-demo/model"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (m *ManagerClientGo) GetResourcesNameByNamespace(kind, namespace string) (*model.ResourceNamesResponse, error) {
	var names []string
	switch strings.ToLower(kind) {
	case "deployment":
		deploys, err := m.clientSet.AppsV1().Deployments(namespace).List(m.ctx, v1.ListOptions{})
		if err != nil {
			logrus.Errorf("get deployment list failure by namespace: %v", namespace)
			return nil, err
		}
		for _, deploy := range deploys.Items {
			names = append(names, deploy.GetName())
		}
	case "service":
		fmt.Println()
	case "statefulSet":
		fmt.Println()
	case "pod":
		pods, err := m.clientSet.CoreV1().Pods(namespace).List(m.ctx, v1.ListOptions{})
		if err != nil {
			logrus.Errorf("get pod list failure by namespace: %v", namespace)
			return nil, err
		}
		for _, pod := range pods.Items {
			names = append(names, pod.GetName())
		}
	case "configMap":
		fmt.Println()
	case "":
		fmt.Println()
	}
	return &model.ResourceNamesResponse{
		Names:     names,
		Namespace: namespace,
		Kind:      kind,
	}, nil
}
