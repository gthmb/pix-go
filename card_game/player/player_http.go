package player

import (
	"fmt"
	"net/http"

	"github.com/gthmb/pix-go/util"
)

// Rendered adds Links
type Rendered struct {
	Player
	Links map[string]string
}

func (player Player) render() (rendered Rendered) {
	rendered = Rendered{
		Player: player,
		Links:  make(map[string]string),
	}

	rendered.Links["self"] = fmt.Sprintf("%s/player/%s", util.Host, player.ID)

	return
}

func (m Map) render() (rendered []Rendered) {
	rendered = make([]Rendered, 0)

	for _, val := range m {
		rendered = append(rendered, val.render())
	}

	return
}

// HandleList handles the http request/response for the player list endpoint
func HandleList(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"GET", "POST"}, w, r); !ok {
		return
	}

	switch r.Method {
	case "GET":
		util.WriteJSONResponse(w, PlayerMap.render())
	case "POST":
		newPlayer, _ := CreateAndWrite()
		util.WriteJSONResponse(w, newPlayer.render())
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
