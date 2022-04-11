package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Create a new type of deck that is a slice of strings
//It its like saying that this type deck that has behaviour of slice string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Espadas", "Diamantes", "Corazones", "Trebol"}
	cardValues := []string{"One", "two", "three", "four", "five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//(d deck) is declaring the receiver of that new function.
// d can be any name, d is a reference, d can be think as self or this. GO do not have slef or this
// It is a convention to name as the first one or 2 of the type name
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Return 2 decks, new one and one with missing cards
//2 types deck

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
