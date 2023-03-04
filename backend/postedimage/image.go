package postedimage

import (
	"bytes"
	"github.com/sunshineplan/imgconv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"image"
	"io"
	"strings"
)

type Image struct {
	Name, Breed, nameFormatted string
	Image                      image.Image
}

func (i Image) Path() string {
	return i.Breed + "/" + i.nameFormatted + ".jpg"
}

func capitalize(s string) string {
	return cases.Title(language.English).String(s)
}

func formatBreed(breed string) string {
	breedString := strings.Split(breed, "-")
	hasSubBreed := len(breedString) > 1
	if !hasSubBreed {
		return breedString[0]
	}
	return breedString[1] + " " + breedString[0]
}

func (i Image) CommitMessage() string {
	breed := capitalize(formatBreed(i.Breed))
	return "Adds a " + breed + " named " + capitalize(i.Name)
}

func (i Image) AsBytes() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := i.save(buf)
	if err == nil && buf.Len() > 300_000 {
		buf.Reset()
		err = i.saveLossy(buf)
	}
	return buf.Bytes(), err
}

func (i Image) save(writer io.Writer) error {
	return imgconv.Write(writer, i.Image, &imgconv.FormatOption{Format: imgconv.JPEG})
}

func (i Image) saveLossy(writer io.Writer) error {
	resizedImage := resize(i.Image, 1080)
	return imgconv.Write(writer, resizedImage, &imgconv.FormatOption{
		Format: imgconv.JPEG,
		EncodeOption: []imgconv.EncodeOption{
			imgconv.Quality(70)},
	})
}
