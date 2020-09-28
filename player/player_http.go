package player

import (
	"net/http"

	"github.com/gthmb/pix-go/util"
)

// HandleList handles the http request/response for the player list endpoint
func HandleList(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"GET", "POST"}, w, r); !ok {
		return
	}

	switch r.Method {
	case "GET":
		util.WriteJSONResponse(w, PlayerMap.ToSlice())
	case "POST":
		newPlayer, _ := CreateAndWrite()
		util.WriteJSONResponse(w, newPlayer)
	}
}

// HandleDetail handles the http request/response for the player detail endpoint
func HandleDetail(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"GET"}, w, r); !ok {
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
