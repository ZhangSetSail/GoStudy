package client_go

import (
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *ManagerClientGo) GetPodByNamespace(namespace string) ([]string, error) {
	logrus.Infof("get pod namespace by %v", namespace)
	var podList []string
	podsObject, err := m.clientSet.CoreV1().Pods(namespace).List(m.ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Errorf("get pod list failture: %v", err)
		return nil, err
	}
	for _, pod := range podsObject.Items {
		podList = append(podList, pod.GetName())
	}
	return podList, nil
}
