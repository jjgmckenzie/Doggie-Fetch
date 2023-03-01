package main

import (
	"gocv.io/x/gocv"
	"image"
)

type ComplianceHandler struct {
	dogClassifier gocv.CascadeClassifier
}

func (c ComplianceHandler) Close() error {
	return c.dogClassifier.Close()
}

func (c ComplianceHandler) containsDog(image image.Image) bool {
	imgMat, err := gocv.ImageToMatRGB(image)
	return err == nil && len(c.dogClassifier.DetectMultiScale(imgMat)) > 0
}

func (c ComplianceHandler) containsHuman(image image.Image) bool {
	return false
}

func (c ComplianceHandler) IsCompliant(image image.Image) bool {
	return c.containsDog(image) && !c.containsHuman(image)
}

func NewComplianceHandler() ComplianceHandler {
	classifier := gocv.NewCascadeClassifier()
	classifier.Load("dog_face_haar_cascade/cascade.xml") // loads Haar cascade classifier
	return ComplianceHandler{
		dogClassifier: classifier,
	}
}
