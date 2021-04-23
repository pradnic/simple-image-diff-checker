package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"
)
import "image"
import _ "image/jpeg"


func main() {
	i1 := readImage("diff-1.jpg")
	i2 := readImage("diff-2.jpg")
	if i1.Bounds() != i2.Bounds() {
		log.Fatalf("image 1 bounds %s, image 2 bounds %s",i1.Bounds(), i2.Bounds())
	}
	bounds := i1.Bounds()
	out1 := image.NewRGBA(bounds)
	out2 := image.NewRGBA(bounds)
	mismatchColor := color.RGBA{R: 10, G: 10, B: 255, A: 1}

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			p1 := i1.At(i,j)
			p2 := i2.At(i,j)
			if mismatch(p1,p2) {
				out1.Set(i,j,mismatchColor)
				out2.Set(i,j,p1)
			} else {
				out1.Set(i,j,p1)
				out2.Set(i,j,mismatchColor)
			}
		}
	}
	saveImage(out1, "out1.jpg")
	saveImage(out2, "out2.jpg")
}

func saveImage(out1 *image.RGBA, path string) {
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

func mismatch(p1 color.Color, p2 color.Color) bool {
	r1, g1, b1, _ := p1.RGBA()
	r2, g2, b2, _ := p2.RGBA()
	diffR := absDiff(r1, r2)
	diffG := absDiff(g1, g2)
	diffB := absDiff(b1, b2)
	diffTot := diffR + diffG + diffB
	return diffTot > 20000
	//return false
}

func absDiff(u1, u2 uint32) int {
	u := int(u1) - int(u2)
	if u > 0 {
		return u
	}
	return -u
}

func readImage(filename string) image.Image {
	open, err := os.Open(filename)
	if err != nil {
		log.Fatal(filename, err)
	}
	decode, s, err := image.Decode(open)
	if err != nil {
		log.Fatal(filename, err)
	}
	log.Print(s)
	return decode
}
