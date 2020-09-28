package deck

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gthmb/pix-go/card"
)

func verifyDeck(t *testing.T, deck *Deck) {
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

	suitLength := len(card.SuitLabels)
	valueLength := len(card.ValueLabels)

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
	deck := CreateDeck()
	verifyDeck(t, deck)
}

func TestShuffleDeck(t *testing.T) {
	deck := CreateDeck()
	shuffled := CreateDeck()
	shuffled.Shuffle()

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

func TestDrawCards(t *testing.T) {
	var deck *Deck = CreateDeck()
	var deckValue = deck

	deckIndex := 0

	// deal groups of cards, from 1 to 10
	for i := 0; i < 10; i++ {
		cards, error := deckValue.DrawCards(i)
		expectedIndex := deckIndex + i
		expectedCards := deck.Cards[deckIndex:expectedIndex]

		if error != nil {
			t.Errorf("error dealing cards: %v", error)
		}

		if l := len(cards); l != i {
			t.Errorf("Should have dealt %d cards, but only got %d", i, l)
		}

		if deckValue.Index != expectedIndex {
			t.Errorf("After the deal the index should be %d, not %d", deckIndex, deckValue.Index)
		}

		deckIndex = expectedIndex

		for i = 0; i < len(cards); i++ {
			if c, e := cards[i], expectedCards[i]; c != e {
				t.Errorf("Got %v, but expected %v", c, e)
			}
		}
	}
}

func TestDrawTooManyCards(t *testing.T) {
	var deck = CreateDeck()

	cards, error := deck.DrawCards(1000)

	if cards != nil || error == nil {
		t.Error("Should have errored on the deal, deck does not contain 1000 cards")
	}
}

func TestDrawCardsConcurrency(t *testing.T) {
	var wg sync.WaitGroup

	deck := CreateDeck()
	hand := make([]card.Card, 0)
	numDeals, numCards := 3, 3
	hasDupes := false

	for i := 0; i < numDeals; i++ {
		wg.Add(1)
		go func(indx int) {
			cards, _ := deck.DrawCards(numCards)

			// if any card dealt back is already in the hand, set hasDupes to true
			for _, el := range cards {
				for _, la := range hand {
					if el.Suit == la.Suit && el.Value == la.Value {
						hasDupes = true
					}
				}
			}

			hand = append(hand, cards...)
			wg.Done()
		}(i)
	}
	wg.Wait()

	l, expected := len(hand), numDeals*numCards
	if l != numDeals*numCards {
		t.Errorf("Hand should have %d cards, but got %d", expected, l)
	}

	if hasDupes {
		t.Errorf("Hand has duplicate cards: %v", hand)
	}

	fmt.Printf("Dealt %d cards\n", l)
	fmt.Printf("hand looks like: %v cards\n", hand)
}

// work around handling flags in tests. @see https://github.com/golang/go/issues/31859#issuecomment-489889428
var _ = func() bool {
	testing.Init()
	return true
}()
