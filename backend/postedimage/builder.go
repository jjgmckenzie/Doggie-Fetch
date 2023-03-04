package postedimage

import (
	"errors"
	goaway "github.com/TwiN/go-away"
	"image"
	"regexp"
	"time"
)

type Builder interface {
	Build(name, breed, base64img string) (Image, error)
}

type builder struct {
}

func asASCII(name string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	return re.ReplaceAllLiteralString(name, "")
}

func truncate(name string) string {
	if len(name) > 20 {
		name = name[:20]
	}
	return name
}

func filter(name string) (string, error) {
	var err error
	asciiName := asASCII(name)
	if goaway.IsProfane(asciiName) { // really hate that this has to be done, but the risk outweighs the false negatives.
		err = errors.New("dog's name is profane :(")
	}
	truncatedAsciiName := truncate(name)
	return truncatedAsciiName, err
}

func formatted(name string) string {
	return "gofetchbot_" + name + "_" + time.Now().UTC().Format(time.RFC3339)
}

func (b builder) Build(name, breed, base64img string) (Image, error) {
	var img image.Image
	filteredName, err := filter(name)
	if err == nil {
		img, err = decode(base64img)
	}
	return Image{
		Name:          filteredName,
		Breed:         breed,
		nameFormatted: formatted(filteredName),
		Image:         img,
	}, err
}

func New() Builder {
	return builder{}
}
