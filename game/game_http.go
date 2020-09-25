package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gthmb/pix-go/deck"
	"github.com/gthmb/pix-go/player"
	"github.com/gthmb/pix-go/util"
)

// JoinPostBody struct
type JoinPostBody struct {
	PlayerID string
}

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

// using a named return, so fancy
func getGameDetail(w http.ResponseWriter, r *http.Request) (foundGame Game, ok bool) {
	_, id, _ := util.GetRouteParams(r.URL.Path)
	foundGame, ok = GameMap[id]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Game not found")
	}

	return
}

// HandleGameDetail handles the http request/response for the game detail endpoint
func HandleGameDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
		return
	}

	foundGame, ok := getGameDetail(w, r)

	if !ok {
		return
	}

	util.WriteJSONResponse(w, foundGame)
}

// HandleGameJoin handles the http request/response for the game join endpoint
func HandleGameJoin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
		return
	}

	foundGame, ok := getGameDetail(w, r)

	if !ok {
		return
	}

	var postData JoinPostBody
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &postData)

	postData, ok = interface{}(postData).(JoinPostBody)

	if ok {
		if foundPlayer, pok := player.PlayerMap[postData.PlayerID]; pok {
			foundPlayer.Games = append(foundPlayer.Games, foundGame.ID)
			for _, pid := range foundGame.PlayerIDs {
				if pid == postData.PlayerID {
					util.WriteErrorResponse(w, http.StatusConflict, "Player is already in Game")
					return
				}
			}
			foundGame.PlayerIDs = append(foundGame.PlayerIDs, postData.PlayerID)
			Put(foundGame.ID, foundGame)
			player.Put(foundPlayer.ID, foundPlayer)
			util.WriteJSONResponse(w, foundGame)
		} else {
			util.WriteErrorResponse(w, http.StatusNotFound, "Player not found")
		}

		return
	}

	util.WriteErrorResponse(w, http.StatusInternalServerError, "Borked")
}

// HandleGameStart handles the http request/response for the game start endpoint
func HandleGameStart(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
		return
	}

	foundGame, ok := getGameDetail(w, r)

	if !ok {
		return
	}

	foundDeck, ok := deck.DeckMap[foundGame.ID]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Cannot find Deck for Game")
		return
	}

	if len(foundGame.PlayerIDs) < 2 {
		util.WriteErrorResponse(w, http.StatusNotFound, "Game needs at least 2 players")
		return
	}

	if foundGame.Start() {
		foundDeck.Shuffle()
		Put(foundGame.ID, foundGame)
		deck.Put(foundDeck.ID, foundDeck)
	} else {
		util.WriteErrorResponse(w, http.StatusConflict, "Game already started")
		return
	}

	util.WriteJSONResponse(w, foundGame)
}
