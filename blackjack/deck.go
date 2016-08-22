package blackjack

type Deck struct {
	cardList []*Card
}

type Card struct {
	suit  string
	data  string
	value int
}
