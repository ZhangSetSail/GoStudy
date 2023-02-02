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

type KeyValue struct {
	Key   string `form:"key"`
	Value string `form:"value"`
}

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

func initDapr() DaprTest {
	var err error
	var daprCli DaprTest
	daprCli.client, err = dapr.NewClient()
	daprCli.ctx = context.Background()
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
