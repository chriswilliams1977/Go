package main

import (
	"fmt"
	"github.com/chriswilliams1977/03_language_basics/03_scope/vis"
)

//whole package scope
var X int = 42

func main() {
	fmt.Println(X)
	foo()

	//only accessible inside function
	y := 17
	fmt.Println(y)
	fmt.Println(vis.MyName)
	vis.PrintVar()
}

func foo() {
	fmt.Println(X)
}
