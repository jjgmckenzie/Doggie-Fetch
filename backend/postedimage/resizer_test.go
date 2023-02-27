package postedimage

import (
	"github.com/sunshineplan/imgconv"
	"image"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestResizeDoesNotChangeSmallerImageDimensions(t *testing.T) {
	// Given an Image Resizer with max size of 1080px
	imgResizer := imgConvResizer{1080}
	// When the Image Resizer is given a file that is smaller than that size
	imgFile, _ := os.Open("../test_images/test_image.png")
	img, _, _ := image.Decode(imgFile)
	processedImg := imgResizer.resize(img)
	// The Image Resizer does not change the file's dimensions
	widthUnchanged := processedImg.Bounds().Dx() == img.Bounds().Dx()
	heightUnchanged := processedImg.Bounds().Dy() == img.Bounds().Dy()
	if !widthUnchanged || !heightUnchanged {
		t.Fail()
	}
}

func TestResizerScalesWidthWhenWidthLarger(t *testing.T) {
	// Given an Image Resizer with max size of clamp, where clamp is a random number lying within [0, 1080)
	rand.Seed(time.Now().UnixNano())
	clamp := rand.Intn(1080)
	imgResizer := imgConvResizer{clamp}
	// When the Image Resizer is given an item with a width of 1920, and height of 1080
	options := imgResizer.getResizeOptions(rectangle{x: 1920, y: 1080})
	// The resizer clamps only width, the larger of the two, to the clamp; maintaining aspect ratio.
	clampedWidth := imgconv.ResizeOption{Width: clamp}
	if options != clampedWidth {
		t.Fail()
	}
}

func TestResizerScalesHeightWhenHeightLarger(t *testing.T) {
	// Given an Image Resizer with max size of clamp, where clamp is a random number lying within [0, 1080)
	rand.Seed(time.Now().UnixNano())
	clamp := rand.Intn(1080)
	imgResizer := imgConvResizer{clamp}
	// When the Image Resizer is given an item with a width of 1080, and height of 1920
	options := imgResizer.getResizeOptions(rectangle{x: 1080, y: 1920})
	// The resizer clamps only height, the larger of the two, to the clamp; maintaining aspect ratio.
	clampedHeight := imgconv.ResizeOption{Height: clamp}
	if options != clampedHeight {
		t.Fail()
	}
}

func TestResizeChangesBiggerImageDimensions(t *testing.T) {
	// Given an Image Resizer with max size of 336px
	imgResizer := imgConvResizer{336}
	// When the Image Resizer is given a file that is twice that size on one axis
	imgFile, _ := os.Open("../test_images/test_image.png")
	img, _, _ := image.Decode(imgFile)
	processedImg := imgResizer.resize(img)
	// The Image Resizer changes both file's dimensions to half, maintaining aspect ratio
	widthRightRatio := processedImg.Bounds().Dx() == (img.Bounds().Dx() / 2)
	heightRightRatio := processedImg.Bounds().Dy() == (img.Bounds().Dy() / 2)
	if !widthRightRatio || !heightRightRatio {
		t.Fail()
	}
}
