package main

import (
	"github.com/wimspaargaren/yolov3"
	"gocv.io/x/gocv"
	"image"
	"log"
)

type ComplianceHandler struct {
	neuralNetwork yolov3.Net
}

func (c ComplianceHandler) getContents(image image.Image) ([]yolov3.ObjectDetection, error) {
	var detection []yolov3.ObjectDetection = nil
	imgMat, err := gocv.ImageToMatRGB(image)
	if err == nil {
		detection, err = c.neuralNetwork.GetDetections(imgMat)
	}
	return detection, err
}

func (c ComplianceHandler) contains(className string, objects []yolov3.ObjectDetection) bool {
	for _, object := range objects {
		if object.ClassName == className {
			return true
		}
	}
	return false
}

func (c ComplianceHandler) containsDog(objects []yolov3.ObjectDetection) bool {
	return c.contains("dog", objects)
}

func (c ComplianceHandler) containsHuman(objects []yolov3.ObjectDetection) bool {
	return c.contains("person", objects)
}

func (c ComplianceHandler) IsCompliant(image image.Image) (bool, error) {
	isCompliant := false
	contents, err := c.getContents(image)
	if err == nil {
		isCompliant = c.containsDog(contents) && !c.containsHuman(contents)
	}
	return isCompliant, err
}

func NewComplianceHandler() ComplianceHandler {

	yolonet, err := yolov3.NewNet("./yolov3/yolov3.weights", "./yolov3/yolov3.cfg", "./yolov3/coco.names")
	if err != nil {
		log.Fatalf("an error occured when setting up compliance handler: %s", err.Error())
	}

	return ComplianceHandler{
		neuralNetwork: yolonet,
	}
}
