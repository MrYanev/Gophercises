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

func TestJoker(t *testing.T) {
	cards := NewDeck(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, received:", count)
	}
}
