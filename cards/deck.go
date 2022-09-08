package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	d := deck{}

	varSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	varNums := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range varSuits {
		for _, num := range varNums {
			d = append(d, num+" of "+suit)
		}
	}

	return d
}
