package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Here we are calling a function with an argument of d and deck and handSize of type of int
// this function below should return to slices.  it should take two slices and split them into two
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// this function will take a deck and return a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// this function will save our data to the machine. first it will take the filenmae,
// then it will take the slice and turn it into a string after that it will be turned to byte.
// the last argument is the premission to write or read the file we are saveing into
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// this function will take a file and read its content
// it will receive the file name and outpot what inside it
// now we will have byte slice that we need to turn into a string se we
// can read it
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and close the applicaion
		fmt.Println("Erorr:", err)
		os.Exit(1)
	}
	// bs => byteSlice. These will rturn a string of this sort
	// Ace of spade, King of tiles. basically it will return a deck
	s := strings.Split(string(bs), ",")
	return deck(s) // this will return a deck!
}

// this function will shuffle our slice and return it
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// we are following convention so instead of using index we are going to use `i`
	// we will always have the same random set of slice since we are always using the
	// same seed, so we will need to truly create a random sequence
	for i := range d {
		newPostion := r.Intn(len(d) - 1)

		d[i], d[newPostion] = d[newPostion], d[i]
	}
}
