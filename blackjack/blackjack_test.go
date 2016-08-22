package blackjack

import "testing"

func TestBlackJackSetup(t *testing.T) {

	game := Game{}

	game.InitCards() // deck of cards and shuffle
	result := len(game.deck.cardList)

	if result != 52 {
		t.Errorf("Expected cardList to be 52, got %d", result)
	}
}

func TestPlaceFourCards(t *testing.T) {

	game := Game{}

	game.InitCards() // deck of cards and shuffle
	initialLength := len(game.deck.cardList)

	playerCardList, dealerCardList := game.PlaceFourCards()

	result := len(game.deck.cardList)
	if initialLength-4 != len(game.deck.cardList) {
		t.Errorf("Expected cardList to be 48, got %d", result)
	}
	result = len(playerCardList)
	if 2 != len(playerCardList) {
		t.Errorf("Expected palyer cardList to be 2, got %d", result)
	}
	result = len(dealerCardList)
	if 2 != len(dealerCardList) {
		t.Errorf("Expected dealer cardList to be 2, got %d", result)
	}

}

func TestCalculateScore(t *testing.T) {
	game := Game{}

	game.InitCards() // deck of cards and shuffle

	playerCardList, dealerCardList := game.PlaceFourCards()

	var playerScore int
	var dealerScore int
	for _, card := range playerCardList {
		playerScore += card.value
	}
	for _, card := range dealerCardList {
		dealerScore += card.value
	}
	if playerScore != 1 {
		t.Errorf("Expected cardList to be 1, got %d", playerScore)
	}
	if dealerScore != 5 {
		t.Errorf("Expected cardList to be 5, got %d", dealerScore)
	}

}
