package main

import "fmt"

type CardSuit = int

const (
	club CardSuit = iota + 1
	heart
	diamond
	spade
)

type CardValue = int

const (
	J CardValue = iota + 11
	D
	K
	A
)

type Card struct {
	Value int
	Suit  int
}



func main() {

	/* possible card values: 2,3,4,5,6,7,8,9,10,J,D,K,A
	otherwise: panic

	possible card suits: club, heart, diamond, spade (1,2,3,4)
	otherwise: panic
	*/

	cardsSlice := []Card{Card{Value: J, Suit: club},
		Card{Value: 2, Suit: diamond},
		Card{Value: 3, Suit: heart},
		Card{Value: 5, Suit: spade},
		Card{Value: 10, Suit: club},
		Card{Value: A, Suit: heart},
		Card{Value: D, Suit: diamond},
		Card{Value: K, Suit: spade},
		Card{Value: 8, Suit: spade},
		Card{Value: 4, Suit: heart},
		Card{Value: 7, Suit: diamond},
		Card{Value: 9, Suit: club},
		Card{Value: 6, Suit: spade},
		Card{Value: 3, Suit: heart},
		Card{Value: J, Suit: club},
		Card{Value: A, Suit: club},
	}

	var temp compareCardsV2
	

	cc := &temp

	printProperly(maxCardV2(cardsSlice, *cc))

}


type compareCardsV2 func(playerOne, playerTwo Card) int


func (cc *compareCardsV2) compareCardsV2 (playerOne, playerTwo Card) int {
	
	playerOneTotal := playerOne.Value + playerOne.Suit
	playerTwoTotal := playerTwo.Value + playerTwo.Suit
	var result int

	if playerOneTotal < playerTwoTotal {
		result = -1
	}

	if playerOneTotal == playerTwoTotal {
		result = 0
	}

	if playerOneTotal > playerTwoTotal {
		result = 1
	}

	return result
}


func maxCardV2(cards []Card, comparatorFunc compareCardsV2) Card {

	var biggestCard Card

	for i := 0; i < len(cards)-1; i++ {
		if comparatorFunc.compareCardsV2(cards[i], biggestCard) == 1 {
			biggestCard = cards[i]
		}
	}

	return biggestCard
}

func printProperly(card Card) {
	if card.Value >= 2 && card.Value <= 10 {
		fmt.Print(card.Value, " ")
	} else {
		switch card.Value {
		case 11:
			fmt.Print("J", " ")
		case 12:
			fmt.Print("D", " ")
		case 13:
			fmt.Print("K", " ")
		case 14:
			fmt.Print("A", " ")
		}
	}

	switch card.Suit {
	case 1:
		fmt.Print("club")
	case 2:
		fmt.Print("heart")
	case 3:
		fmt.Print("diamond")
	case 4:
		fmt.Print("spade")
	}
}
