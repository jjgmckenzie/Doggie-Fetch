package main

import (
	"github.com/gin-gonic/gin"
	"gofetch/postedimage"
	"net/http"
)

type ImageUploadHandler struct {
	builder postedimage.Builder
}

func processImage(image postedimage.Image) (int, any) {
	if !isCompliant(image.Image) {
		return http.StatusPreconditionFailed, nil
	}
	return http.StatusAccepted, postImageToGitHub(image)
}

func (i ImageUploadHandler) getPostedImage(c *gin.Context) (postedimage.Image, error) {
	name := c.GetString("name")
	breed := c.GetString("breed")
	base64img := c.GetString("image")
	return i.builder.Build(name, breed, base64img)
}

func (i ImageUploadHandler) HandleImageUpload(c *gin.Context) {
	image, err := i.getPostedImage(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(processImage(image))
}
