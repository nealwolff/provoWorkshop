package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nealwolff/provoWorkshop/crud"
	"github.com/nealwolff/provoWorkshop/types"
)

//AmoritizationHandler handles logic associated with calculating amoritization schedule
func AmoritizationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID := params["id"]
	loanID := params["loanId"]

	userRaw, err := crud.GetOne("users", ID, w)

	if err != nil {
		return
	}

	user := types.User{}
	json.Unmarshal(userRaw, &user)

	if loan, ok := user.Loans[loanID]; ok {
		aSched := loan.CalculateASchedule()

		ret, _ := json.MarshalIndent(aSched, "", "    ")
		w.Write(ret)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("The loan does not exist"))
}

//LoanHandler adds loans to the user in the database
func LoanHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID := params["id"]
	loanID := params["loanId"]

	userRaw, err := crud.GetOne("users", ID, w)

	if err != nil {
		return
	}

	user := types.User{}
	json.Unmarshal(userRaw, &user)

	if user.Loans == nil {
		user.Loans = make(map[string]types.Loan)
	}

	var loan types.Loan
	json.NewDecoder(r.Body).Decode(&loan)

	user.Loans[loanID] = loan

	resp, err := crud.Update("users", ID, user, w)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(resp)

}
