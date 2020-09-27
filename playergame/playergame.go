package playergame

import (
	"errors"
	"fmt"

	"github.com/gthmb/pix-go/card"
	"github.com/gthmb/pix-go/deck"
	"github.com/gthmb/pix-go/game"
	"github.com/gthmb/pix-go/player"
)

// PlayerGame struct
type PlayerGame struct {
	ID       string
	GameID   string
	PlayerID string
	Hand     []card.Card
}

// Map struct
type Map map[string]PlayerGame

// PlayerGameMap is a map of all the GamePlayer instances
var PlayerGameMap Map = make(map[string]PlayerGame)

// CreatePlayerGame makes a game with the supplied players
func CreatePlayerGame(playerID, gameID string) (PlayerGame, error) {

	if _, ok := player.PlayerMap[playerID]; !ok {
		return PlayerGame{}, errors.New("Cannot find Player")
	}

	foundGame, ok := game.GameMap[gameID]

	if !ok {
		return PlayerGame{}, errors.New("Cannot find Game")
	}

	for _, pid := range foundGame.PlayerGameIDs {
		if pid == playerID {
			return PlayerGame{}, fmt.Errorf("Player %s is already in game %s", playerID, gameID)
		}
	}

	playerGame := PlayerGame{
		ID:       fmt.Sprint(len(PlayerGameMap) + 1),
		GameID:   gameID,
		PlayerID: playerID,
		Hand:     make([]card.Card, 0),
	}

	foundGame.PlayerGameIDs = append(foundGame.PlayerGameIDs, playerGame.ID)
	game.Put(foundGame)

	return playerGame, nil
}

// DrawCards attempts to draw cards from the game deck into the players hand
func (pg PlayerGame) DrawCards(num int) (PlayerGame, []card.Card, error) {
	var cards = make([]card.Card, 0)

	if num < 1 {
		return pg, cards, errors.New("invalid NumberOfCards")
	}

	foundGame, ok := game.GameMap[pg.GameID]

	if !ok {
		return pg, cards, errors.New("could not find Game")
	}

	if !foundGame.Started {
		return pg, cards, fmt.Errorf("game %s not started", foundGame.ID)
	}

	foundDeck, ok := deck.DeckMap[foundGame.DeckID]

	if !ok {
		return pg, cards, errors.New("could not find Deck")
	}

	cards, error := foundDeck.DrawCards(num)

	if error != nil {
		return pg, cards, error
	}

	pg.Hand = append(pg.Hand, cards...)

	Put(pg)
	deck.Put(foundDeck)

	return pg, cards, nil
}

// ToSlice converts Map.Games into a slice of Games
func (m Map) ToSlice() []PlayerGame {
	slice := make([]PlayerGame, len(m))
	index := 0
	for _, el := range m {
		slice[index] = el
		index++
	}
	return slice
}

func init() {
	PlayerGameMap, _ = FetchAll()
}
