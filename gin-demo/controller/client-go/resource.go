package client_go

import (
	"github.com/ZhangSetSail/GoStudy/gin-demo/handle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResourceManage struct {
}

func (p *ResourceManage) GetResourcesName(c *gin.Context) {
	namespace := c.Query("namespace")
	kind := c.Param("kind")
	resourceNames, err := handle.GetManagerClientGo().GetResourcesNameByNamespace(kind, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: 500,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": resourceNames,
		})
	}
}

func ()  {

}
