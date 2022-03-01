package Task1

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
	CardValue int
	CardSuit int
}

func main() {

	/* possible card values: 2,3,4,5,6,7,8,9,10,J,D,K,A
	otherwise: panic
	
	possible card suits: club, heart, diamond, spade (1,2,3,4)
	otherwise: panic
	*/

	playerOne := Card{CardValue: J, CardSuit: club}
	playerTwo := Card{CardValue: A, CardSuit: spade}

	cardCheker(playerOne, playerTwo)
	fmt.Println(compareCards(playerOne,playerTwo))
}

func compareCards(playerOne,playerTwo Card) int {

	playerOneTotal := playerOne.CardValue + playerOne.CardSuit
	playerTwoTotal := playerTwo.CardValue + playerTwo.CardSuit
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

func cardCheker(playerOne, playerTwo Card) {

	if playerOne.CardValue < 2 || playerOne.CardValue > A || playerTwo.CardValue < 2 || playerTwo.CardValue > A {
		panic("Wrong card value!")
	}

	if playerOne.CardSuit < club || playerOne.CardSuit > spade || playerTwo.CardSuit < club || playerTwo.CardSuit > spade {
		panic("Wrong card suit!")
	}

}
