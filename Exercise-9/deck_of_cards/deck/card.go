package deck

type Card struct {
}

// A function to Shuffle the deck
func (c *Card) Shuffle([]Card) ([]Card, error) {
	return nil, nil
}

// A function to generate new deck of cards
func NewDeck() []Card {
	return nil
}

// A function to remove specific type of cards
func (c *Card) SortOut(deck []Card, suit string) ([]Card, error) {
	return nil, nil
}

// A function to build deck from multiple decks
func (c *Card) MultiDeck(count int) []Card {
	return nil
}
