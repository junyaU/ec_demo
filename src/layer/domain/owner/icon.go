package owner

import (
	"errors"
	"strings"
)

type icon struct {
	path      string
	extension string
}

func newIcon(iconPath string) (icon, error) {
	pathSplit := strings.Split(iconPath, ".")
	if len(pathSplit) < 2 || pathSplit[0] == "" {
		return icon{}, errors.New("invalid image path")
	}

	switch pathSplit[1] {
	case "jpg", "jpeg", "png":
		return icon{path: pathSplit[0], extension: pathSplit[1]}, nil

	default:
		return icon{}, errors.New("cannot register extensions other than jpg and png")
	}
}
