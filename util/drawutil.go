package util

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"strings"
)

func DrawOver(baseImage draw.Image, layer image.Image) {
	draw.Draw(baseImage, baseImage.Bounds(), layer, image.ZP, draw.Over)
}

func DrawOverLocation(baseImage draw.Image, layer image.Image, x, y int) {
	draw.Draw(baseImage, baseImage.Bounds(), layer, image.Pt(x, y), draw.Over)
}

func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("Error cropping image")
	}

	return simg.SubImage(crop), nil
}

func GetImageFromFileNoCache(fileName string) (image.Image, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening file %s", fileName)
	}
	defer f.Close()

	if f.Name() != fileName {
		return nil, fmt.Errorf("File names don't match %s : %s", f.Name(), fileName)
	}

	var imgD image.Image
	if strings.Contains(fileName, "jpg") {
		imgD, err = jpeg.Decode(f)
		if err != nil {
			return nil, fmt.Errorf("Error decoding image %s", fileName)
		}
	} else {
		imgD, _, err = image.Decode(f)
		if err != nil {
			return nil, fmt.Errorf("Error decoding image %s", fileName)
		}
	}

	return imgD, nil
}
