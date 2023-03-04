package postedimage

import (
	"image"
	"os"
	"testing"
)

func TestImagePath(t *testing.T) {
	img := Image{nameFormatted: "maya",
		Breed: "pyrenees",
	}
	result := img.Path()
	if result != "pyrenees/maya.jpg" {
		t.Fail()
	}
}

func TestImage_AsBytes(t *testing.T) {
	loadDir := "../test_images/large_dog_image.jpg"
	imgFile, _ := os.Open(loadDir)
	img, _, _ := image.Decode(imgFile)
	newImage := Image{Image: img}
	bytes, err := newImage.AsBytes()
	if err != nil {
		t.Fail()
	}
	if len(bytes) > 300_000 {
		t.Fail()
	}
}
