package game

import (
	"fmt"
	"net/http"

	"github.com/gthmb/pix-go/deck"
	"github.com/gthmb/pix-go/util"
)

// HandleGameList handles the http request/response for the game list endpoint
func HandleGameList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		util.WriteJSONResponse(w, GameMap.ToSlice())
	case "POST":
		newGame, newDeck, _ := CreateGame(nil)
		Put(newGame.ID, newGame)
		deck.Put(newDeck.ID, newDeck)
		util.WriteJSONResponse(w, newGame)
	default:
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
	}
}

// HandleGameDetail handles the http request/response for the game detail endpoint
func HandleGameDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
		return
	}

	_, id, _ := util.GetRouteParams(r.URL.Path)
	foundGame, ok := GameMap[id]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Game not found")
		return
	}

	util.WriteJSONResponse(w, foundGame)
}

// HandleGameJoin handles the http request/response for the game join endpoint
func HandleGameJoin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
		return
	}
}

// HandleGameStart handles the http request/response for the game start endpoint
func HandleGameStart(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
		return
	}

	_, id, _ := util.GetRouteParams(r.URL.Path)
	foundGame, ok := GameMap[id]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Game not found")
		return
	}

	foundDeck, ok := deck.DeckMap[foundGame.ID]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Cannot find Deck for Game")
		return
	}

	if foundGame.Start() {
		foundDeck.Shuffle()
		Put(id, foundGame)
		deck.Put(id, foundDeck)
	} else {
		util.WriteErrorResponse(w, http.StatusConflict, "Game already started")
		return
	}

	util.WriteJSONResponse(w, foundGame)
}
