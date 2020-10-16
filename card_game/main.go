package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gthmb/pix-go/deck"
	"github.com/gthmb/pix-go/game"
	"github.com/gthmb/pix-go/player"
	"github.com/gthmb/pix-go/playergame"
	"github.com/gthmb/pix-go/util"
)

var rGameDetail = regexp.MustCompile(`/games/\d+$`)
var rGameList = regexp.MustCompile(`/games$`)
var rGameStart = regexp.MustCompile(`/games/\d+/start$`)
var rDeckDetail = regexp.MustCompile(`/decks/\d+$`)
var rPlayerList = regexp.MustCompile(`/players$`)
var rPlayerGameDetail = regexp.MustCompile(`/playergames/\d+$`)
var rPlayerGameList = regexp.MustCompile(`/playergames$`)
var rPlaygerGameDraw = regexp.MustCompile(`/playergames/\d+/draw$`)

func handleRoute(w http.ResponseWriter, r *http.Request) {
	switch {
	case rGameList.MatchString(r.URL.Path):
		game.HandleList(w, r)
	case rGameDetail.MatchString(r.URL.Path):
		game.HandleDetail(w, r)
	case rGameStart.MatchString(r.URL.Path):
		game.HandleGameStart(w, r)
	case rDeckDetail.MatchString(r.URL.Path):
		deck.HandleDetail(w, r)
	case rPlayerList.MatchString(r.URL.Path):
		player.HandleList(w, r)
	case rPlayerGameList.MatchString(r.URL.Path):
		playergame.HandleList(w, r)
	case rPlayerGameDetail.MatchString(r.URL.Path):
		playergame.HandleDetail(w, r)
	case rPlaygerGameDraw.MatchString(r.URL.Path):
		playergame.HandleDraw(w, r)
	default:
		util.WriteErrorResponse(w, http.StatusNotFound, "Noooooope!")
	}
}

func main() {
	http.HandleFunc("/", handleRoute)
	fmt.Printf("Server is running: %s\n", util.Host)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOSTNAME"), os.Getenv("SERVER_PORT")), nil))
}
