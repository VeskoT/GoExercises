package main

import (
	"bufio"
	"fmt"
	"github.com/VeskoT/GoExercises/deck"
	"os"
	"strings"
)

type Hand []deck.Card

type State uint8

const (
	PlayerTurn State = iota
	DealerTurn
	HandOver
)

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ",")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", HIDDEN"
}

func draw(cardDeck []deck.Card) (deck.Card, []deck.Card){
	return cardDeck[0], cardDeck[1:]
}

type GameState struct {
	Deck []deck.Card
	State State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case PlayerTurn:
		return &gs.Player
	case DealerTurn:
		return &gs.Dealer
	default:
		panic("It isn't anyone's turn yet")
	}
}

func clone(gs GameState) GameState {
	clonedGs := GameState {
		Deck: make([]deck.Card, len(gs.Deck)),
		State: gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(clonedGs.Deck, gs.Deck)
	copy(clonedGs.Player, gs.Player)
	copy(clonedGs.Dealer, gs.Dealer)
	return clonedGs
}

func (h Hand)Score() int {
	lesserScore := h.getLesserScore()

	// No chance of there being an ace in our hand
	if lesserScore > 11 {
		return lesserScore
	}

	for _, card := range h {
		if card.Rank == deck.Ace {
			// we have only counted the Ace as a one instead of eleven -> need to add 10 more to the score
			return lesserScore + 10
		}
	}
	return lesserScore
}

// Lesser because Aces are still counted as having a score of 1. Will be handled in the Score() function
func (h Hand)getLesserScore() int {
	score := 0
	for _,card := range h {
		score = score + min(int(card.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	cardDeck := deck.New()
	cardDeck2 := deck.New()
	cardDeck3 := deck.New()

	cards := deck.CombineDecks(cardDeck, cardDeck2, cardDeck3)
	cards = deck.Shuffle(cards)
	ret.Deck = cards
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 4)
	ret.Dealer = make(Hand, 0, 4)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = PlayerTurn
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func EndGame(gs GameState) GameState {
	ret := clone(gs)
	playerScore, dealerScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("Game ending. Final hands:")
	fmt.Println("Player's hand:", ret.Player)
	fmt.Println("Dealer's hand:", ret.Dealer)
	switch {
	case playerScore > 21:
		fmt.Println("Player busted.")
	case dealerScore > 21:
		fmt.Println("Dealer busted.")
	case playerScore > dealerScore:
		fmt.Println("Player won!")
	case dealerScore > playerScore:
		fmt.Println("Dealer won!")
	case playerScore == dealerScore:
		fmt.Println("Draw")
	}
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string
	reader := bufio.NewReader(os.Stdin)
	for gs.State == PlayerTurn {
		fmt.Println("Player's hand:", gs.Player)
		fmt.Println("Dealer's hand", gs.Dealer.DealerString())
		fmt.Print("What do you want to do? (s)tand or (h)old: ")
		input, _ = reader.ReadString('\n')
		input = input[0:len(input)-1]
		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Not a valid option: ", input)
		}
	}

	for gs.State == DealerTurn {
		if gs.Dealer.Score() < 16 || (gs.Dealer.Score() == 17 && gs.Dealer.getLesserScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	gs = EndGame(gs)




}
