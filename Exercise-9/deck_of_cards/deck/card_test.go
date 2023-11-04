package deck

import "testing"

func TestNew(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 13*4 {
		t.Error("Wrong amount of cards")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := NewDeck(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades, received card:", cards[0])
	}
}
