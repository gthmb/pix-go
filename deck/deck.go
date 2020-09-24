package deck

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gthmb/pix-go/card"
)

// Deck struct
type Deck struct {
	ID    string
	Cards []card.Card
	Index int
}

// Map struct
type Map map[string]Deck

// DeckMap is a map of all the Decks
var DeckMap Map = make(map[string]Deck)

// CurrentDeckID is a sill ID index trakcer
var CurrentDeckID int = 0

// CreateDeck creates a deck of cards
func CreateDeck() Deck {

	deck := Deck{
		ID:    fmt.Sprint(len(DeckMap) + 1),
		Cards: make([]card.Card, 0),
		Index: 0,
	}

	for i := 0; i < len(card.SuitLabels); i++ {
		for j := 0; j < len(card.ValueLabels); j++ {
			deck.Cards = append(deck.Cards, card.Card{Suit: i, Value: j})
		}
	}
	return deck
}

// DealCards deals cards from a deck
func (deck *Deck) DealCards(num int) ([]card.Card, error) {
	newIndex := deck.Index + num

	if l := len(deck.Cards); newIndex >= l {
		return nil, fmt.Errorf("cannot deal %d cards since the Index is %d and there are only %d in the deck", num, deck.Index, l)
	}

	cards := deck.Cards[deck.Index:newIndex]
	deck.Index = newIndex

	return cards, nil
}

// Shuffle shuffles the cards in a deck
func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

func init() {
	DeckMap, _ = FetchAll()
}
