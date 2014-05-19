package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func Size(bds image.Rectangle) (width, height int) {

	return bds.Max.X - bds.Min.X, bds.Max.Y - bds.Min.Y
}

func main() {

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
}
