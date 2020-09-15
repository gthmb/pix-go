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
	deck := shuffle(createDeck())
	for i := 0; i < len(deck.Cards); i++ {
		card, newDeck := dealOneCard(deck)
		deck = newDeck
		fmt.Printf("Dealing: %s\n", getCardLabel(card))
	}
}

func createDeck() Deck {
	deck := Deck{Cards: make([]Card, 0), Index: 0}
	for i := 0; i < len(suitLabels); i++ {
		for j := 0; j < len(valueLabels); j++ {
			deck.Cards = append(deck.Cards, Card{i, j})
		}
	}
	return deck
}

func dealOneCard(deck Deck) (Card, Deck) {
	return deck.Cards[deck.Index], Deck{Cards: deck.Cards, Index: deck.Index + 1}
}

func shuffle(deck Deck) Deck {
	rand.Seed(time.Now().UnixNano())
	shuffled := Deck{Cards: make([]Card, 0), Index: 0}
	var count = len(deck.Cards)
	var indexes = rand.Perm(count)
	for i := 0; i < count; i++ {
		shuffled.Cards = append(shuffled.Cards, deck.Cards[indexes[i]])
	}
	return shuffled
}

func getCardLabel(card Card) string {
	return fmt.Sprintf("%s of %s", valueLabels[card.Value], suitLabels[card.Suit])
}
