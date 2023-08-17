package main

import (
	"fmt"

	"github.com/MustafaNafizDurukan/Guard/internal/hunters"
)

func main() {
	hunter := hunters.Select("-lnk", "")
	if hunter == nil {
		fmt.Println("The hunter you requested was not found.")
		return
	}

	hunters.Init(hunter)

	hunters.Hunt(hunter)
}
