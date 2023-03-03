package main

import (
	"context"
	"flag"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/gin-gonic/gin"
	"gofetch/postedimage"
	"log"
	"net/http"
)

func main() {
	var appId int64
	var installationId int64
	var privateKey string

	flag.Int64Var(&appId, "appId", 0, "the app ID of the the bot to run")
	flag.Int64Var(&installationId, "installation ID", 0, "the installation ID of the the bot to run")
	flag.StringVar(&privateKey, "debug", "", "the private PEM encoded key of the bot to run")

	println("initializing git...")
	gh, err := ghinstallation.New(http.DefaultTransport, appId, installationId, []byte(privateKey))
	if err != nil {
		log.Fatalln(err)
	}
	token, err := gh.Token(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	git, err := NewGitHandler("./dog-api-images", Auth{username: "git", token: token})
	if err != nil {
		log.Fatalln(err)
	}
	println("git initialized.")
	complianceHandler, err := NewComplianceHandler()
	if err != nil {
		log.Fatalf("an error occured when setting up compliance handler: %s", err.Error())
	}
	handler := ImageUploadHandler{builder: postedimage.New(), gitHubHandler: gitHubHandler{gitHandler: git}, complianceHandler: complianceHandler}
	router := gin.Default()
	router.POST("/upload", handler.HandleImageUpload)
	log.Fatalln(router.Run(":8080"))
}
