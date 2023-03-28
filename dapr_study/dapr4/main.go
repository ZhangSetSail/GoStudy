package main

import (
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	// 创建 Dapr 客户端
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 定义存储的命名空间
	namespace := "my-namespace"

	// 通过 Dapr 客户端创建存储操作的请求对象
	req := client.WithNamespace(namespace).SaveState("my-store", "my-key", []byte("my-value"))

	// 执行请求
	resp, err := req.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", resp)
}
