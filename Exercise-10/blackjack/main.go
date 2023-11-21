package main

import (
	"fmt"

	"github.com/MrYanev/Gophercises/Exercise-11/blackjack_ai/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
