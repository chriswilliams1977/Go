package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T)  {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got: %v", len(d))
	}

	if d[0] !=  "Ace of Spades" {
		t.Errorf("Exepected Ace of Spades but got %v ", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v ",d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T)  {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	LoadedDeck := newDeckFromFile("_decktesting")

	if len(LoadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but got %v: ", len(LoadedDeck))
	}

	os.Remove("_decktesting")
}
