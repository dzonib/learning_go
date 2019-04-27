package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSign := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, sign := range cardSign {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+sign)
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

	s := []string(d)

	return strings.Join(s, ",")
}

func (d deck) saveToFile(filename string) error {
	// last argument 0666 is promision (anyone can read and write)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		// exit program
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",")

	return deck(s)
}



func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) -1) 
		
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
