package main

import (
	"github.com/gin-gonic/gin"
	"gofetch/postedimage"
)

func main() {
	println("initializing git...")
	_, err := NewGitHandler("https://github.com/jigsawpieces/dog-api-images.git", "https://github.com/gofetchbot/dog-api-images.git")
	if err != nil {
		println("an error occurred, ")
		print(err.Error())
		return
	}
	println("git initialized.")
	handler := ImageUploadHandler{builder: postedimage.New()}
	router := gin.Default()
	router.POST("/upload", handler.HandleImageUpload)
	router.Run(":8080")
}
