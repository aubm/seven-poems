package webapi

import (
	"github.com/aubm/seven-poems/backend/poems"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func ListPoems(w http.ResponseWriter, r *http.Request) {
	poemsList := poems.GetPoemsList()
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, poemsList.ToJson())
}

func GetOnePoemBySlug(w http.ResponseWriter, r *http.Request) {
	urlParameters := mux.Vars(r)
	poem, err := poems.GetOnePoemBySlug(urlParameters["slug"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, poem.ToJson())
	}
}
