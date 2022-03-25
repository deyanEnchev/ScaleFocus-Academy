package cardgame

import (
	"math/rand"
	"testing"
	"time"
)

var CardsSlice = []Card{{CardValue: J, CardSuit: Club},
{CardValue: 2, CardSuit: Diamond},
{CardValue: 3, CardSuit: Heart},
{CardValue: 5, CardSuit: Spade},
{CardValue: 10, CardSuit: Club},
{CardValue: A, CardSuit: Heart},
{CardValue: D, CardSuit: Diamond},
{CardValue: K, CardSuit: Spade},
{CardValue: 8, CardSuit: Spade},
{CardValue: 4, CardSuit: Heart},
{CardValue: 7, CardSuit: Diamond},
{CardValue: 9, CardSuit: Club},
{CardValue: 6, CardSuit: Spade},
{CardValue: 3, CardSuit: Heart},
{CardValue: J, CardSuit: Club},
{CardValue: A, CardSuit: Club},
}
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(14)
}

func TestCompareCards(t *testing.T) {
	playerOne, playerTwo := Card{CardValue: 2, CardSuit: Club},
		Card{CardValue: K, CardSuit: Diamond}

	result := CompareCards(playerOne, playerTwo)

	if result != -1 {
		t.Errorf("Expected -1 got %d", result)
	}

	result = CompareCards(playerOne, playerOne)
	if result != 0 {
		t.Errorf("Expected 0 got %d", result)

	}

	result = CompareCards(playerTwo, playerOne)
	if result != 1 {
		t.Errorf("Expected 1 got %d", result)

	}
}

func TestMaxCard(t *testing.T) {
	result := MaxCard(CardsSlice)
	biggestCard := Card{CardValue: K, CardSuit: Spade}
	if result != biggestCard {
		t.Errorf("Expected K Spade got %v", result)
	}
}

//My tests are shit for sure, but will watch some courses and will improve.