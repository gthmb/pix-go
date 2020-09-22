package cards

import (
	"fmt"
	"math/rand"
	"time"
)

// Card struct
type Card struct {
	Suit int
	Value int
}

// Deck struct
type Deck struct {
	Cards []Card
	Index int
}

// ValueLabels is an array of value labels
var ValueLabels = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

// SuitLabels is an array of suit labels
var SuitLabels = [4]string{"♠", "♥", "♦", "♠"}

// CreateDeck creates a deck of cards
func CreateDeck() *Deck {
	deck := Deck{Cards: make([]Card, 0), Index: 0}
	for i := 0; i < len(SuitLabels); i++ {
		for j := 0; j < len(ValueLabels); j++ {
			deck.Cards = append(deck.Cards, Card{i, j})
		}
	}
	return &deck
}

// DealCards deals cards from a deck
func (deck *Deck) DealCards(num int) ([]Card, error) {
	newIndex := deck.Index + num

	if l := len(deck.Cards); newIndex >= l {
		return nil, fmt.Errorf("cannot deal %d cards since the Index is %d and there are only %d in the deck", num, deck.Index, l)
	}

	cards := deck.Cards[deck.Index:newIndex]
	deck.Index = newIndex;

	return cards, nil
}

// Shuffle shuffles the cards in a deck
func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

// Describe prints a string representation of a Card
func (card Card) Describe() string {
	return fmt.Sprintf("%s of %s", ValueLabels[card.Value], SuitLabels[card.Suit])
}
