package main

import "fmt"

//const (
//	A = iota
//	B
//	C
//)
//
//const (
//	D = iota
//	E
//	F
//)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

func main() {
	const q = 42

	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(C)
	//
	//fmt.Println(D)
	//fmt.Println(E)
	//fmt.Println(F)

	fmt.Println(KB)
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)

}
