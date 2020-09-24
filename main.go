package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gthmb/pix-go/deck"
	"github.com/gthmb/pix-go/game"
	"github.com/gthmb/pix-go/player"
	"github.com/gthmb/pix-go/util"
)

var rGameStart = regexp.MustCompile(`/games/\d+/start$`)
var rGameJoin = regexp.MustCompile(`/games/\d+/join$`)
var rGameDetail = regexp.MustCompile(`/games/\d+$`)
var rGameList = regexp.MustCompile(`/games$`)
var rDeckDetail = regexp.MustCompile(`/decks/\d+$`)
var rPlayerList = regexp.MustCompile(`/players$`)

func handleRoute(w http.ResponseWriter, r *http.Request) {
	switch {
	case rGameList.MatchString(r.URL.Path):
		game.HandleGameList(w, r)
	case rGameDetail.MatchString(r.URL.Path):
		game.HandleGameDetail(w, r)
	case rGameJoin.MatchString(r.URL.Path):
		game.HandleGameJoin(w, r)
	case rGameStart.MatchString(r.URL.Path):
		game.HandleGameStart(w, r)
	case rDeckDetail.MatchString(r.URL.Path):
		deck.HandleDeckDetail(w, r)
	case rPlayerList.MatchString(r.URL.Path):
		player.HandlePlayerList(w, r)
	default:
		util.WriteErrorResponse(w, http.StatusNotFound, "Noooooope!")
	}
}

func main() {
	http.HandleFunc("/", handleRoute)
	fmt.Println("Server is running on port :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
