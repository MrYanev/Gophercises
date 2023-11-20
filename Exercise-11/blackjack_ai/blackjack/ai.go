package blackjack

import (
	"fmt"

	"github.com/MrYanev/Gophercises/Exercise-9/deck_of_cards/deck"
)

type AI interface {
	Bet() int
	Results(hand [][]deck.Card, dealer []deck.Card)
	Play(hand []deck.Card, dealer deck.Card)
}

type HumanAI struct{}

type dealerAI struct{}

func (ai *dealerAI) Bet() int {
	return 1
}

func (ai *dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	} else {
		return MoveStand
	}
}

func (ai *dealerAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	//Noop
}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) {
	for {
		fmt.Println("Player: ", hand)
		fmt.Println("Dealer: ", dealer)
		fmt.Println("What will you do? (h)it, (s)tands?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit()
		case "s":
			return MoveStand()
		default:
			fmt.Println("That's not a valid option!")
		}
	}
}

func (ai *HumanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Dealer:", dealer)
}
