package deck

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gthmb/pix-go/util"
)

const filePath string = "./data/decks.json"

// FetchAll fetchs Deck data from the filesystem
func FetchAll() (Map, error) {
	var dest = make(map[string]*Deck)
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return dest, err
	}

	err = json.Unmarshal([]byte(data), &dest)

	if err != nil {
		return dest, err
	}

	return dest, nil
}

// WriteAll blah
func WriteAll() error {
	return util.WriteJSONFile(filePath, DeckMap)
}

// Put blah
func Put(deck *Deck) error {
	DeckMap[deck.ID] = deck
	return WriteAll()
}
