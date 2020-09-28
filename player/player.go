package player

import (
	"fmt"
)

// Player struct
type Player struct {
	ID string
}

// Map struct
type Map map[string]*Player

// PlayerMap is a map of all the Players
var PlayerMap Map = make(map[string]*Player)

// CreatePlayer makes a game with the supplied players
func CreatePlayer() (*Player, error) {
	return &Player{
		ID: fmt.Sprint(len(PlayerMap) + 1),
	}, nil
}

// ToSlice converts Map.Games into a slice of Games
func (m Map) ToSlice() []*Player {
	slice := make([]*Player, len(m))
	index := 0
	for _, el := range m {
		slice[index] = el
		index++
	}
	return slice
}

func init() {
	PlayerMap, _ = FetchAll()
}
