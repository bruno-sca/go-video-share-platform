package main

import (
	"github.com/bruno-sca/go-video-share-platform/config"
	"github.com/bruno-sca/go-video-share-platform/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	if err := config.Init(); err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}