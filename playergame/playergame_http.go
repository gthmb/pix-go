package playergame

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gthmb/pix-go/util"
)

// ListPostBody struct
type ListPostBody struct {
	PlayerID string
	GameID   string
}

// DrawPostBody struct
type DrawPostBody struct {
	NumberOfCards int
}

// HandleList handles the http request/response for the PlayerGame list endpoint
func HandleList(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"GET", "POST"}, w, r); !ok {
		return
	}

	switch r.Method {
	case "GET":
		util.WriteJSONResponse(w, PlayerGameMap.ToSlice())
	case "POST":
		var postData ListPostBody
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &postData)
		postData, ok := interface{}(postData).(ListPostBody)
		if !ok {
			util.WriteErrorResponse(w, http.StatusBadRequest, "Could not parse request body")
		}
		playerGame, error := CreateAndWrite(postData.PlayerID, postData.GameID)
		if error != nil {
			util.WriteErrorResponse(w, http.StatusInternalServerError, error.Error())
		}
		util.WriteJSONResponse(w, playerGame)
	}
}

// using a named return, so fancy
func getPlayerGameDetail(w http.ResponseWriter, r *http.Request) (found *PlayerGame, ok bool) {
	_, id, _ := util.GetRouteParams(r.URL.Path)
	found, ok = PlayerGameMap[id]
	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "PlayerGame not found")
	}
	return
}

// HandleDetail handles the http request/response for the PlayerGame detail endpoint
func HandleDetail(w http.ResponseWriter, r *http.Request) {
	if playerGame, ok := getPlayerGameDetail(w, r); ok {
		util.WriteJSONResponse(w, playerGame)
	}
}

// HandleDraw handles the http request/response for a request to draw cards from the Game deck into a players hand
func HandleDraw(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"POST"}, w, r); !ok {
		return
	}

	playerGame, ok := getPlayerGameDetail(w, r)

	if !ok {
		util.WriteErrorResponse(w, http.StatusBadRequest, "cannot find PlayerGame")
		return
	}

	var postData DrawPostBody
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &postData)
	postData, ok = interface{}(postData).(DrawPostBody)

	if !ok {
		util.WriteErrorResponse(w, http.StatusBadRequest, "unexpected POST body")
		return
	}

	playerGame, cards, err := playerGame.DrawCards(postData.NumberOfCards)

	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("Drew %v from the deck", cards)

	util.WriteJSONResponse(w, playerGame)
}
