package game

import (
	"fmt"

	"github.com/gthmb/pix-go/deck"
)

// Game struct
type Game struct {
	ID        string
	DeckID    string
	PlayerIDs []string
	Started   bool
}

// Map struct
type Map map[string]Game

// GameMap is a map of all the Games
var GameMap Map = make(map[string]Game)

// CreateGame makes a game with the supplied players
func CreateGame(playerIDs []string) (Game, deck.Deck, error) {
	if playerIDs == nil {
		playerIDs = make([]string, 0)
	}

	deck := deck.CreateDeck()

	return Game{
		ID:        fmt.Sprint(len(GameMap) + 1),
		PlayerIDs: playerIDs,
		DeckID:    deck.ID,
		Started:   false,
	}, deck, nil
}

// Start starts a game
func (game *Game) Start() bool {
	if game.Started {
		return false
	}
	game.Started = true
	return true
}

// ToSlice converts Map.Games into a slice of Games
func (m Map) ToSlice() []Game {
	slice := make([]Game, len(m))
	index := 0
	for _, el := range m {
		slice[index] = el
		index++
	}
	return slice
}

func init() {
	GameMap, _ = FetchAll()
}
