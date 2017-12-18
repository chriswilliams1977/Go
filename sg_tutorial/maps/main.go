package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//use this is you want to later define key values
	//var colours map[string]string
	//make is built in function, same as using var
	//colours := make(map[string]string)

	//colours["white"] = "#ffffff"
	//with structs you can get a value using structName.white with amps you use colours[]
	//can also use colours[2] = "#ffffff", all key must be same type and all values same type

	//to delete pass in map and key you want to delete
	//delete(colours,"white")

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for ", colour, " is ", hex)
	}
}


//Difference to structs
//All key and values are of same types, struct you can kix and match
//All key are indexed so you can interate over a map, you cannot do this for a struct as not indexed
//Map is a reference type, struct is a value type
//Reference type, when you pass a map you pass a reference to the data structure for example printMap
//So if you change colours in printMap, colours will ses that change (you copy a reference to the map
//With structs when you pass them to a function you make a copy of the entire struct, so if it changes then the copy changes not the original
//Thus why you need to use pointers to modufy original struct

//Use map when dealing with related properties like string of colours and you dont need to know all keys are compile time (add and remove later)
//Struct you have to define all properties at compile time. Closed set of keys like a person
//Structs used more