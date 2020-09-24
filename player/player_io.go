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

// WriteAll blah
func WriteAll() error {
	return util.WriteJSONFile(filePath, PlayerMap)
}

// Put blah
func Put(id string, player Player) error {
	PlayerMap[id] = player
	return WriteAll()
}
