package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"log"
)

type ComplianceHandler struct {
	dogClassifier gocv.CascadeClassifier
	humanDetector gocv.HOGDescriptor
}

func (c ComplianceHandler) Close() error {
	return c.dogClassifier.Close()
}

func (c ComplianceHandler) containsDog(image image.Image) bool {
	println("converting image to be readable")
	imgMat, err := gocv.ImageToMatRGB(image)
	println("converted image, searching for dog")
	dogs := c.dogClassifier.DetectMultiScale(imgMat)
	fmt.Println(len(dogs), " dogs found")
	return err == nil && len(dogs) > 0
}

func (c ComplianceHandler) containsHuman(image image.Image) bool {
	println("converting image to be readable")
	imgMat, err := gocv.ImageToMatRGB(image)
	println("converted image, searching for human")
	humans := c.humanDetector.DetectMultiScale(imgMat)
	fmt.Println(len(humans), " humans found at", humans[0].String())
	return err == nil && len(humans) > 0
}

func (c ComplianceHandler) IsCompliant(image image.Image) bool {
	return c.containsDog(image) && !c.containsHuman(image)
}

func NewComplianceHandler() ComplianceHandler {
	classifier := gocv.NewCascadeClassifier()
	classifier.Load("dog_face_haar_cascade/cascade.xml") // loads Haar cascade classifier
	humanDetector := gocv.NewHOGDescriptor()
	err := humanDetector.SetSVMDetector(gocv.HOGDefaultPeopleDetector())
	if err != nil {
		log.Fatalf("an error occured when setting up compliance handler: %s", err.Error())
	}
	return ComplianceHandler{
		dogClassifier: classifier,
		humanDetector: humanDetector,
	}
}
