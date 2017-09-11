package main

func main() {

	//cards := []string{"Ace of Spades", newCard()}
	//using deck type
	//cards := newDeck()
	//cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()

}
