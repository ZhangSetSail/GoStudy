package api

import "github.com/gin-gonic/gin"

type PodInterface interface {
	GetPodDetails(c *gin.Context)
	GetPodsName(c *gin.Context)
	DeletePod(c *gin.Context)
}

type WatchPodInterface interface {
	GetWatchPodsName(c *gin.Context)
}
