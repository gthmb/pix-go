package card


import (
	"fmt"
)

// Card struct
type Card struct {
	Suit int
	Value int
}

// ValueLabels is an array of value labels
var ValueLabels = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

// SuitLabels is an array of suit labels
var SuitLabels = [4]string{"♠", "♥", "♦", "♠"}

// Describe prints a string representation of a Card
func (card Card) Describe() string {
	return fmt.Sprintf("%s of %s", ValueLabels[card.Value], SuitLabels[card.Suit])
}
