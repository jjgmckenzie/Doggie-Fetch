package postedimage

import (
	"encoding/base64"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"
)

func decode(encodedImage string) (image.Image, error) {
	index := strings.Index(encodedImage, ",")
	if index < 0 {
		return nil, image.ErrFormat
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encodedImage[index+1:]))
	img, _, err := image.Decode(reader)
	return img, err
}
