package forest

import (
	"fmt"
)

// Automaton cell
type Cell uint

const (
	Tree Cell = iota
	Space
	Fire
	Ready
	Ash

	T, S, F, R, A = 'T', 'S', 'F', 'R', 'A'
)

func (c *Cell) Scan(state fmt.ScanState, verb rune) error {

	state.SkipSpace()

	r, _, err := state.ReadRune()
	if err != nil {

		return err

	} else if r == T {

		*c = Tree

	} else if r == S {

		*c = Space

	} else if r == F {

		*c = Fire

	} else if r == R {

		*c = Ready

	} else if r == A {

		*c = Ash

	} else {

		state.UnreadRune()
		return fmt.Errorf(
			"The rune '%c' does not represent a cell's content.",
			r,
		)
	}

	return nil
}

func (c Cell) String() string {

	var r rune

	if c == Tree {
		r = T
	} else if c == Fire {
		r = F
	} else if c == Space {
		r = S
	} else if c == Ash {
		r = A
	} else if c == Ready {
		r = R
	}

	return fmt.Sprintf("%c", r)
}

// A square grid of cells
type Grid [][]Cell

func NewGrid(width, height int) Grid {

	g := make([][]Cell, height)

	for i, _ := range g {

		g[i] = make([]Cell, width)
	}

	return Grid(g)
}

func (g Grid) Size() (width, height int) {

	height = len(g)
	width = len(g[0])
	return
}

func (g Grid) Copy() (g_ Grid) {

	w, h := g.Size()

	g_ = NewGrid(w, h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			g_[i][j] = g[i][j]
		}
	}

	return
}

func (g *Grid) Scan(state fmt.ScanState, verb rune) error {

	var w, h int

	_, err := fmt.Fscan(state, &w)
	defer func() {
		if err != nil {
			g = nil
		}
	}()
	if err != nil {
		return err
	}

	_, err = fmt.Fscan(state, &h)
	if err != nil {
		return err
	}

	*g = NewGrid(w, h)

	var c Cell
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			_, err = fmt.Fscan(state, &c)
			if err != nil {
				return err
			}

			(*g)[i][j] = c
		}
	}

	return err
}

func (g Grid) Format(state fmt.State, c rune) {

	w, h := g.Size()
	fmt.Fprintf(state, "%d %d\n", w, h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			fmt.Fprintf(state, "%s", g[i][j])

			if j+1 == w {
				fmt.Fprintln(state)

			} else {
				fmt.Fprint(state, " ")
			}
		}
	}
}
