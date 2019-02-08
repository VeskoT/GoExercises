//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
	)

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) getSuit() Suit {
	return c.Suit
}

func (c Card) getRank() Rank {
	return c.Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("{%s of %ss}", c.Rank.String(), c.Suit.String())
}

func newCard(suit Suit, rank Rank) *Card {
	c := new(Card)
	c.Suit = suit
	c.Rank = rank
	return c
}

var suits = [...]Suit{Spade, Diamond, Club, Heart}

func New() []Card{
	var deck []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			card := newCard(suit, rank)
			deck = append(deck, *card)
		}
	}
	return deck
}

func DefaultSort(cards []Card) []Card{
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool{
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)* int(c.Rank)
}

func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffledDeck := make([]Card, len(cards))
	permutations := r.Perm(len(cards))
	for i, randIndex := range permutations {
		shuffledDeck[i] = cards[randIndex]
	}
	return shuffledDeck
}

func AddJokers(cards []Card, numberOfJokers int) []Card{
	for i:=0; i<numberOfJokers; i++ {
		joker := newCard(Joker, Rank(i))
		cards = append(cards, *joker)
	}
	return cards
}

func FilterOut(cards []Card, ranks ...Rank) []Card{
	for _, rank := range ranks {
		for i, card := range cards {
			if card.Rank == rank {
				cards = append(cards[:i], cards[i+1:]...)
			}
		}
	}
	return cards
}

func CombineDecks(cardDecks ...[]Card) []Card {
	var ret []Card
	for _, cardDeck := range cardDecks {
		ret = append(ret, cardDeck...)
	}
	return ret
}
