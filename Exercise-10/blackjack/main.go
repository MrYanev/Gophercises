package main

import (
	"fmt"

	"github.com/MrYanev/Gophercises/Exercise-11/blackjack_ai/blackjack"
	"github.com/MrYanev/Gophercises/Exercise-9/deck_of_cards/deck"
)

type basicAI struct {
	count int
	seen  int
	decks int
}

func (ai *basicAI) Bet(shuffled bool) int {
	if shuffled {
		ai.count = 0
		ai.seen = 0
	}
	trueScore := ai.count / ((ai.decks*52 - ai.seen) / 52)
	switch {
	case trueScore >= 14:
		return 100000
	case trueScore >= 8:
		return 5000
	default:
		return 100
	}
}

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	score := blackjack.Score(hand...)
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := blackjack.Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				return blackjack.MoveSplit
			}
		}
		if (score == 10 || score == 11) && !blackjack.Soft(hand...) {
			return blackjack.MoveDouble
		}
	}
	dScore := blackjack.Score(dealer)
	if dScore >= 5 && dScore <= 6 {
		return blackjack.MoveStand
	}
	if score < 13 {
		return blackjack.MoveHit
	}
	return blackjack.MoveStand
}

func (ai *basicAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	for _, card := range dealer {
		ai.Count(card)
	}
	for _, hand := range hands {
		for _, card := range hand {
			ai.Count(card)
		}
	}
}

func (ai *basicAI) Count(card deck.Card) {
	score := blackjack.Score(card)
	switch {
	case score >= 10:
		ai.count--
	case score <= 6:
		ai.count++
	}
	ai.seen++
}

func main() {
	opts := blackjack.Options{
		Decks:           3,
		Hands:           2,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(&basicAI{
		decks: 4,
	})
	fmt.Println(winnings)
}
