package main

import "fmt"

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
