package postedimage

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestDecoderErrorsOnInvalidString(t *testing.T) {
	// given the image decoder and an invalid string
	decoder := base64Decoder{}
	invalidString := ""
	// when the decoder is given an invalid string
	_, err := decoder.decode(invalidString)
	// then the decoder will return an error
	if err == nil {
		t.Fail()
	}
}

func encodeTestImage(extension string) string {
	base64image := "data:image/" + extension + ";base64,"
	fileName := "../test_images/test_image." + extension
	bytes, _ := os.ReadFile(fileName)
	base64image += base64.StdEncoding.EncodeToString(bytes)
	return base64image
}

func canDecodeFormat(t *testing.T, format string) {
	// given the decoder and a valid encoded image
	decoder := base64Decoder{}
	base64image := encodeTestImage(format)
	// when the decoder is given a valid encoded image
	img, err := decoder.decode(base64image)
	// then the decoder will not error,
	if err != nil {
		println(base64image)
		t.Fail()
		return
	}

	// and will produce an image with the correct dimensions,
	imageHeightIsCorrect := img.Bounds().Dy() == 504
	imageWidthIsCorrect := img.Bounds().Dx() == 672
	if !imageHeightIsCorrect || !imageWidthIsCorrect {
		t.Fail()
	}

	// and the image produced matches the encoded image
	r, g, b, a := img.At(0, 0).RGBA()
	if r>>8 != 192 || g>>8 != 192 || b>>8 != 192 || a>>8 != 255 {
		t.Fail()
	}
}

func TestDecoderCanDecodePng(t *testing.T) {
	canDecodeFormat(t, "png")
}

func TestDecoderCanDecodeJpg(t *testing.T) {
	canDecodeFormat(t, "jpg")
}

func TestDecoderCanDecodeWebp(t *testing.T) {
	canDecodeFormat(t, "webp")
}

func TestDecoderCanDecodeGif(t *testing.T) {
	canDecodeFormat(t, "gif")
}

// no golang package for this, yet.
/* func TestDecoderCanDecodeAvif(t *testing.T) {
	canDecodeFormat(t, "avif")
} */
