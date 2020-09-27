package game

import (
	"fmt"

	"github.com/gthmb/pix-go/deck"
)

// Game struct
type Game struct {
	ID            string
	DeckID        string
	PlayerGameIDs []string
	Started       bool
}

// Map struct
type Map map[string]Game

// GameMap is a map of all the Games
var GameMap Map = make(map[string]Game)

// CreateGame makes a game with the supplied players
func CreateGame() (Game, deck.Deck, error) {
	deck := deck.CreateDeck()

	return Game{
		ID:            fmt.Sprint(len(GameMap) + 1),
		PlayerGameIDs: make([]string, 0),
		DeckID:        deck.ID,
		Started:       false,
	}, deck, nil
}

// Start starts a game
func (game Game) Start() (Game, error) {
	if game.Started {
		return game, fmt.Errorf("game %s is already started", game.ID)
	}

	foundDeck, ok := deck.DeckMap[game.ID]

	if !ok {
		return game, fmt.Errorf("cannot find deck for game %s", game.ID)
	}

	if len(game.PlayerGameIDs) < 2 {
		return game, fmt.Errorf("game %s needs at least 2 players", game.ID)
	}

	game.Started = true
	foundDeck.Shuffle()
	Put(game)
	deck.Put(foundDeck)

	return game, nil
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
