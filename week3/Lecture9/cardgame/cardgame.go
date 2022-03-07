package cardgame

import (
	"errors"
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

func (d *Deck) Done() bool {
	return len(d.Slice) == 0
}

func (d *Deck) Deal() (*Card, error) {
	
	var err error = nil

	if len(d.Slice) == 0 {
		err = errors.New("something went wrong in Deal()")
		return &Card{}, err
	}

	currentCard := d.Slice[0]
	d.Slice = d.Slice[1:]

	return &currentCard, err
}

func random() int {
	rand.Seed(time.Now().Local().UnixMicro())
	return rand.Intn(52)
}

func random2() int {
	rand.Seed(time.Now().Local().UnixMilli())
	return rand.Intn(52)
}

func (d Deck) Shuffle() *Deck {
	for i := 0; i < 300; i++ {
		randomNumber := random()
		randomNumber2 := random2()
		oldCard := d.Slice[randomNumber]

		d.Slice[randomNumber] = d.Slice[randomNumber2]
		d.Slice[randomNumber2] = oldCard
	}

	return &d
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
