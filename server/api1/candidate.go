package api1

import (
	"fmt"
	"net/http"

	"github.com/go-http-utils/headers"

	"github.com/gorilla/mux"
	"template.github.com/server/api1/model"
	"template.github.com/server/app"
)

var App = app.App

//InitCandidate initialize handlers for candidates
func (api *API) InitCandidate() {
	api.BaseRoutes.Candidates.Handle("", api.APIHandler(createCandidate)).Methods("POST").HeadersRegexp(headers.ContentType, "application/json")
	api.BaseRoutes.Candidate.Handle("", api.APIHandler(getCandidateByID)).Methods("GET")
}

//createCandidate handler function for create candidate
func createCandidate(ctx *context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling createCandidate")
	candidate := model.CandidateFromJSON(r.Body)
	result, error := App.Candidate.CreateCandidate(ctx, model.CandidateToInternalCandidate(candidate))
	if error != nil {
		handleAPIError(error, w)
	} else {
		candidate.ID = result
		w.Write([]byte(model.CandidateToJSON(candidate)))
	}
}

//getCandidateById handler function to get candidate by Id
func getCandidateByID(ctx *context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	candidateId := vars["candidateId"]
	iCandidate, error := App.Candidate.FindCandidateByID(ctx, candidateId)
	if error != nil {
		handleAPIError(error, w)
		return
	}
	CandidateToJSON := model.CandidateToJSON(model.InternalCandidateToCandidate(iCandidate))
	w.Write([]byte(CandidateToJSON))
}
