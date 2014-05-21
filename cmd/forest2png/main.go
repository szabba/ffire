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

var (
	maxZeros   int
	nameFormat string

	treeColor  = colorval.NRGBA{color.NRGBA{0, 127, 0, 255}}
	spaceColor = colorval.NRGBA{color.NRGBA{127, 255, 50, 255}}
	fireColor  = colorval.NRGBA{color.NRGBA{255, 0, 0, 255}}
	ashColor   = colorval.NRGBA{color.NRGBA{127, 127, 127, 255}}
	errorColor = colorval.NRGBA{color.NRGBA{255, 0, 255, 255}}
)

func init() {

	flag.IntVar(&maxZeros, "zeros", 4, "How many zeros leading zeros to put in filenames?")

	flag.Var(&treeColor, "tree-color", "The color of trees")
	flag.Var(&spaceColor, "space-color", "The color of free space")
	flag.Var(&fireColor, "fire-color", "The color of fire")
	flag.Var(&ashColor, "ash-color", "The color of ashes")
	flag.Var(&errorColor, "err-color", "Color for errnoerously encoded cells")

	flag.Parse()

	nameFormat = fmt.Sprintf("forest_%%0%dd.png", maxZeros)
}

func cellColor(cell forest.Cell) color.Color {

	if cell == forest.Tree {
		return treeColor

	} else if cell == forest.Space {
		return spaceColor

	} else if cell == forest.Ash {
		return ashColor

	} else if cell == forest.Fire {
		return fireColor
	}

	return errorColor
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
