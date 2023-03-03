package postedimage

import (
	"github.com/sunshineplan/imgconv"
	"image"
	"os"
)

type Image struct {
	Name, Breed string
	Image       image.Image
}

func (i Image) Save(directory string) error {
	subDirectory := directory + i.Breed

	destination := subDirectory + "/" + i.Name + ".jpg"
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
