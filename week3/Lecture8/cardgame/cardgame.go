package cardgame

import (
	"math/rand"
	"time"
)

type Deck struct {
	Slice []Card
}

type Card struct {
	Value int
	Suit  string
}

func (d *Deck) Deal() *Card {

	if len(d.Slice) == 0 {
		return &Card{}
	}
	currentCard := d.Slice[0]
	d.Slice = d.Slice[1:]

	return &currentCard
}

func random() int {
	rand.Seed(time.Now().Local().UnixMicro())
	return rand.Intn(52)
}

func random2() int {
	rand.Seed(time.Now().Local().UnixMilli())
	return rand.Intn(52)
}

func NewDeck(cv []int, cs []string) *Deck {
	d := Deck{make([]Card, 0, 52)}
	currentCard := Card{}

	for i := 0; i < 13; i++ {
		currentCard.Value = cv[i]
		for k := 0; k < 4; k++ {
			currentCard.Suit = cs[k]
			d.Slice = append(d.Slice, currentCard)
		}
	}
	return &d
}