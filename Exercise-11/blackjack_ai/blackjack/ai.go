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
			return Hit
		case "s":
			return Stand
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

type GameState struct{}

func Hit(gs GameState) GameState {
	return gs
}

func Stand(gs GameState) GameState {
	return gs
}
