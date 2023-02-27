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

type imgConvResizer struct {
	maxSize int
}

func (i imgConvResizer) getResizeOptions(imgSize rectangle) imgconv.ResizeOption {
	if imgSize.y < imgSize.x {
		return imgconv.ResizeOption{Width: i.maxSize}
	}
	return imgconv.ResizeOption{Height: i.maxSize}
}

func (i imgConvResizer) resize(img image.Image) image.Image {
	bounds := img.Bounds()
	imgSize := rectangle{x: bounds.Dx(), y: bounds.Dy()}
	maxSizeRect := rectangle{x: i.maxSize, y: i.maxSize}
	if imgSize.fitsInside(maxSizeRect) {
		return img
	}
	resizeOptions := i.getResizeOptions(imgSize)
	return imgconv.Resize(img, &resizeOptions)
}
