package carddraw

import (
	"SFA/week3/Lecture9/cardgame"
)

type dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	currentSlice := make([]cardgame.Card, 0, 52)
	currentCard, err := dealer.Deal()
	for currentCard.Value != 0 {
		if err != nil {
			if dealer.Done() {
				return currentSlice, nil
			} else {
				return nil, err
			}
		}
		currentSlice = append(currentSlice, *currentCard)
		currentCard, err = dealer.Deal()
	}

	return currentSlice, err
}
