package postedimage

import (
	"github.com/sunshineplan/imgconv"
	"image"
)

type rectangle struct {
	x, y int
}

func (r rectangle) fitsInside(bounds rectangle) bool {
	return r.x < bounds.x && r.y < bounds.y
}

func getResizeOptions(imgSize rectangle, maxSize int) imgconv.ResizeOption {
	if imgSize.y < imgSize.x {
		return imgconv.ResizeOption{Width: maxSize}
	}
	return imgconv.ResizeOption{Height: maxSize}
}

func resize(img image.Image, maxSize int) image.Image {
	bounds := img.Bounds()
	imgSize := rectangle{x: bounds.Dx(), y: bounds.Dy()}
	maxSizeRect := rectangle{x: maxSize, y: maxSize}
	if imgSize.fitsInside(maxSizeRect) {
		return img
	}
	resizeOptions := getResizeOptions(imgSize, maxSize)
	return imgconv.Resize(img, &resizeOptions)
}
