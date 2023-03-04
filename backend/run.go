package main

import (
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/gin-gonic/gin"
	"gofetch/ghapp"
	"gofetch/postedimage"
	"net/http"
)

func run(repository string) error {
	transport, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 0, 0, "./key.pem")
	if err != nil {
		return err
	}
	compliant, err := NewComplianceHandler()
	if err != nil {
		return err
	}
	handler := ImageUploadHandler{builder: postedimage.New(), app: ghapp.New(transport, repository), complianceHandler: compliant}
	router := gin.Default()
	router.POST("/upload", handler.HandleGinUpload)
	return router.Run(":8080")
}
