package handler

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/bruno-sca/go-video-share-platform/helpers"
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

	filePath, fileName, err := createFileFolderAndReturnPath(file.Filename)

	if err != nil {
		logger.Errorf("Error creating file folder: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SaveUploadedFile(file, filePath)
	go processVideo(filePath, fileName)

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

func processVideo(filePath string, fileName string) {
	targetDir := filepath.Join("public", "videos", fileName)
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		logger.Errorf("Error creating target directory: %v", err)
		return
	}

	wg := helpers.WaitGroupHelper{}

	wg.Wrap(func() {transcodeVideo(filePath, fileName, targetDir)})
	wg.Wrap(func() {createThumbnail(filePath, fileName, targetDir)})

	wg.Wait()

	logger.Info("Removing temp file", filePath)

	os.Remove(filePath)
}

func createThumbnail(fp string, fn string, td string) {
	logger.Info("Creating thumbnail")
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i",
		fp,
		"-vf",
		"thumbnail",
		"-frames:v",
		"1",
		filepath.Join(td, fn + ".png"),
	)

	if output, err := ffmpegCmd.CombinedOutput(); err != nil {
		logger.Errorf("Error creating thumbnail: %v\nOutput: %s", err, string(output))
		return
	}
	logger.Info("Thumbnail created")
}

func transcodeVideo(fp string, fn string, td string) {
	logger.Info("Transcoding video")
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i",
		fp,
		"-filter:v",
		"scale='min(1280,iw)':min'(720,ih)':force_original_aspect_ratio=decrease,pad=1280:720:(ow-iw)/2:(oh-ih)/2",
		filepath.Join(td, fn + ".mp4"),
	)

	if output, err := ffmpegCmd.CombinedOutput(); err != nil {
		logger.Errorf("Error creating thumbnail: %v\nOutput: %s", err, string(output))
		return
	}
	logger.Info("Video transcoded")
}