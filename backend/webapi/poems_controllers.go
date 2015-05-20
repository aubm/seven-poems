package webapi

import (
	"github.com/aubm/seven-poems/backend/poems"
	"io"
	"net/http"
)

func ListPoems(w http.ResponseWriter, r *http.Request) {
	poemsList := poems.GetPoemsList()
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, poemsList.ToJson())
}
