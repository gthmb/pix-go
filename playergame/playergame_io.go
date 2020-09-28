package playergame

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gthmb/pix-go/util"
)

const filePath string = "./data/playergames.json"

// FetchAll fetchs GamePlayer data from the filesystem
func FetchAll() (Map, error) {
	var dest = make(map[string]*PlayerGame)
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

// CreateAndWrite creates a playergame and saves it to disk
func CreateAndWrite(playerID, gameID string) (newPlayerGame *PlayerGame, err error) {
	newPlayerGame, err = CreatePlayerGame(playerID, gameID)
	Put(newPlayerGame)
	return
}

// WriteAll blah
func WriteAll() error {
	return util.WriteJSONFile(filePath, PlayerGameMap)
}

// Put blah
func Put(el *PlayerGame) error {
	PlayerGameMap[el.ID] = el
	return WriteAll()
}
