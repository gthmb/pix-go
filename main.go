package main

import (
	"fmt"
	"github.com/gthmb/pix-go/cards"
)

func main() {
	var deck *cards.Deck = cards.CreateDeck()

	deck.Shuffle()

	for i := 0; i < len(deck.Cards) + 1; i++ {
		dealtCards, error := deck.DealCards(1)
		if error == nil {
			for _, i := range dealtCards {
				fmt.Printf("Dealing: %s\n", i.Describe())
			}
		}
	}
}
