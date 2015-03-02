package cast

import (
	"encoding/hex"
	"hash/fnv"
	"io"
)

// Image is the *abstract* representation of an image file.
type Image struct {
	Id               string // Hash-based identity for the image
	Title            string
	OriginalFilename string
	Caption          string
	FocalPoint       Coordinate
}

type Coordinate struct {
	X, Y Pixel
}

type Pixel float32

func ImageId(data io.Reader) (string, error) {
	hasher := fnv.New32a()
	_, err := io.Copy(hasher, data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
