package main

import (
	"flag"
	"fmt"
	"github.com/szabba/ffire/colorval"
	"github.com/szabba/ffire/forest"
	"image"
	"image/color"
	"image/png"
	"os"
)

func cellColor(cell forest.Cell) color.Color {

	if cell == forest.Tree {

		return colorval.NRGBA{color.NRGBA{0, 127, 0, 255}}
	} else if cell == forest.Space {
		return colorval.NRGBA{color.NRGBA{127, 255, 50, 255}}
	} else if cell == forest.Ash {
		return colorval.NRGBA{color.NRGBA{0, 0, 0, 255}}
	} else if cell == forest.Fire {
		return colorval.NRGBA{color.NRGBA{255, 0, 0, 255}}
	}

	return colorval.NRGBA{color.NRGBA{255, 0, 255, 255}}
}

var (
	maxZeros   int
	nameFormat string
)

func init() {

	flag.IntVar(&maxZeros, "zeros", 4, "How many zeros leading zeros to put in filenames?")

	flag.Parse()

	nameFormat = fmt.Sprintf("forest_%%0%dd.png", maxZeros)
}

func main() {

	var (
		g   forest.Grid
		err error
	)

	i := 0
	for _, err = fmt.Scan(&g); err == nil; _, err = fmt.Scan(&g) {

		w, h := g.Size()
		bounds := image.Rect(0, 0, w, h)
		img := image.NewNRGBA(bounds)

		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {

				img.Set(x, y, cellColor(g[y][x]))
			}
		}

		imgFile, _ := os.Create(fmt.Sprintf(nameFormat, i))
		png.Encode(imgFile, img)
		imgFile.Close()

		i++
	}

	fmt.Fprint(os.Stderr, err.Error())
}
