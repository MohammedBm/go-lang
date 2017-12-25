package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	//both of card variables do the same thing
	// the first line we declared var and what type of vat do we need
	// in the second line we didn't, we let go decide what kind of var do we need
  // we dont need `:= ` if we want to change the value of card, we can do it without the column
	// card = "King of Cards" to reassign a value to card
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}