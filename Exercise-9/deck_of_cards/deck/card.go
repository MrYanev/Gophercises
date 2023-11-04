package deck

import (
	"fmt"
	"sort"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Joker
	Heart
)

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Card struct {
	Suit
	Rank
}

func (r Rank) String() string {
	return ""
}

func (s Suit) String() string {
	return ""
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// A function to Shuffle the deck
func (c Card) Shuffle([]Card) ([]Card, error) {
	return nil, nil
}

// A function to generate new deck of cards
func NewDeck(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// A function to check whether a value at index is less than
// a value at another index
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// A function to remove specific type of cards
func (c Card) SortOut(deck []Card, suit string) ([]Card, error) {
	return nil, nil
}

// A function to build deck from multiple decks
func (c Card) MultiDeck(count int) []Card {
	return nil
}
