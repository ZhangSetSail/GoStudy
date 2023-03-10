package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DaprTest struct {
	client dapr.Client
	ctx    context.Context
	store  string
}

//KeyValue 解析请求数据所用到的结构体
type KeyValue struct {
	Key   string `form:"key"`
	Value string `form:"value"`
}

//Add 添加
func (d DaprTest) Add(c *gin.Context) {
	var kv KeyValue
	if err := c.ShouldBindQuery(&kv); err != nil {
		c.Error(err)
	}
	data := []byte(kv.Value)
	if err := d.client.SaveState(d.ctx, d.store, kv.Key, data, nil); err != nil {
		c.Error(err)
	}
	c.String(http.StatusOK, "Successfully added")
}

//Get 获取
func (d DaprTest) Get(c *gin.Context) {
	var kv KeyValue
	if err := c.ShouldBindQuery(&kv); err != nil {
		c.Error(err)
	}
	item, err := d.client.GetState(d.ctx, d.store, kv.Key, nil)
	if err != nil {
		c.Error(err)
	}
	ret := fmt.Sprintf("value is %v", string(item.Value))
	c.String(http.StatusOK, ret)
}

//Delete 删除
func (d DaprTest) Delete(c *gin.Context) {
	var kv KeyValue
	if err := c.ShouldBindQuery(&kv); err != nil {
		c.Error(err)
	}
	if err := d.client.DeleteState(d.ctx, d.store, kv.Key, nil); err != nil {
		c.Error(err)
	}
	c.String(http.StatusOK, "Successfully delete")
}

//initDapr 初始化 Dapr
func initDapr() DaprTest {
	var err error
	var daprCli DaprTest
	//初始化 dapr 客户端
	daprCli.client, err = dapr.NewClient()
	daprCli.ctx = context.Background()
	//所用到的存储名称，也就是我们上面添加的redis.yaml中的name字段
	daprCli.store = "statestore"
	if err != nil {
		panic(err)
	}
	return daprCli
}

func main() {
	r := gin.Default()
	daprCli := initDapr()
	defer daprCli.client.Close()
	r.POST("/add", daprCli.Add)
	r.GET("/get", daprCli.Get)
	r.DELETE("/delete", daprCli.Delete)
	r.Run()
}
