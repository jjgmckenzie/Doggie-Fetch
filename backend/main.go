package main

import (
	"github.com/gin-gonic/gin"
	"gofetch/postedimage"
	"log"
)

func main() {

	gh := NewGitHubHandler("./dog-api-images")
	compliant, err := NewComplianceHandler()
	if err != nil {
		log.Fatalf("an error occured when setting up compliance handler: %s", err.Error())
	}
	handler := ImageUploadHandler{builder: postedimage.New(), gitHubHandler: gh, complianceHandler: compliant}
	router := gin.Default()
	router.POST("/upload", handler.HandleGinUpload)
	log.Fatalln(router.Run(":8080"))
}
