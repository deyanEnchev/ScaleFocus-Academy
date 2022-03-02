package Task2

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
	CardSuit  int
}

func main() {

	/* possible card values: 2,3,4,5,6,7,8,9,10,J,D,K,A
	otherwise: panic

	possible card suits: club, heart, diamond, spade (1,2,3,4)
	otherwise: panic
	*/
	
	cardsSlice := []Card{ Card{CardValue: J, CardSuit: club},
	Card{CardValue: 2, CardSuit: diamond},
	Card{CardValue: 3, CardSuit: heart},
	Card{CardValue: 5, CardSuit: spade},
	Card{CardValue: 10, CardSuit: club},
	Card{CardValue: A, CardSuit: heart},
	Card{CardValue: D, CardSuit: diamond},
	Card{CardValue: K, CardSuit: spade},
	Card{CardValue: 8, CardSuit: spade},
	Card{CardValue: 4, CardSuit: heart},
	Card{CardValue: 7, CardSuit: diamond},
	Card{CardValue: 9, CardSuit: club},
	Card{CardValue: 6, CardSuit: spade},
	Card{CardValue: 3, CardSuit: heart},
	Card{CardValue: J, CardSuit: club},
	Card{CardValue: A, CardSuit: club},
 }
	

	printProperly(maxCard(cardsSlice))

	
}

func compareCards(playerOne, playerTwo Card) int {

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



func maxCard (cards []Card) Card {
	var biggestCard Card;
	
	for i := 0; i < len(cards) - 1; i++ {
		if compareCards(cards[i],biggestCard) == 1 {
			biggestCard = cards[i]
		}
	}

	return biggestCard
}

func printProperly (card Card) {
	if card.CardValue >= 2 && card.CardValue <= 10 {
		fmt.Print(card.CardValue, " ")
	} else {
		switch card.CardValue {
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

	switch card.CardSuit {
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
