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

// Host is the hostname:port value concatented from environment variables
var Host string

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

// ValidateRequestMethod returns true if request method is in the allowed methods, otherwise write an http error and returns false
func ValidateRequestMethod(methods []string, w http.ResponseWriter, r *http.Request) (ok bool) {
	for _, method := range methods {
		if r.Method == method {
			return true
		}
	}
	WriteErrorResponse(w, http.StatusMethodNotAllowed, fmt.Sprintf("%s not allowed", r.Method))
	return false
}

// probably bad practive to init env variables outside of the main/main
func init() {
	os.Setenv("SERVER_PROTOCOL", "http")
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("SERVER_HOSTNAME", "localhost")
	Host = fmt.Sprintf("%s://%s:%s", os.Getenv("SERVER_PROTOCOL"), os.Getenv("SERVER_HOSTNAME"), os.Getenv("SERVER_PORT"))
}
