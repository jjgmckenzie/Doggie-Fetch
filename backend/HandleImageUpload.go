package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"gofetch/ghapp"
	"gofetch/postedimage"
	"log"
	"net/http"
	"time"
)

type ImageUploadHandler struct {
	builder           postedimage.Builder
	app               ghapp.App
	complianceHandler IsCompliantChecker
}

type httpContext interface {
	AbortWithStatus(code int)
	JSON(code int, object any)
	BindJSON(object any) error
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

func (i ImageUploadHandler) processImage(ctx context.Context, image postedimage.Image) (int, any) {
	isCompliant, err := i.complianceHandler.IsCompliant(image.Image)
	if err != nil {
		log.Printf("[DOG] [ERROR] An Error occured when determining Image Compliance: %s", err)
		return http.StatusInternalServerError, nil
	}
	if !isCompliant {
		println("[DOG] [ERROR] Image is not compliant.")
		return http.StatusPreconditionFailed, nil
	}
	println("[DOG] Image is compliant, pushing to Github")
	link, err := i.app.MakePullRequest(ctx, image)
	if err != nil {
		log.Printf("[DOG] [ERROR] An Error occured when posting to GitHub: %s", err)
		return http.StatusInternalServerError, nil
	}
	log.Printf("[DOG] Image was successfully pushed! GitHub Link : %s", link)
	return http.StatusAccepted, link
}

type imageReq struct {
	Name  string
	Breed string
	Image string
}

func (i ImageUploadHandler) getPostedImage(c httpContext) (postedimage.Image, error) {
	var imgReq imageReq
	var postedImage postedimage.Image
	err := c.BindJSON(&imgReq)
	if err == nil {
		postedImage, err = i.builder.Build(imgReq.Name, imgReq.Breed, imgReq.Image)
	}
	return postedImage, err
}

func (i ImageUploadHandler) HandleImageUpload(ctx httpContext) {
	img, err := i.getPostedImage(ctx)
	if err != nil {
		log.Printf("[DOG] [ERROR] An Error Occured: %s", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(i.processImage(ctx, img))
}

func (i ImageUploadHandler) HandleGinUpload(c *gin.Context) {
	println("[DOG] New Image Received!")
	i.HandleImageUpload(c)
}
