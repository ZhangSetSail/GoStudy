package client_go

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WatchPodManage struct {
}

func (w WatchPodManage) GetWatchPodsName(c *gin.Context) {
	namespace := c.Query("namespace")
	podList, err := handle.GetManagerClientGo().GetWatchPodByNamespace(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: 500,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": podList,
	})
}
