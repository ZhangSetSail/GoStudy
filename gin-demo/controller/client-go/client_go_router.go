package client_go

import "github.com/ZhangSetSail/GoStudy/gin-demo/api"

//ClientManageInterface 汇总所有 client-go 的相关的路由的接口
type ClientManageInterface interface {
	api.ResourceInterface
	api.WatchResourceInterface
}

func GetClientSetManage() ClientManageInterface {
	return defaultClientManager
}

//CreateClientSetManage 初始化 client-go 相关路由
func CreateClientSetManage() {
	defaultClientManager = &ClientGoStruct{}
}

var defaultClientManager ClientManageInterface

//ClientGoStruct 汇总所有 client-go 的相关的路由的结构体
type ClientGoStruct struct {
	ResourceManage
	WatchResourceManage
}
