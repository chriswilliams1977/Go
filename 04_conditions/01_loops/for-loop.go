package main

import (
	"fmt"
)

func main() {

	i := 0
	//same as for i := 0; i < 10; i++
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//to break out of a loop
	for {
		if i >= 10 {
			fmt.Println(i)
			break
		}
		i++
	}
}
