package watch

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/model"
	appsv1 "k8s.io/api/apps/v1"
	"sync"
)

type ManagerWatch interface {
	Start()
	OnAdd(obj interface{})
	OnDelete(obj interface{})
	GetApp(kind string) *App
	OnUpdate(oldObj interface{}, newObj interface{})
	Ready() bool
}

type managerWatch struct {
	informers *Informer
	listers   *Lister
	stopch    chan struct{}
	app       sync.Map
}

func (mw *managerWatch) Start() {
	stopch := make(chan struct{})
	mw.informers.Start(stopch)
	mw.stopch = stopch
	for !mw.Ready() {
	}
}

func (mw *managerWatch) getApp(kind string) *App {
	var app *App
	app = mw.GetApp(kind)
	if app == nil {
		app = InitCacheApp()
		mw.app.Store(kind, app)
	}
	return app
}

func (mw *managerWatch) GetApp(kind string) *App {
	app, ok := mw.app.Load(kind)
	if ok {
		app := app.(*App)
		return app
	}
	return nil
}

func (mw *managerWatch) OnAdd(obj interface{}) {
	if deployment, ok := obj.(*appsv1.Deployment); ok {
		app := mw.getApp(model.Deployment)
		app.SetDeployment(deployment)
		return
	}
}

func (mw *managerWatch) OnDelete(obj interface{}) {

}

func (mw *managerWatch) OnUpdate(oldObj interface{}, newObj interface{}) {

}

func (mw *managerWatch) Ready() bool {
	return mw.informers.Ready()
}
