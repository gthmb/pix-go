package game

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gthmb/pix-go/util"
)

const filePath string = "./data/games.json"

// FetchAll fetchs Game data from the filesystem
func FetchAll() (Map, error) {
	var dest = make(map[string]Game)
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
	return util.WriteJSONFile(filePath, GameMap)
}

// Put blah
func Put(id string, game Game) error {
	GameMap[id] = game
	return WriteAll()
}
