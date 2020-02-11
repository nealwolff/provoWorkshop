package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nealwolff/provoWorkshop/crud"
	"github.com/nealwolff/provoWorkshop/types"
)

//UserHandler handles user functionality
func UserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var user types.User

		json.NewDecoder(r.Body).Decode(&user)

		if user.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("A user must have a name"))
			return
		}

		ret, err := crud.Insert("users", user, w)
		if err != nil {
			return
		}

		json.NewEncoder(w).Encode(ret)
	} else if r.Method == http.MethodGet {
		params := mux.Vars(r)

		ID := params["id"]

		if ID == "" {
			//TODO get all
			return
		}

		data, err := crud.GetOne("users", ID, w)

		if err != nil {
			return
		}

		w.Write(data)
	}

}
