package cardgame

type CardSuit = int

const (
	Club CardSuit = iota + 1
	Heart
	Diamond
	Spade
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

func CompareCards(playerOne, playerTwo Card) int {

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

func MaxCard(cards []Card) Card {
	var biggestCard Card

	for i := 0; i < len(cards)-1; i++ {
		if CompareCards(cards[i], biggestCard) == 1 {
			biggestCard = cards[i]
		}
	}

	return biggestCard
}
