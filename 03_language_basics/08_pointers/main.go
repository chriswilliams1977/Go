package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	//*int is an int pointer
	var b *int = &a

	fmt.Println(b)
}
