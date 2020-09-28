package deck

import (
	"net/http"

	"github.com/gthmb/pix-go/util"
)

// HandleDetail handles the http request/response for the game detail endpoint
func HandleDetail(w http.ResponseWriter, r *http.Request) {
	if ok := util.ValidateRequestMethod([]string{"GET"}, w, r); !ok {
		return
	}

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
