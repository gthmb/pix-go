package player

import (
	"fmt"
	"net/http"

	"github.com/gthmb/pix-go/util"
)

// HandlePlayerList handles the http request/response for the player list endpoint
func HandlePlayerList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		util.WriteJSONResponse(w, PlayerMap.ToSlice())
	case "POST":
		newPlayer, _ := CreatePlayer()
		Put(newPlayer.ID, newPlayer)
		util.WriteJSONResponse(w, newPlayer)
	default:
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
	}
}

// HandlePlayerDetail handles the http request/response for the player detail endpoint
func HandlePlayerDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
		return
	}

	_, id, _ := util.GetRouteParams(r.URL.Path)
	foundGame, ok := PlayerMap[id]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Player not found")
		return
	}

	util.WriteJSONResponse(w, foundGame)
}
