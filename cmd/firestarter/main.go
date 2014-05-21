package main

import (
	"flag"
	"fmt"
	"github.com/szabba/ffire/forest"
	"math/rand"
	"time"
)

var p float64

func init() {

	flag.Float64Var(&p, "p", 0.01, "Probability of lighting a tree")

	flag.Parse()
}

func main() {

	rand.Seed(time.Now().Unix())

	var grid forest.Grid
	fmt.Scan(&grid)

	for i, row := range grid {
		for j, cell := range row {

			if cell == forest.Tree && rand.Float64() < p {

				grid[i][j] = forest.Fire
			}
		}
	}

	fmt.Print(grid)
}
