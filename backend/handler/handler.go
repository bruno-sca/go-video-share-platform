package handler

import "github.com/bruno-sca/go-video-share-platform/config"


var (
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	// Initialize dependencies here
}