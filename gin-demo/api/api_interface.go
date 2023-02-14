package api

import "github.com/gin-gonic/gin"

type PodInterface interface {
	GetPod(c *gin.Context)
	GetPodsName(c *gin.Context)
	DeletePod(c *gin.Context)
}
