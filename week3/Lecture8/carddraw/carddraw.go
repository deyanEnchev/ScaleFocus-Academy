package carddraw

import "SFA/week3/Lecture8/cardgame"

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	currentSlice := make([]cardgame.Card, 0, 52)
	currentCard := dealer.Deal()
	for currentCard.Value != 0 {
		currentSlice = append(currentSlice, *currentCard)
		currentCard = dealer.Deal()
	}

	return currentSlice
}