package main

import (
	"SFA/week3/Lecture9/carddraw"
	"SFA/week3/Lecture9/cardgame"
	"fmt"
	"log"
)



type CardValue = int

const (
	J CardValue = iota + 11
	D
	K
	A
)

func main() {
	cardSuit := []string{"♠", "♥", "♦", "♣"}
	cardValue := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, J, D, K, A}
	d := cardgame.NewDeck(cardValue, cardSuit)
	d.Shuffle()

	slice, err := carddraw.DrawAllCards(d)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(slice)

}