package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() error {
	r := gin.Default()
	PodRouter(r)
	return r.Run(":54321")
}
