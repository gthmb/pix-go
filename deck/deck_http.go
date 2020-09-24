package deck

import (
	"net/http"

	"github.com/gthmb/pix-go/util"
)

// HandleDeckDetail handles the http request/response for the game detail endpoint
func HandleDeckDetail(w http.ResponseWriter, r *http.Request) {
	_, id, action := util.GetRouteParams(r.URL.Path)

	deckVal, ok := DeckMap[id]

	if !ok {
		util.WriteErrorResponse(w, http.StatusNotFound, "Deck not found")
		return
	}

	if action != "" {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Bad Request")
	}

	util.WriteJSONResponse(w, deckVal)
}
