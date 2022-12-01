package main

import (
	"SFA/Lecture17/Task/cardgame"
	"fmt"
)

func main() {

	/* possible card values: 2,3,4,5,6,7,8,9,10,J,D,K,A
	otherwise: panic

	possible card suits: club, heart, diamond, spade (1,2,3,4)
	otherwise: panic
	*/

	CardsSlice := []cardgame.Card{{CardValue: cardgame.J, CardSuit: cardgame.Club},
		{CardValue: 2, CardSuit: cardgame.Diamond},
		{CardValue: 3, CardSuit: cardgame.Heart},
		{CardValue: 5, CardSuit: cardgame.Spade},
		{CardValue: 10, CardSuit: cardgame.Club},
		{CardValue: cardgame.A, CardSuit: cardgame.Heart},
		{CardValue: cardgame.D, CardSuit: cardgame.Diamond},
		{CardValue: cardgame.K, CardSuit: cardgame.Spade},
		{CardValue: 8, CardSuit: cardgame.Spade},
		{CardValue: 4, CardSuit: cardgame.Heart},
		{CardValue: 7, CardSuit: cardgame.Diamond},
		{CardValue: 9, CardSuit: cardgame.Club},
		{CardValue: 6, CardSuit: cardgame.Spade},
		{CardValue: 3, CardSuit: cardgame.Heart},
		{CardValue: cardgame.J, CardSuit: cardgame.Club},
		{CardValue: cardgame.A, CardSuit: cardgame.Club},
	}

	fmt.Println(cardgame.MaxCard(CardsSlice))

}
