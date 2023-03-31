package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Text struct {
	Message     string `json:"message"`
	MessageType string `json:"messageType"`
}

type Body struct {
	Data Text `json:"data"`
}

func A(c *gin.Context) {
	var b Body
	if err := c.ShouldBindQuery(&b); err != nil {
		c.Error(err)
		return
	}
	fmt.Println(b.Data.MessageType, ":", b.Data.Message)
	c.String(http.StatusOK, "Successfully A")
}

func Subscribe(c *gin.Context) {
	c.JSON(http.StatusOK, []map[string]string{
		{
			"pubsubname": "pubsub",
			"topic":      "A",
			"route":      "A",
		},
	},
	)
}

func main() {
	r := gin.Default()
	r.POST("/A", A)
	r.GET("/dapr/subscribe", Subscribe)
	r.Run()
}
