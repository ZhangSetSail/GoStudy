package client_go

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PodsManage struct {
}

func (p *PodsManage) GetPodDetails(c *gin.Context) {
}

func (p *PodsManage) GetPodsName(c *gin.Context) {
	namespace := c.Query("namespace")
	podList, err := handle.GetManagerClientGo().GetPodByNamespace(namespace)
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

func (p *PodsManage) DeletePod(c *gin.Context) {

}
