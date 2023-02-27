package main

import (
	"github.com/gin-gonic/gin"
	"gofetch/postedimage"
)

func main() {
	handler := ImageUploadHandler{postedimage.New()}
	router := gin.Default()
	router.POST("/upload", handler.HandleImageUpload)
	router.Run(":8080")
}
