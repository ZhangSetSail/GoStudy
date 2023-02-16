package router

import (
	client_go "github.com/ZhangSetSail/GoStudy/gin-demo/router/client-go"
	"github.com/gin-gonic/gin"
)

func InitRouter() error {
	r := gin.Default()
	client_go.Router(r.Group("/client-go"))
	return r.Run(":54321")
}
