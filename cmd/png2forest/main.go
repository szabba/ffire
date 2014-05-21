package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"math/rand"
	"os"
	"time"
)

func EqualColors(a, b color.Color) bool {

	aR, aG, aB, aA := a.RGBA()
	bR, bG, bB, bA := b.RGBA()

	return aR == bR && aG == bG && aB == bB && aA == bA
}

func Size(bds image.Rectangle) (width, height int) {

	return bds.Max.X - bds.Min.X, bds.Max.Y - bds.Min.Y
}

var fire_prob float64

func init() {
	flag.Float64Var(&fire_prob, "p", 0.01, "probability of fire")

	flag.Parse()
}

func main() {

	rand.Seed(time.Now().Unix())

	imgFile, err := os.Open("forest.png")
	if err != nil {

		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	var img image.Image
	img, _, err = image.Decode(imgFile)
	if err != nil {

		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(2)
	}

	bds := img.Bounds()
	w, h := Size(bds)

	fmt.Printf("%d %d\n", w, h)

	for y := bds.Min.Y; y < bds.Max.Y; y++ {
		for x := bds.Min.X; x < bds.Max.X; x++ {

			color := img.At(x, y)

			red, _, _, _ := color.RGBA()

			if red != 0 {
				fmt.Print("S")
			} else {
				if rand.Float64() < fire_prob {
					fmt.Print("F")
				} else {
					fmt.Print("T")
				}
			}

			if x+1 == bds.Max.X {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}
