package colorval

import (
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

func (nrgba *NRGBA) Set(in string) {

	if len(in) == 6 || len(in) == 8 {

		fmt.Sscanf(in[:2], "%x", &nrgba.R)
		fmt.Sscanf(in[2:4], "%x", &nrgba.G)
		fmt.Sscanf(in[4:6], "%x", &nrgba.B)
	}

	if len(in) == 8 {

		fmt.Sscanf(in[6:], "%x", &nrgba.A)

	} else {

		nrgba.A = 255
	}

}
