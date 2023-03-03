package postedimage

import (
	"errors"
	goaway "github.com/TwiN/go-away"
	"image"
	"regexp"
	"time"
)

type resizer interface {
	resize(img image.Image) image.Image
}

type decoder interface {
	decode(string) (image.Image, error)
}

type Builder interface {
	Build(name, breed, base64img string) (Image, error)
}

type builder struct {
	decoder decoder
	resizer resizer
}

func asASCII(name string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	return re.ReplaceAllLiteralString(name, "")
}

func filter(name string) (string, error) {
	var err error
	asciiName := asASCII(name)
	if goaway.IsProfane(asciiName) { // really hate that this has to be done, but the risk outweighs the false negatives.
		err = errors.New("dog's name is profane :(")
	}
	return asciiName, err
}

func formatted(name string) string {
	return "gofetchbot_" + name + "_" + time.Now().UTC().Format(time.RFC3339)
}

func (b builder) Build(name, breed, base64img string) (Image, error) {
	var img image.Image
	filteredName, err := filter(name)
	if err == nil {
		img, err = b.decodeAndResize(base64img)
	}
	return Image{
		Name:          filteredName,
		Breed:         breed,
		nameFormatted: formatted(filteredName),
		Image:         img,
	}, err
}

func (b builder) decodeAndResize(encodedImage string) (image.Image, error) {
	img, err := b.decoder.decode(encodedImage)
	if err == nil {
		img = b.resizer.resize(img)
	}
	return img, err
}

func New() Builder {
	return builder{
		decoder: base64Decoder{},
		resizer: imgConvResizer{maxSize: 1080},
	}
}
