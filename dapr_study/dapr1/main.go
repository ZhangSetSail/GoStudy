package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sync"
)

type DaprTest struct {
	client dapr.Client
	ctx    context.Context
	store  string
	pubsub string
}

// KeyValue 解析请求数据所用到的结构体
type KeyValue struct {
	Key   string `form:"key"`
	Value string `form:"value"`
}

// Add 添加
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

// Get 获取
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

// Delete 删除
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

type Body struct {
	Message string `json:"message"`
}

func (d DaprTest) A(c *gin.Context) {
	xxx, err := RequestInputs(c)
	if err != nil {
		c.Error(err)
		return
	}
	fmt.Println("A: ", xxx)
	c.String(http.StatusOK, "Successfully A")
}

func (d DaprTest) Subscribe(c *gin.Context) {
	c.JSON(http.StatusOK, []map[string]string{
		{
			"pubsubname": "pubsub",
			"topic":      "A",
			"route":      "A",
		},
	},
	)
}

// initDapr 初始化 Dapr
func initDapr() DaprTest {
	var err error
	var daprCli DaprTest
	//初始化 dapr 客户端
	daprCli.client, err = dapr.NewClient()
	daprCli.ctx = context.Background()
	//所用到的存储名称，也就是我们上面添加的redis.yaml中的name字段
	daprCli.store = "statestore"
	daprCli.pubsub = "pubsub"
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
	r.POST("/A", daprCli.A)
	r.GET("/dapr/subscribe", daprCli.Subscribe)
	r.Run()
}

// RequestInputs 获取所有参数
func RequestInputs(c *gin.Context) (map[string]interface{}, error) {

	const defaultMemory = 32 << 20
	contentType := c.ContentType()

	var (
		dataMap  = make(map[string]interface{})
		queryMap = make(map[string]interface{})
		postMap  = make(map[string]interface{})
	)

	// @see gin@v1.7.7/binding/query.go ==> func (queryBinding) Bind(req *http.Request, obj interface{})
	for k := range c.Request.URL.Query() {
		queryMap[k] = c.Query(k)
	}

	if "application/json" == contentType {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// @see gin@v1.7.7/binding/json.go ==> func (jsonBinding) Bind(req *http.Request, obj interface{})
		if c.Request != nil && c.Request.Body != nil {
			if err := json.NewDecoder(c.Request.Body).Decode(&postMap); err != nil {
				return nil, err
			}
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	} else if "multipart/form-data" == contentType {
		// @see gin@v1.7.7/binding/form.go ==> func (formMultipartBinding) Bind(req *http.Request, obj interface{})
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	} else {
		if err := c.Request.ParseForm(); err != nil {
			return nil, err
		}
		if err := c.Request.ParseMultipartForm(defaultMemory); err != nil {
			if err != http.ErrNotMultipart {
				return nil, err
			}
		}
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
	}

	var mu sync.RWMutex
	for k, v := range queryMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}
	for k, v := range postMap {
		mu.Lock()
		dataMap[k] = v
		mu.Unlock()
	}

	return dataMap, nil
}
