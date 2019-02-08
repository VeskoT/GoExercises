package main

import (
	"bufio"
	"fmt"
	"github.com/VeskoT/GoExercises/deck"
	"os"
	"strings"
)

type Hand []deck.Card

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

func main() {
	cardDeck := deck.New()
	cardDeck2 := deck.New()
	cardDeck3 := deck.New()

	cards := deck.CombineDecks(cardDeck, cardDeck2, cardDeck3)
	cards = deck.Shuffle(cards)

	var player, dealer Hand
	var card deck.Card

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}

	var input string
	reader := bufio.NewReader(os.Stdin)

	for input != "s" {
		fmt.Println("Player's hand:", player)
		fmt.Println("Dealer's hand", dealer.DealerString())
		fmt.Print("What do you want to do? (s)tand or (h)old: ")
		input, _ = reader.ReadString('\n')
		input = input[0:len(input)-1]
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	if dealer.Score() < 16 || (dealer.Score() == 17 && dealer.getLesserScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	playerScore, dealerScore := player.Score(), dealer.Score()
	fmt.Println("Game ending. Final hands:")
	fmt.Println("Player's hand:", player)
	fmt.Println("Dealer's hand:", dealer)
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


}
