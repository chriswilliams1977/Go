package main

import "fmt"

//If there is any other type which has getGreeting and returns a string then is is a member of type bot
//englishBot and spainishbot are therefore also type bot
//interface type. Cannot create a value directly.
type bot interface {
	getGreeting() string
}

//concrete type like map, struct, int, string. You can create a value out of it directly.
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

//english and spainish bot satify the bot interface specifications so they can be passed in
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//Same as other existing function so use an interface for the implementation
//func printGreeting(sb spainishBot) {
//	fmt.Println(sb.getGreeting())
//}

//Ommitted value for englishBot as we are not using it
func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
