package postedimage

import "image"

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

func (b builder) Build(name, breed, base64img string) (Image, error) {
	img, err := b.decodeAndResize(base64img)
	return Image{
		Name:  name,
		Breed: breed,
		Image: img,
	}, err
}

func (b builder) decodeAndResize(encodedImage string) (image.Image, error) {
	img, err := b.decoder.decode(encodedImage)
	if err != nil {
		return nil, err
	}
	return b.resizer.resize(img), nil
}

func New() Builder {
	return builder{
		decoder: base64Decoder{},
		resizer: imgConvResizer{maxSize: 1080},
	}
}
