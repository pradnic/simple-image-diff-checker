package util

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func ReadImage(filename string) image.Image {
	open, err := os.Open(filename)
	if err != nil {
		log.Fatal(filename, " ", err)
	}
	decode, s, err := image.Decode(open)
	if err != nil {
		log.Fatal(filename, " ", err)
	}
	log.Print(s)
	return decode
}

func SaveImage(out1 *image.RGBA, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = jpeg.Encode(file, out1, nil)
	if err != nil {
		log.Fatal(err)
	}
}
