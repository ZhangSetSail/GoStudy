package main

import (
	"github.com/dapr/go-sdk/service/common"
)

func main() {
	common.Service()
	common.Content{
		Data:        nil,
		ContentType: "",
		DataTypeURL: "",
	}
}
