package main

import (
	"github.com/gin-gonic/gin"
	"gofetch/postedimage"
	"net/http"
)

type ImageUploadHandler struct {
	builder           postedimage.Builder
	gitHubHandler     GitHubHandler
	complianceHandler ComplianceHandler
}

func (i ImageUploadHandler) processImage(image postedimage.Image) (int, any) {
	if !i.complianceHandler.IsCompliant(image.Image) {
		return http.StatusPreconditionFailed, nil
	}
	link, err := i.gitHubHandler.PostToGithub(image)
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	return http.StatusAccepted, link
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
	c.JSON(i.processImage(image))
}
