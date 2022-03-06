package main

import (
	"SFA/week3/Lecture8/carddraw"
	"SFA/week3/Lecture8/cardgame"
	"fmt"
)



type CardValue = int

const (
	J CardValue = iota + 11
	D
	K
	A
)



// func (d Deck) Shuffle() *Deck {
// 	for i := 0; i < 300; i++ {
// 		randomNumber := random()
// 		randomNumber2 := random2()
// 		oldCard := d.slice[randomNumber]

// 		d.slice[randomNumber] = d.slice[randomNumber2]
// 		d.slice[randomNumber2] = oldCard
// 	}

// 	return &d
// }



func main() {
	cardSuit := []string{"♠", "♥", "♦", "♣"}
	cardValue := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, J, D, K, A}
	d := cardgame.NewDeck(cardValue, cardSuit)
	
	slice := carddraw.DrawAllCards(d)


	fmt.Println(slice)

}
