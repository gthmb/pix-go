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

	for i := 0; i < len(deck.Cards) + 1; i++ {
		cards, error := dealCards(deck, 1)
		if error == nil {
			for _, i := range cards {
				fmt.Printf("Dealing: %s\n", getCardLabel(i))
			}
		}
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

func dealCards(deck *Deck, num int) ([]Card, error) {
	newIndex := deck.Index + num

	if l := len(deck.Cards); newIndex >= l {
		return nil, fmt.Errorf("cannot deal %d cards since the Index is %d and there are only %d in the deck", num, deck.Index, l)
	}

	cards := deck.Cards[deck.Index:newIndex]
	deck.Index = newIndex;

	return cards, nil
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
