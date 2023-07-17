package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}