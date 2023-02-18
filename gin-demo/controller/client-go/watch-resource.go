package client_go

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WatchResourceManage struct {
}

func (w WatchResourceManage) GetWatchResourcesName(c *gin.Context) {
	namespace := c.Query("namespace")
	kind := c.Param("kind")
	resourceList, err := handle.GetManagerClientGo().GetWatchResourcesByNamespace(kind, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: 500,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": resourceList,
	})
}
