package deck

import "testing"

func TestNew(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 13*4 {
		t.Error("Wrong amount of cards")
	}
}
