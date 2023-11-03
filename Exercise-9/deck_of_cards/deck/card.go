package deck

type Suit uint8

const (
	Spade   Suit = iota
	Diamond Suit = iota
	Club    Suit = iota
	Heart   Suit = iota
	Joker   Suit = iota
)

type Rank uint8

const (
	_     Rank = iota
	Ace   Rank = iota
	Two   Rank = iota
	Three Rank = iota
	Four  Rank = iota
	Five  Rank = iota
	Six   Rank = iota
	Seven Rank = iota
	Eight Rank = iota
	Nine  Rank = iota
	Ten   Rank = iota
	Jack  Rank = iota
	Queen Rank = iota
	King  Rank = iota
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return ""
}

// A function to Shuffle the deck
func (c Card) Shuffle([]Card) ([]Card, error) {
	return nil, nil
}

// A function to generate new deck of cards
func (c Card) NewDeck() []Card {
	return nil
}

// A function to remove specific type of cards
func (c Card) SortOut(deck []Card, suit string) ([]Card, error) {
	return nil, nil
}

// A function to build deck from multiple decks
func (c Card) MultiDeck(count int) []Card {
	return nil
}
