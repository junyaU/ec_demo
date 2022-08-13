package merchandise

import (
	"errors"
	"strings"
)

type printImage struct {
	path      string
	extension string
}

func newPrintImage(image string) (printImage, error) {
	imageSplit := strings.Split(image, ".")
	if len(imageSplit) < 2 || imageSplit[0] == "" {
		return printImage{}, errors.New("invalid image path")
	}

	switch imageSplit[1] {
	case "jpg", "jpeg", "png":
		return printImage{path: imageSplit[0], extension: imageSplit[1]}, nil

	default:
		return printImage{}, errors.New("cannot register extensions other than jpg and png")
	}
}
