package colorval

import (
	"flag"
	"fmt"
	"image/color"
)

// A wrapper around color.NRGBA that's also a flag.Value
type NRGBA struct {
	color.NRGBA
}

func (nrgba NRGBA) String() string {

	return (fmt.Sprintf("%x", nrgba.R) +
		fmt.Sprintf("%x", nrgba.G) +
		fmt.Sprintf("%x", nrgba.B) +
		fmt.Sprintf("%x", nrgba.A))
}

func (nrgba *NRGBA) Set(in string) error {

	var err error

	if len(in) == 6 || len(in) == 8 {

		_, err = fmt.Sscanf(in[:2], "%x", &nrgba.R)
		if err != nil {
			return err
		}
		_, err = fmt.Sscanf(in[2:4], "%x", &nrgba.G)
		if err != nil {
			return err
		}
		_, err = fmt.Sscanf(in[4:6], "%x", &nrgba.B)
		if err != nil {
			return err
		}

	} else {

		return fmt.Errorf(
			"A color needs to be 6 or 8 characters long, not %d",
			len(in),
		)
	}

	if len(in) == 8 {

		_, err = fmt.Sscanf(in[6:], "%x", &nrgba.A)
		if err != nil {
			return err
		}

	} else {

		nrgba.A = 255
	}

	return nil
}

var (
	Tree  = NRGBA{color.NRGBA{0, 127, 0, 255}}
	Space = NRGBA{color.NRGBA{127, 255, 50, 255}}
	Fire  = NRGBA{color.NRGBA{255, 0, 0, 255}}
	Ash   = NRGBA{color.NRGBA{127, 127, 127, 255}}
	Error = NRGBA{color.NRGBA{255, 0, 255, 255}}
)

func init() {

	flag.Var(&Tree, "tree-color", "The color of trees")
	flag.Var(&Space, "space-color", "The color of free space")
	flag.Var(&Fire, "fire-color", "The color of fire")
	flag.Var(&Ash, "ash-color", "The color of ashes")
	flag.Var(&Error, "err-color", "Color for errnoerously encoded cells")
}
