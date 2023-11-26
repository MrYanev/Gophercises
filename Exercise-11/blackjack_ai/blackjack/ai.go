package blackjack

import (
	"fmt"

	"github.com/MrYanev/Gophercises/Exercise-9/deck_of_cards/deck"
)

type AI interface {
	Bet(shuffled bool) int
	Results(hands [][]deck.Card, dealer []deck.Card)
	Play(hand []deck.Card, dealer deck.Card) Move
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffled bool) int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	//Noop
	///
}

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled!")
	}
	fmt.Println("What's your Bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player: ", hand)
		fmt.Println("Dealer: ", dealer)
		fmt.Println("What will you do? (h)it, (s)tands?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "p":
			return MoveSplit
		default:
			fmt.Println("That's not a valid option!")
		}
	}
}

func (ai humanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Dealer:", dealer)
}
