package main

import (
	"fmt"

	"github.com/radoslavboychev/gophercises-link/parser"
)

var filename string = "../.././src/ex3.html"

func main() {
	res, err := parser.ParseHTML(filename)
	if err != nil {
		return
	}

	fmt.Printf("res: %v\n", res)
}
