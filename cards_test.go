package main

import (
	"fmt"
	"testing"
)

func verifyDeck(t *testing.T, deck Deck) {
	length, expectedLength := len(deck.Cards), 52
	index, expectedIndex := deck.Index, 0

	if length != expectedLength {
		t.Errorf("A Deck should have %d cards, but got %d", length, expectedLength)
	}

	if deck.Index != expectedIndex {
		t.Errorf("A Deck should have a starting index of %d cards, but got %d", index, expectedIndex)
	}

	verifySuits := make(map[int]int)
	verifyValues := make(map[int]int)

	for i := 0; i < length; i++ {
		card := deck.Cards[i]
		verifySuits[card.Suit]++
		verifyValues[card.Value]++
	}

	suitLength := len(suitLabels)
	valueLength := len(valueLabels)

	for i := 0; i < suitLength; i++ {
		if verifySuits[i] != valueLength {
			t.Errorf("Suit%d should have %d Cards, but got %d", i, verifySuits[1], valueLength)
		}
	}

	for i := 0; i < valueLength; i++ {
		if verifyValues[i] != suitLength {
			t.Errorf("Value%d should have %d Cards, but got %d", i, verifyValues[1], suitLength)
		}
	}
}

func TestCreateDeck(t *testing.T) {
	deck := createDeck()
	verifyDeck(t, deck)
}

func TestShuffleDeck(t *testing.T) {
	deck := createDeck()
	shuffled := shuffle(deck)

	verifyDeck(t, shuffled)

	deckHash := ""
	shuffleHash := ""

	for i := 0; i < len(deck.Cards); i++ {
		deckHash = fmt.Sprintf("%v%d%d", deckHash, deck.Cards[i].Value, deck.Cards[i].Value)
		shuffleHash = fmt.Sprintf("%v%d%d", shuffleHash, shuffled.Cards[i].Value, shuffled.Cards[i].Value)
	}

	if deckHash == shuffleHash {
		t.Error("The hash of shuffled deck should be different than that of the unshuffled deck")
	}

}

func TestDealOneCard(t *testing.T) {
	deck := createDeck()
	deckIndex := deck.Index
	card := Card{}

	// deal ten and test after each
	for i := 0; i < 10; i++ {
		card, deck = dealOneCard(deck)
		expectedCard := deck.Cards[i]

		if card != expectedCard {
			t.Errorf("The Card should be the same as the deck's card at it's Index. Got %v, expected %v", card, expectedCard)
		}

		index, expectedIndex := deck.Index, deckIndex+i+1
		if index != expectedIndex {
			t.Errorf("Dealing a card should return a deck with an incremented Index. Got %d, expected %d", index, expectedIndex)
		}
	}

}

func TestGetCardLabel(t *testing.T) {
	for i := 0; i < 52; i++ {
		card := Card{Suit: i % 4, Value: i % 13}
		cardLabel, expectedLabel := getCardLabel(card), fmt.Sprintf("%s of %s", valueLabels[card.Value], suitLabels[card.Suit])
		if( cardLabel != expectedLabel){
			t.Errorf("Expected card label to be %s, but got %s", cardLabel, expectedLabel)
		}
	}
}
