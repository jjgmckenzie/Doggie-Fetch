package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.Fatalln(DefaultRunner(":8080").run("dog-api-images"))
}
