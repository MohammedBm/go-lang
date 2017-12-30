package main

func main() {
	// go is not a Object Oriatned language, so there is no classes on Go.

	// variables
	/*
			// var card string = "Ace of Spades"
			// card := "Ace of Spades"
			//both of card variables do the same thing
			// the first line we declared var and what type of vat do we need
			// in the second line we didn't, we let go decide what kind of var do we need
		  // we dont need `:= ` if we want to change the value of card, we can do it without the column
			// card = "King of Cards" to reassign a value to card
			card := newCard()

			fmt.Println(card)
	*/

	//arrays and slice
	/*
		arrays can only be of one fixed type, for example they can only be array of strings
		or an array of integer. You cannot have an array of both types. also array is a fixed length
		on the other hand, slice can have a different length.complex128
	*/
	// now we will use the newDeck function that we created on the deck.go file
	cards := newDeck()

	// hand, remaningCards := deal(cards, 5)

	// cards.print()
	// hand.print()
	// remaningCards.print()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}
