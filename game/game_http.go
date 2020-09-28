package game

import (
	"fmt"
	"net/http"

	"github.com/gthmb/pix-go/util"
)

// HandleList handles the http request/response for the game list endpoint
func HandleList(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"GET", "POST"}, w, r); !ok {
		return
	}
	switch r.Method {
	case "GET":
		util.WriteJSONResponse(w, GameMap.render())
	case "POST":
		newGame, _, _ := CreateAndWrite()
		util.WriteJSONResponse(w, newGame.render())
	}
}

// Rendered adds Links
type Rendered struct {
	Game
	Links map[string]string
}

func (game Game) render() (rendered Rendered) {
	rendered = Rendered{
		Game:  game,
		Links: make(map[string]string),
	}

	rendered.Links["self"] = fmt.Sprintf("%s/games/%s", util.Host, game.ID)
	rendered.Links["start"] = fmt.Sprintf("%s/games/%s/start", util.Host, game.ID)

	return
}

func (m Map) render() (rendered []Rendered) {
	rendered = make([]Rendered, 0)

	for _, val := range m {
		rendered = append(rendered, val.render())
	}

	return
}

// using named return values, so fancy
func getGameDetail(w http.ResponseWriter, r *http.Request) (foundGame *Game, ok bool) {
	_, id, _ := util.GetRouteParams(r.URL.Path)

	foundGame, ok = GameMap[id]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Game not found")
	}

	return
}

// HandleDetail handles the http request/response for the game detail endpoint
func HandleDetail(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"GET"}, w, r); !ok {
		return
	}

	foundGame, ok := getGameDetail(w, r)

	if !ok {
		return
	}

	util.WriteJSONResponse(w, foundGame.render())
}

// HandleGameStart handles the http request/response for the game start endpoint
func HandleGameStart(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"POST"}, w, r); !ok {
		return
	}

	foundGame, ok := getGameDetail(w, r)

	if !ok {
		util.WriteErrorResponse(w, http.StatusBadRequest, "cannot find Game")
		return
	}

	foundGame, err := foundGame.Start()

	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	util.WriteJSONResponse(w, foundGame)
}
