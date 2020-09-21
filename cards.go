package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card has an annoying comment
type Card struct {
	Suit int
	Value int
}

// Deck has an annoying comment
type Deck struct {
	Cards []Card
	Index int
}

var valueLabels = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var suitLabels = [4]string{"♠", "♥", "♦", "♠"}

func main() {
	var deck *Deck = createDeck()

	shuffle(deck);

	for i := 0; i < len(deck.Cards); i++ {
		card := dealOneCard(deck)
		fmt.Printf("Dealing: %s\n", getCardLabel(card))
	}
}

func createDeck() *Deck {
	deck := Deck{Cards: make([]Card, 0), Index: 0}
	for i := 0; i < len(suitLabels); i++ {
		for j := 0; j < len(valueLabels); j++ {
			deck.Cards = append(deck.Cards, Card{i, j})
		}
	}
	return &deck
}

func dealOneCard(deck *Deck) (Card) {
	card := deck.Cards[deck.Index];
	deck.Index++
	return card
}

func shuffle(deck *Deck) *Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return deck
}

func getCardLabel(card Card) string {
	return fmt.Sprintf("%s of %s", valueLabels[card.Value], suitLabels[card.Suit])
}
