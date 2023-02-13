package controller

import "github.com/ZhangSetSail/GoStudy/gin-demo/api"

//ClientSetManageInterface 汇总所有 client-go 的相关的路由的接口
type ClientSetManageInterface interface {
	api.PodInterface
}

func GetClientSetManage() ClientSetManageInterface {
	return defaultClientManager
}

//CreateClientSetManage 初始化 client-go 相关路由
func CreateClientSetManage() {
	defaultClientManager = &ClientGoStruct{}
}

var defaultClientManager ClientSetManageInterface

//ClientGoStruct 汇总所有 client-go 的相关的路由的结构体
type ClientGoStruct struct {
	PodsManage
}
