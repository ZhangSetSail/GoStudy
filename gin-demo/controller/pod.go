package controller

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PodsManage struct {
}

func (p *PodsManage) GetPod(c *gin.Context) {
	namespace := c.Query("namespace")
	pm := handle.GetManagerPod()
	podList, err := pm.GetPodByNamespace(namespace)
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

func (p *PodsManage) GetPods(c *gin.Context) {

}

func (p *PodsManage) DeletePod(c *gin.Context) {

}
