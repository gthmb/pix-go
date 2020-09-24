package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// WriteJSONResponse marshals and writes a json response, or writes an error response
func WriteJSONResponse(w http.ResponseWriter, data interface{}) {
	str, err := json.Marshal(data)

	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(str))
}

// WriteJSONFile writes a stringified JSON version of the data to a file with the supplied path
func WriteJSONFile(writepath string, data interface{}) error {
	str, _ := json.Marshal(data)

	os.MkdirAll(filepath.Dir(writepath), os.ModePerm)

	return ioutil.WriteFile(writepath, []byte(str), 0644)
}

// WriteErrorResponse writes an error response
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	fmt.Fprintf(w, "%d Error - %s", code, message)
}

// GetRouteParams splits a route string and returns the resource, id, and action
func GetRouteParams(route string) (resource, id, action string) {
	pathSlice := append(strings.Split(route, "/"), "", "", "", "")
	return pathSlice[1], pathSlice[2], pathSlice[3]
}
