package main

import (
	"fmt"
)

// create a new type of deck
// which is a slice of strings
type deck []string

// create a new function that delete
// d is equal to Cards variable. it the cards deck we passed from the main.go folder

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
