package game

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gthmb/pix-go/deck"
	"github.com/gthmb/pix-go/util"
)

const filePath string = "./data/games.json"

// FetchAll fetchs Game data from the filesystem
func FetchAll() (Map, error) {
	var dest = make(map[string]*Game)
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

// CreateAndWrite creates a game with a deck and saves the resources to disk
func CreateAndWrite() (newGame *Game, newDeck *deck.Deck, err error) {
	newGame, newDeck, err = CreateGame()
	Put(newGame)
	deck.Put(newDeck)
	return
}

// WriteAll blah
func WriteAll() error {
	return util.WriteJSONFile(filePath, GameMap)
}

// Put blah
func Put(game *Game) error {
	GameMap[game.ID] = game
	return WriteAll()
}
