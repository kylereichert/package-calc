package main

import (
	"fmt"
	"github.com/kylereichert/package-calc/internal/package"
)

func main() {
	fmt.Println("testing")

	test1 := materials.Test{
		Name: "Jimbo",
		Age: 30,
		City: "Cool Zone",
	}

	fmt.Println(test1.Name)
}