package blackjack

type Game struct {
	deck   Deck
	player Player
	dealer Dealer
}

func (game *Game) InitCards() {

	var cardList []*Card
	suitesList := []string{"clubs", "hearts", "diamonds", "spades"}
	contentList := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "A", "J", "K", "Q"}
	facevalueMap := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "A": 11, "J": 10, "K": 10, "Q": 10}
	for _, suite := range suitesList {
		for _, content := range contentList {
			cardList = append(cardList, &Card{suit: suite, data: content, value: facevalueMap[content]})

		}

	}
	deck := Deck{cardList: cardList}
	game.deck = deck

}

func (game *Game) PlaceFourCards() ([]*Card, []*Card) {

	cardList := game.deck.cardList

	playerCardList := cardList[0:2]
	dealerCardList := cardList[2:4]
	newcardList := cardList[4:len(cardList)]
	game.deck.cardList = newcardList
	return playerCardList, dealerCardList

}

type GameController interface {
	placeCard()
	calculateScore()
	checkWin()
}
