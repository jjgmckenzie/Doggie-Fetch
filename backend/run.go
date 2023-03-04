package main

import (
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/gin-gonic/gin"
	"gofetch/ghapp"
	"gofetch/postedimage"
	"log"
	"net/http"
)

type Runner struct {
	isCompliant ComplianceHandler
	key         *ghinstallation.Transport
	port        string
}

func (r Runner) run(repository string) error {
	handler := ImageUploadHandler{builder: postedimage.New(), app: ghapp.New(r.key, repository), complianceHandler: r.isCompliant}
	router := gin.Default()
	router.POST("/upload", handler.HandleGinUpload)
	return router.Run(r.port)
}

func NewRunner(weightPath string, configPath string, cocoNamePath string, appId int64, installationId int64, pemKeyPath string, port string) (Runner, error) {
	compliant, err := NewComplianceHandler(weightPath, configPath, cocoNamePath)
	if err != nil {
		log.Printf("could not initialize neural network: %s", err.Error())
		return Runner{}, err
	}
	key, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appId, installationId, pemKeyPath)
	if err != nil {
		log.Printf("could not initialize runner: %s", err.Error())
		return Runner{}, err
	}
	return Runner{
		isCompliant: compliant,
		key:         key,
		port:        port,
	}, nil
}

func DefaultRunner(port string) Runner {
	runner, _ := NewRunner("./yolov3/yolov3.weights", "./yolov3/yolov3.cfg", "./yolov3/coco.names", 0, 0, "./key.pem", port)
	return runner
}
