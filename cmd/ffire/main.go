package main

import (
	"fmt"
	"os"
)

// Kind of content a cell can be filled with
type Content uint

const (
	Tree Content = iota
	Space
	Fire
	Ash

	T, S, F, A = 'T', 'S', 'F', 'A'
)

// An automaton cell
type Cell struct {
	Content Content
	Phase   int
}

func (c *Cell) Scan(state fmt.ScanState, verb rune) error {

	//if verb != 'q' {

	//	return fmt.Errorf(
	//		"'%c' is not a verb for Cells",
	//		verb,
	//	)
	//}

	state.SkipSpace()

	r, _, err := state.ReadRune()
	if err != nil {

		return err

	} else if r == T {

		c.Content = Tree

	} else if r == S {

		c.Content = Space

	} else if r == F {

		c.Content = Fire

	} else if r == A {

		c.Content = Ash

	} else {

		state.UnreadRune()
		return fmt.Errorf(
			"The rune '%c' does not represent a cell's content.",
			r,
		)
	}

	_, err = fmt.Fscan(state, &c.Phase)

	return err
}

func (c Cell) String() string {

	var r rune

	if c.Content == Tree {
		r = T
	} else if c.Content == Fire {
		r = F
	} else if c.Content == Space {
		r = S
	} else if c.Content == Ash {
		r = A
	}

	return fmt.Sprintf("%c %d", r, c.Phase)
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

// A neighbourhood
type Neighbourhood interface {
	Size() int
	For(g Grid, i, j int, ns []Cell)
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

func (m Moore) For(g Grid, i, j int, ns []Cell) {

	w, h := g.Size()

	k := 0
	for p := -1; p < 2; p++ {
		for q := -1; q < 2; q++ {

			if !(q == 0 && p == 0) {

				r := wrap(i, p, w)
				s := wrap(j, q, h)

				fmt.Fprintln(
					os.Stderr,
					"\t", r, s,
				)

				ns[k] = g[r][s]
				k++
			}
		}
	}
}

// An automaton
type Automaton struct {
	next, now     Grid
	neighbourhood Neighbourhood
}

func NewAutomaton(g Grid, n Neighbourhood) (a *Automaton) {

	a = new(Automaton)
	a.now = g
	a.next = g.Copy()
	a.neighbourhood = n

	return
}

func (auto *Automaton) Step(step func(Cell, []Cell) Cell) {

	w, h := auto.now.Size()
	ns := make([]Cell, auto.neighbourhood.Size())

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
	step func(c Cell, ns []Cell) Cell,
	each func(then, now Grid),
) {

	for i := 0; i < steps; i++ {

		auto.Step(step)
		each(auto.next, auto.now)
	}
}

func main() {

	var (
		m    Moore
		g    Grid
		i, j int
	)
	cs := make([]Cell, m.Size())

	fmt.Scan(&g, &i, &j)
	fmt.Println(g)

	m.For(g, i, j, cs)
	fmt.Println(cs)
}
