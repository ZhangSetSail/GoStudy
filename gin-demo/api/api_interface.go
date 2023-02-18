package api

import "github.com/gin-gonic/gin"

type ResourceInterface interface {
	GetResourcesName(c *gin.Context)
}

type WatchResourceInterface interface {
	GetWatchResourcesName(c *gin.Context)
}
