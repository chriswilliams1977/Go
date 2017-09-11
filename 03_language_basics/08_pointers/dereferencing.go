package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	//*b get the value from that memory address
	fmt.Println(*b)

	//make the value at the pointer address 42
	*b = 42

	fmt.Println(a)
}
