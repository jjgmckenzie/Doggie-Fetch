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

func (m mockDecoder) decode(_ string) (image.Image, error) {
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

func TestProfanityFilter(t *testing.T) {
	// given a profane word and a not profane, similar word
	profaneWord := "ass"
	notProfaneWord := "lassie"
	// when the profanity filter considers them
	_, profaneWordProfane := filter(profaneWord)
	_, notProfaneWordProfane := filter(notProfaneWord)
	// then the profane word will error, and the not profane word will not
	if profaneWordProfane == nil {
		t.Fail()
	}
	if notProfaneWordProfane != nil {
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

func TestTruncateString(t *testing.T) {
	if truncate("test") != "test" {
		t.Fail()
	}
	if len(truncate("abcdefghijklmnopqrstuvwxyz")) != 20 {
		t.Fail()
	}
}
