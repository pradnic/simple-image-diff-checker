package main

import (
	"flag"
	"image"
	"image/color"
	"log"
	"simple-image-diff-checker/util"
)

var (
	help = flag.Bool("h", false, "Prints help")
	in1 = flag.String("in1", "in1.jpg", "Set path to input 1")
	in2 = flag.String("in2", "in2.jpg", "Set path to input 2")
	out1 = flag.String("out1", "out1.jpg", "Set path to output 1")
	out2 = flag.String("out2", "out2.jpg", "Set path to output 2")
	R = flag.Uint("r", 255, "Set RGBA R value")
	G = flag.Uint("g", 255, "Set RGBA G value")
	B = flag.Uint("b", 255, "Set RGBA B value")
	A = flag.Uint("a", 1, "Set RGBA A value")
)

func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}
	mismatchColor := color.RGBA{R: uint8(*R), G: uint8(*G), B: uint8(*B), A: uint8(*A)}
	compareImages(*in1, *in2, *out1, *out2, mismatchColor)
}

func compareImages(in1, in2, out1, out2 string, mismatchColor color.RGBA) {
	i1 := util.ReadImage(in1)
	i2 := util.ReadImage(in2)
	if i1.Bounds() != i2.Bounds() {
		log.Fatalf("image 1 bounds %s, image 2 bounds %s", i1.Bounds(), i2.Bounds())
	}
	bounds := i1.Bounds()
	imgOut1 := image.NewRGBA(bounds)
	imgOut2 := image.NewRGBA(bounds)

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			p1 := i1.At(i, j)
			p2 := i2.At(i, j)
			if util.Mismatch(p1, p2) {
				imgOut1.Set(i, j, mismatchColor)
				imgOut2.Set(i, j, p1)
			} else {
				imgOut1.Set(i, j, p1)
				imgOut2.Set(i, j, mismatchColor)
			}
		}
	}
	util.SaveImage(imgOut1, out1)
	util.SaveImage(imgOut2, out2)
}