package backend

import (
	"fmt"
	"github.com/aubm/seven-poems/backend/webapi"
	"github.com/gorilla/mux"
	"net/http"
)

func StartOnAddr(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/poems", webapi.ListPoems)
	router.HandleFunc("/api/poems/{slug}", webapi.GetOnePoemBySlug)
	http.Handle("/", router)
	fmt.Println("Application started on " + addr)
	http.ListenAndServe(addr, nil)
}
