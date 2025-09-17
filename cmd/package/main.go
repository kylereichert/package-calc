package main

import (
	"fmt"
	"github.com/kylereichert/package-calc/internal/package"
)

func main() {
	fmt.Println("testing")

	test := materials.Materials{
		Rebar10M: 15,
	}

	fmt.Println(test.Rebar10M)
}
