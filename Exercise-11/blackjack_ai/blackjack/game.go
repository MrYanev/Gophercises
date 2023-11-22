package blackjack

import (
	"errors"
	"fmt"

	"github.com/MrYanev/Gophercises/Exercise-9/deck_of_cards/deck"
)

const (
	stateBet state = iota
	statePlayerTurn
	stateDealerTurn
	stateHandOver
)

type state int8

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 100
	}
	if opts.BlackjackPayout == 0.0 {
		opts.BlackjackPayout = 1.5
	}
	g.nDecks = opts.Decks
	g.Hands = opts.Hands
	g.blackjackPayout = opts.BlackjackPayout
	return g
}

type Game struct {
	nDecks          int
	Hands           int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []hand
	handIdx   int
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI AI
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player[g.handIdx].cards
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

type hand struct {
	cards []deck.Card
	bet   int
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	g.playerBet = bet
}

func deal(g *Game) {
	playerHand := make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		playerHand = append(playerHand, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.player = []hand{
		{
			cards: playerHand,
			bet:   g.playerBet,
		},
	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) int {
	g.deck = nil
	min := 53 * g.nDecks / 3
	for i := 0; i < g.Hands; i++ {
		shuffled := false
		if len(g.deck) < min {
			g.deck = deck.NewDeck(deck.Deck(g.nDecks), deck.Shuffle)
			shuffled = true
		}

		bet(g, ai, shuffled)
		deal(g)

		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(*g.currentHand()))
			copy(hand, *g.currentHand())
			move := ai.Play(hand, g.dealer[0])
			err := move(g)
			switch err {
			case errBust:
				MoveStand(g)
			case nil:
				//noop
			default:
				panic(err)
			}
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endRound(g, ai)
	}
	return g.balance
}

var (
	errBust = errors.New("Hand score exceeded 21")
)

type Move func(*Game) error

func MoveHit(g *Game) error {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		return errBust
	}
	return nil
}

func MoveStand(g *Game) error {
	if g.state == stateDealerTurn {
		g.state++
		return nil
	}
	if g.state == statePlayerTurn {
		g.handIdx++
		if g.handIdx == len(g.player) {
			g.state++
		}
		return nil
	}
	return errors.New("Invalid state")
}

func MoveSplit(g *Game) error {
	cards := g.currentHand()
	if len(*cards) != 2 {
		return errors.New("You can only slplit two of the same card!")
	}
	if (*cards)[0].Rank != (*cards)[1].Rank {
		return errors.New("Both cards must be the same rank!")
	}
	g.player = append(g.player, hand{
		cards: []deck.Card{(*cards)[1]},
		bet:   g.player[g.handIdx].bet,
	})
	*cards = (*cards)[:1]
	return nil
}

func MoveDouble(g *Game) error {
	if len(g.player) != 2 {
		return errors.New("Can't double a single card hand!")
	}
	g.playerBet *= 2
	MoveHit(g)
	return MoveStand(g)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func endRound(g *Game, ai AI) {
	dScore := Score(g.dealer...)
	dBlackjack := Blackjack(g.dealer...)
	allHands := make([][]deck.Card, len(g.player))
	for hi, hand := range g.player {
		cards := hand.cards
		allHands[hi] = cards
		pScore, pBlackjack := Score(cards...), Blackjack(cards...)
		winnings := hand.bet
		switch {
		case pBlackjack && dBlackjack:
			winnings = 0
		case dBlackjack:
			winnings = -winnings
		case pBlackjack:
			winnings = int(float64(winnings) * g.blackjackPayout)
		case pScore > 21:
			fmt.Println("You busted")
			winnings *= -1
		case dScore > 21:
			fmt.Println("Dealer busted")
		case pScore > dScore:
			fmt.Println("You win!")
		case dScore > pScore:
			fmt.Println("You lose")
			winnings *= -1
		case dScore == pScore:
			fmt.Println("Draw")
			winnings = 0
		}
		g.balance += winnings
	}

	fmt.Println()
	ai.Results(allHands, g.dealer)
	g.player = nil
	g.dealer = nil
}

// Score will take in a hand of cards and return the best blackjack score
// possible with the hand.
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// Soft returns true if the score of a hand is a soft score - that is if an ace
// is being counted as 11 points.
func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
}

func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
