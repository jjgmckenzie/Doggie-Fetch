package main

import (
	"github.com/gin-gonic/gin"
	"gofetch/postedimage"
	"net/http"
)

type ImageUploadHandler struct {
	builder           postedimage.Builder
	gitHubHandler     GitHubHandler
	complianceHandler IsCompliantChecker
}

type httpContext interface {
	GetString(key string) string
	AbortWithStatus(code int)
	JSON(code int, object any)
}

func (i ImageUploadHandler) processImage(image postedimage.Image) (int, any) {
	isCompliant, err := i.complianceHandler.IsCompliant(image.Image)
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	if !isCompliant {
		return http.StatusPreconditionFailed, nil
	}
	link, err := i.gitHubHandler.PostToGithub(image, image.GetCommitMessage())
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	return http.StatusAccepted, link
}

func (i ImageUploadHandler) getPostedImage(c httpContext) (postedimage.Image, error) {
	name := c.GetString("name")
	breed := c.GetString("breed")
	base64img := c.GetString("image")
	return i.builder.Build(name, breed, base64img)
}

func (i ImageUploadHandler) HandleImageUpload(c httpContext) {
	img, err := i.getPostedImage(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(i.processImage(img))
}

func (i ImageUploadHandler) HandleGinUpload(c *gin.Context) {
	i.HandleImageUpload(c)
}
