package main

import (
	"image"
)

func containsDog(image image.Image) bool {
	return true
}

func containsHuman(image image.Image) bool {
	return false
}

func isCompliant(image image.Image) bool {
	return containsDog(image) && !containsHuman(image)
}
