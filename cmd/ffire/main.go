package main

import (
	"fmt"
	"github.com/szabba/ffire/forest"
	"math/rand"
	"time"
)

// A neighbourhood
type Neighbourhood interface {
	Size() int
	For(g forest.Grid, i, j int, ns []forest.Cell)
}

// The Moore neighbourhood
type Moore struct{}

func (_ Moore) Size() int {
	return 8
}

func wrap(from, delta, width int) int {

	max := width - 1

	for from+delta > max {
		delta -= width
	}

	for from+delta < 0 {
		delta += width
	}

	return from + delta
}

func (m Moore) For(g forest.Grid, i, j int, ns []forest.Cell) {

	w, h := g.Size()

	k := 0
	for p := -1; p < 2; p++ {
		for q := -1; q < 2; q++ {

			if !(q == 0 && p == 0) {

				r := wrap(i, p, h)
				s := wrap(j, q, w)

				ns[k] = g[r][s]
				k++
			}
		}
	}
}

// An automaton
type Automaton struct {
	next, now     forest.Grid
	neighbourhood Neighbourhood
}

func NewAutomaton(g forest.Grid, n Neighbourhood) (a *Automaton) {

	a = new(Automaton)
	a.now = g
	a.next = g.Copy()
	a.neighbourhood = n

	return
}

func (auto *Automaton) Step(step func(forest.Cell, []forest.Cell) forest.Cell) {

	w, h := auto.now.Size()
	ns := make([]forest.Cell, auto.neighbourhood.Size())

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			auto.neighbourhood.For(auto.now, i, j, ns)

			auto.next[i][j] = step(
				auto.now[i][j], ns,
			)
		}
	}

	auto.next, auto.now = auto.now, auto.next
}

func (auto *Automaton) Run(
	steps int,
	step func(c forest.Cell, ns []forest.Cell) forest.Cell,
	each func(then, now forest.Grid),
) {

	for i := 0; i < steps; i++ {

		auto.Step(step)
		each(auto.next, auto.now)
	}
}

func Hell(c forest.Cell, ns []forest.Cell) forest.Cell {

	return forest.Fire
}

func Spread(c forest.Cell, ns []forest.Cell) forest.Cell {

	if c == forest.Space || c == forest.Ash {
		return c
	}

	fires := 0
	if c == forest.Fire {

		fires++
	}
	for _, n := range ns {
		if n == forest.Fire {

			fires++
		}
	}

	if c == forest.Tree {

		if fires > 0 {

			return forest.Fire
		}
		return forest.Tree
	}

	if fires > 6 {

		return forest.Ash
	}
	return forest.Fire
}

func SetFireToTheRain() func(forest.Cell, []forest.Cell) forest.Cell {

	rng := rand.New(rand.NewSource(time.Now().Unix()))

	return func(c forest.Cell, ns []forest.Cell) forest.Cell {

		fires := 0
		for _, n := range ns {
			if n == forest.Fire {
				fires++
			}
		}

		p := float64(fires) / 8

		if fires == 0 || c == forest.Space || c == forest.Ash {

			return c

		} else if c == forest.Fire {

			return forest.Ash

		} else {

			r := rng.Float64()

			if r <= p {

				return forest.Fire
			}

			return c
		}
	}
}

func main() {

	var (
		steps = 100
		g     forest.Grid
	)
	fmt.Scan(&g)

	auto := NewAutomaton(g, Moore{})

	run := 0
	auto.Run(
		steps,
		SetFireToTheRain(),
		func(then, now forest.Grid) {

			run++

			fmt.Print(then)
			fmt.Println()

			fmt.Scanf("\n")

			if run == steps {
				fmt.Print(now)
			}
		},
	)
}
