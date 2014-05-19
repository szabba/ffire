package main

import (
	"fmt"
	"github.com/szabba/ffire/forest"
	"image"
	"image/color"
	"image/png"
	"os"
)

func cellColor(cell forest.Cell) color.Color {

	if cell == forest.Tree {

		return color.NRGBA{0, 127, 0, 255}
	} else if cell == forest.Space {
		return color.NRGBA{127, 255, 50, 255}
	} else if cell == forest.Ash {
		return color.NRGBA{0, 0, 0, 255}
	} else if cell == forest.Fire {
		return color.NRGBA{255, 0, 0, 255}
	}

	return color.NRGBA{255, 0, 255, 255}
}

const nameFormat = "forest%07d.png"

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
