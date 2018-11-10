package api1

import (
	"net/http"

	"template.github.com/server/api1/model"

	"github.com/gorilla/mux"
	"template.github.com/server/web"
)

const (
	API_URL_SUFFIX = "/api/v1"
)

//Routes all routes of systems
type Routes struct {
	Root    *mux.Router // ''
	ApiRoot *mux.Router // 'api/v4'

	Candidates *mux.Router // 'api/v4/candidates'
	Candidate  *mux.Router // 'api/v4/candidates/{user_id:[A-Za-z0-9]+}'
}

type API struct {
	BaseRoutes *Routes
}

func Init(root *mux.Router) *API {
	api := &API{
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = root
	api.BaseRoutes.ApiRoot = root.PathPrefix(API_URL_SUFFIX).Subrouter()
	api.BaseRoutes.Candidates = api.BaseRoutes.ApiRoot.PathPrefix("/candidates").Subrouter()
	api.BaseRoutes.Candidate = api.BaseRoutes.ApiRoot.PathPrefix("/candidates/{candidateId:[A-Za-z0-9]+}").Subrouter()
	api.InitCandidate()

	root.Handle("/api/v4/{anything:.*}", http.HandlerFunc(api.Handle404))
	return api
}

//APIHandler handler wrapper
func (api *API) APIHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &web.Handler{
		HandleFunc: h,
	}
}

//Handle404 Handle404
func (api *API) Handle404(w http.ResponseWriter, r *http.Request) {
	web.Handle404(w, r)
}

//HandleAPIError Handle API Error
func HandleAPIError(e *model.APIError, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(model.APIErrorToJSON(e)))
}

var ReturnStatusOK = web.ReturnStatusOK

type Context = web.Context
