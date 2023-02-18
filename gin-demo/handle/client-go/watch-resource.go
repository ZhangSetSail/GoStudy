package client_go

import (
	"github.com/sirupsen/logrus"
)

func (m *ManagerClientGo) GetWatchResourcesByNamespace(kind, namespace string) ([]string, error) {
	logrus.Infof("get watch pod namespace by %v", namespace)
	var podList []string
	app := m.mw.GetApp(namespace)
	pods := app.GetPod()
	for _, pod := range pods {
		podList = append(podList, pod.GetName())
	}
	return podList, nil
}
