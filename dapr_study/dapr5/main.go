package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	if err := c.ShouldBindBodyWith(&b, binding.JSON); err != nil {
		c.Error(err)
		return
	}
	fmt.Println("A:", b.Data.Message)
	c.String(http.StatusOK, "Successfully A")
}

func C(c *gin.Context) {
	var b Body
	if err := c.ShouldBindBodyWith(&b, binding.JSON); err != nil {
		c.Error(err)
		return
	}
	fmt.Println("C:", b.Data.Message)
	c.String(http.StatusOK, "Successfully C")
}

func Subscribe(c *gin.Context) {
	c.JSON(http.StatusOK, []map[string]string{
		{
			"pubsubname": "pubsub",
			"topic":      "A",
			"route":      "A",
		}, {
			"pubsubname": "pubsub",
			"topic":      "C",
			"route":      "C",
		},
	},
	)
}

func main() {
	r := gin.Default()
	r.POST("/A", A)
	r.POST("/C", A)
	r.GET("/dapr/subscribe", Subscribe)
	r.Run()
}
