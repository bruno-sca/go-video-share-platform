package router

import (
	"github.com/bruno-sca/go-video-share-platform/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api"

	v1 := router.Group(basePath + "/v1")
	{
		v1.POST("/upload", handler.UploadVideoHandler)
	}
}