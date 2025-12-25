package main

import (
	"fmt"

	"github.com/samber/lo"
)

type X struct {
}

func main() {

	var x X

	var x2 *X = new(X)

	var emptyString = ""
	var notEmptyString = "123"

	fmt.Println(emptyString, lo.IsEmpty(emptyString))
	fmt.Println(x, lo.IsEmpty(x))
	fmt.Println(notEmptyString, lo.IsEmpty("123"))
	fmt.Println(x2, lo.IsNotEmpty(&x2))
}