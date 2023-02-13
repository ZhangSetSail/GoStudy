package watch

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type App struct {
	service     []*corev1.Service
	statefulSet []*v1.StatefulSet
	deployment  []*v1.Deployment
	pod         []*corev1.Pod
	cm          []*corev1.ConfigMap
	pvc         []*corev1.PersistentVolumeClaim
}

func InitCacheApp() *App {
	return &App{}
}

func (a *App) SetDeployment(d *v1.Deployment) {
	if len(a.deployment) > 0 {
		for i, deploy := range a.deployment {
			if deploy.GetName() == d.GetName() {
				a.deployment[i] = d
				return
			}
		}
	}
	logrus.Infof("captures deployment created: %v", d.GetName())
	a.deployment = append(a.deployment, d)
}

func (a *App) SetStatefulSet() {

}

func (a *App) SetService() {

}

func (a *App) SetPod() {

}

func (a *App) SetConfigMap() {

}

func (a *App) SetPVC() {

}
