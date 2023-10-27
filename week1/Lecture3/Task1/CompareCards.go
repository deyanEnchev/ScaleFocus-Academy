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

func main() {

	/* possible card values: 2,3,4,5,6,7,8,9,10,J,D,K,A
	otherwise: panic

	possible card suits: club, heart, diamond, spade (1,2,3,4)
	otherwise: panic
	*/
	var p1v, p2v int = 2, A
	var p1s, p2s int = club, spade

	cardCheker(p1v, p1s, p2v, p2s)
	fmt.Println(compareCards(p1v, p1s, p2v, p2s))
}

func compareCards(cardOneVal int,
	cardOneSuit int,
	cardTwoVal int,
	cardTwoSuit int) int {

	playerOneTotal := cardOneVal + cardOneSuit
	playerTwoTotal := cardTwoVal + cardTwoSuit
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

func cardCheker(p1v, p1s, p2v, p2s int) {

	if p1v < 2 || p1v > A || p2v < 2 || p2v > A {
		panic("Wrong card value!")
	}

	if p1s < club || p1s > spade || p2s < club || p2s > spade {
		panic("Wrong card suit!")
	}
}
