package main

import (
	"os"
	"testing"
)

// this is our test file where we test our code if it works as intended
func TestNewDeck(t *testing.T) {
	d := newDeck()
	// this will test if the deck total number is 16,
	// if it's less it will throw an error
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// this code will check if the first card genrated is
	// 'Ace of Sapdes'
	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected first card of Ace of Spades %v", d[0])
	}

	// this code will test if the last card is 'Four Of Clubs'
	if d[len(d)-1] != "Four Of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
