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
	fmt.Printf("res: %v\n", res)

}
