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
