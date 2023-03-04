package postedimage

import (
	"github.com/sunshineplan/imgconv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"image"
	"os"
	"strings"
)

type Image struct {
	Name, Breed, nameFormatted string
	Image                      image.Image
}

func (i Image) Save(directory string) error {
	subDirectory := directory + i.Breed

	destination := subDirectory + "/" + i.nameFormatted + ".jpg"
	f, err := os.Create(destination)
	if err == nil {
		err = imgconv.Write(f, i.Image, &imgconv.FormatOption{
			Format: imgconv.JPEG,
			EncodeOption: []imgconv.EncodeOption{
				imgconv.Quality(70)},
		})
	}
	return err
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

func (i Image) GetCommitMessage() string {
	breed := capitalize(formatBreed(i.Breed))
	return "GOFETCHBOT: Adds " + capitalize(i.Name) + ", a user submitted " + (breed)
}
