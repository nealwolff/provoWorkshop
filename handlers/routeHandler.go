package handlers

import (
	"encoding/json"
	"net/http"
)

//BasicHandler handles the basic route
func BasicHandler(w http.ResponseWriter, r *http.Request) {
	ret := map[string]string{
		"key": "Hello World",
	}

	retRaw, err := json.Marshal(ret)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(retRaw)
}
