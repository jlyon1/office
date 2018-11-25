package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/jlyon1/office/search"
)

// API is a struct that represents what is needed to serve the search api
type API struct {
	handler  http.Handler
	episodes search.Episodes
}

// New creates a new api struct
func New(episodes search.Episodes) *API {
	api := &API{
		episodes: episodes,
	}
	r := chi.NewRouter()
	r.Route("/office", func(r chi.Router) {
		r.Get("/search", api.Search)
	})
	api.handler = r
	return api
}

// Search handles get requests for episodes, searches using the query at url parameter search
func (a *API) Search(w http.ResponseWriter, r *http.Request) {

	keys, _ := r.URL.Query()["query"]
	if len(keys) == 0 {
		w.WriteHeader(400)
	}
	resp := a.episodes.Query(keys[0])

	str, err2 := json.Marshal(resp)
	if err2 != nil {
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(str)
}

// Run starts the server
func (a *API) Run() {
	if err := http.ListenAndServe("0.0.0.0:8080", a.handler); err != nil {
		fmt.Println("unable to serve")
	}

}
