package main

import (
	"flag"
	"fmt"
	"github.com/szabba/ffire/colorval"
	"image"
	"image/color"
	_ "image/png"
	"os"
)

func EqualColors(a, b color.Color) bool {

	aR, aG, aB, aA := a.RGBA()
	bR, bG, bB, bA := b.RGBA()

	return aR == bR && aG == bG && aB == bB && aA == bA
}

func Size(bds image.Rectangle) (width, height int) {

	return bds.Max.X - bds.Min.X, bds.Max.Y - bds.Min.Y
}

func PrintImage(filename string) {

	imgFile, err := os.Open(filename)
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

			if EqualColors(colorval.Space, color) {
				fmt.Print("S")

			} else if EqualColors(colorval.Tree, color) {
				fmt.Print("T")

			} else if EqualColors(colorval.Fire, color) {
				fmt.Print("F")

			} else if EqualColors(colorval.Ash, color) {
				fmt.Print("A")

			} else {

				r, g, b, a := color.RGBA()

				fmt.Fprintf(
					os.Stderr,
					"Don't know how to interpret color #%x%x%x%x",
					r, g, b, a,
				)
				os.Exit(1)
			}

			if x+1 == bds.Max.X {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}

var fire_prob float64

func init() {
	flag.Float64Var(&fire_prob, "p", 0.01, "probability of fire")

	flag.Parse()
}

func main() {

	for _, arg := range flag.Args() {

		PrintImage(arg)
	}
}
