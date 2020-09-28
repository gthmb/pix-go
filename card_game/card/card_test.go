package card

import (
	"fmt"
	"testing"
)

func TestDescribeCard(t *testing.T) {
	for i := 0; i < 52; i++ {
		card := Card{Suit: i % 4, Value: i % 13}
		cardLabel, expectedLabel := card.Describe(), fmt.Sprintf("%s of %s", ValueLabels[card.Value], SuitLabels[card.Suit])
		if cardLabel != expectedLabel {
			t.Errorf("Expected card label to be %s, but got %s", cardLabel, expectedLabel)
		}
	}
}
