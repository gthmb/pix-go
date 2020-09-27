package player

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gthmb/pix-go/util"
)

const filePath string = "./data/players.json"

// FetchAll fetchs Deck data from the filesystem
func FetchAll() (Map, error) {
	var dest = make(map[string]Player)
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

// CreateAndWrite creates a player and saves it to disk
func CreateAndWrite() (newPlayer Player, err error) {
	newPlayer, err = CreatePlayer()
	Put(newPlayer)
	return
}

// WriteAll blah
func WriteAll() error {
	return util.WriteJSONFile(filePath, PlayerMap)
}

// Put blah
func Put(player Player) error {
	PlayerMap[player.ID] = player
	return WriteAll()
}
