package main

import (
	"fmt"

	"github.com/radoslavboychev/gophercises-link/parser"
)

func main() {
	res, err := parser.ParseHTML(".././src/ex1.html")
	if err != nil {
		return
	}

	for _, v := range res {
		fmt.Printf("v: %v\n", v)
	}

}
