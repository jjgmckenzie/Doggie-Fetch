package postedimage

import (
	"image"
	"os"
	"testing"
)

func TestImageRefusesInvalidDirectory(t *testing.T) {
	loadDir := "../test_images/test_image.png"
	// given an image file with a valid image
	imgFile, _ := os.Open(loadDir)
	img, _, _ := image.Decode(imgFile)
	newImage := Image{Image: img, Breed: "invalid_directory", Name: "test_saved_image"}
	// when asked to save the image to an invalid directory
	err := newImage.Save("../")
	if err == nil || !os.IsNotExist(err) {
		println(err.Error())
		t.Fail()
	}
}

func TestSaveImage(t *testing.T) {
	loadDir := "../test_images/test_image.png"
	saveDir := "../test_images/test_saved_image.jpg"
	// given an image file with a valid image
	imgFile, _ := os.Open(loadDir)
	img, _, _ := image.Decode(imgFile)

	newImage := Image{Image: img, Breed: "test_images", Name: "test_saved_image"}
	// when asked to save the image to a valid directory
	err := newImage.Save("../")
	// then there will not be an error,
	if err != nil {
		t.Fail()
	}
	// then the file will be readable,
	savedImgFile, err := os.Open(saveDir)
	if err != nil {
		t.Fail()
	}
	// then the file will be able to be decoded as an image
	savedImg, format, err := image.Decode(savedImgFile)
	if err != nil {
		t.Fail()
	}
	// and that format will be a jpeg
	if format != "jpeg" {
		t.Fail()
	}
	// and the image will be the same image as before
	r, g, b, _ := savedImg.At(0, 0).RGBA()
	oldR, oldG, oldB, _ := newImage.Image.At(0, 0).RGBA()
	if r>>8 != oldR>>8 || g>>8 != oldG>>8 || b>>8 != oldB>>8 {
		t.Fail()
	}

	// now clean up the file.
	_ = os.Remove(saveDir)

}
