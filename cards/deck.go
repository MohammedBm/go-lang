package main

import (
	"fmt"
)

// create a new type of deck
// which is a slice of strings
// now that we have type deck, instead of writing string[] everytime we are writing a function or a new
// variable we can instead write deck
type deck []string

// this function should return a deck. by having the word `deck` after the function name.
// we are telling the function that we are always returning a string from this function
func newDeck() deck {
	cards := deck{}

	// we are declaring this function using string instead of using our word `deck`. We are doing that
	// since this is not an actual deck
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// also instead of writing all the 52 cards, we are going to use a for loop that move throw the
	// slices and assign one value for each of the suits
	// we are using _ instead of `i` or `j`. If we are using them go will give us an error message,
	// because we are not using them. To make sure go undrstand our code we use underscore to tell it that we do not care
	// about the indices. Also We can use them on variables we dont need.
	for _, suit := range cardSuits { // this is just a regular codeing stuff, we dont need any new knowledge to understand this code!!!
		for _, value := range cardValues {
			cards = append(cards, value+" Of "+suit)
		}
	}

	return cards
}

// create a new function that delete
// d is equal to Cards variable. it the cards deck we passed from the main.go folder
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
