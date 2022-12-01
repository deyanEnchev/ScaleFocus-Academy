package main

import (
	"fmt"
	"math/rand"
	"time"
)

type deck struct {
	slice []card
}

type card struct {
	Value int
	Suit  string
}

type CardValue = int

const (
	J CardValue = iota + 11
	D
	K
	A
)

func random() int {
	rand.Seed(time.Now().Local().UnixMicro())
	return rand.Intn(52)
}

func random2() int {
	rand.Seed(time.Now().Local().UnixMilli())
	return rand.Intn(52)
}

func NewDeck(cv []int, cs []string) *deck {
	d := deck{make([]card, 0, 52)}
	currentCard := card{}

	for i := 0; i < 13; i++ {
		currentCard.Value = cv[i]
		for k := 0; k < 4; k++ {
			currentCard.Suit = cs[k]
			d.slice = append(d.slice, currentCard)
		}
	}
	return &d
}

func (d deck) Shuffle() *deck {
	for i := 0; i < 300; i++ {
		randomNumber := random()
		randomNumber2 := random2()
		oldCard := d.slice[randomNumber]

		d.slice[randomNumber] = d.slice[randomNumber2]
		d.slice[randomNumber2] = oldCard
	}

	return &d
}

func (d *deck) Deal() *card {

	if len(d.slice) == 0 {
		return &card{}
	}
	currentCard := d.slice[0]
	d.slice = d.slice[1:]

	return &currentCard
}

func main() {
	cardSuit := []string{"♠", "♥", "♦", "♣"}
	cardValue := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, J, D, K, A}
	d := NewDeck(cardValue, cardSuit)
	d.Shuffle()

	for i := 0; i < 55; i++ {
		fmt.Println(*d.Deal())
	}

}

//The implementation will go like this:

/*
Make a slice "deck" of cards (deck []card)
Put cardSuit and cardValue in every card
Check if the random generated card is equal to some card in the slice
Iterate 52 times to fill the whole deck

Other variant to think about:

Make a map[card]bool
Fill the map with cards
If card already created, then i++ (to keep iterating so the deck can be of 52 cards)


Check if you can compare cards!
*/
