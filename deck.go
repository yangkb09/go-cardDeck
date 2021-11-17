package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // return entire ln bc WriteFile returns an err
}

func newDeckFromFile(filename string) deck { // no receiver bc dont have deck yet
	bs, err := ioutil.ReadFile(filename)
	if err != nil { // error handling; log err and quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s) // able to do this bc deck is basically []string
}

func (d deck) shuffle() {
	for i := range d { // don't always need to reference item we're iterating over
		newPosition := rand.Intn(len(d) - 1) // len is how we get length of slice

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
