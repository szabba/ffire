package main

import (
	"fmt"
	"github.com/szabba/ffire/forest"
	"os"
)

func main() {

	var (
		g   forest.Grid
		err error
	)

	for _, err = fmt.Scan(&g); err == nil; _, err = fmt.Scan(&g) {

		fmt.Print(g)
		fmt.Println()
	}

	fmt.Fprint(os.Stderr, err.Error())
}
