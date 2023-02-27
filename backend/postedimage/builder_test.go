package postedimage

import (
	"errors"
	"image"
	"os"
	"testing"
)

type mockDecoder struct {
	err string
}

func (m mockDecoder) decode(base64 string) (image.Image, error) {
	if m.err != "" {
		return nil, errors.New(m.err)
	}
	return nil, nil
}

type mockResizer struct {
	image image.Image
}

func (m mockResizer) resize(_ image.Image) image.Image {
	return m.image
}

func TestBuilderPropagatesError(t *testing.T) {
	_, err := builder{decoder: mockDecoder{err: "decoder that errors"}, resizer: mockResizer{}}.Build("", "", "")
	if err == nil || err.Error() != "decoder that errors" {
		t.Fail()
	}
}

func TestBuilderCallsResizer(t *testing.T) {
	imgFile, _ := os.Open("../test_images/test_image.png")
	mockResizedImg, _, _ := image.Decode(imgFile)
	mockResizer := mockResizer{mockResizedImg}
	returnedImage, _ := builder{decoder: mockDecoder{}, resizer: mockResizer}.Build("", "", "")
	if returnedImage.Image.At(0, 0) != mockResizedImg.At(0, 0) {
		t.Fail()
	}
}
