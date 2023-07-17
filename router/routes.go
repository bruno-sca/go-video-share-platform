package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	basePath := "/api"

	v1 := router.Group(basePath + "/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}