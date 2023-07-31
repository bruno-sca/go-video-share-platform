package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadVideoHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		logger.Errorf("Error uploading file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Uploading file: ", file.Filename)

	filePath, _, err := createFileFolderAndReturnPath(file.Filename)

	if err != nil {
		logger.Errorf("Error creating file folder: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SaveUploadedFile(file, filePath)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully!",
	})
}

func createFileFolderAndReturnPath(fn string) (string, string, error) {
	fileExt := filepath.Ext(fn)
	fileName := strings.Split(fn, fileExt)[0]

	destFolder := filepath.Join("tmp", "processing")
	destPath := filepath.Join(destFolder, uuid.New().String() + fileExt);

	if err := os.MkdirAll(destFolder, os.ModePerm); err != nil {
		return "", "", err
	}

	return destPath, fileName, nil
}

