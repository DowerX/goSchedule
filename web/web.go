package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

var addFunc func(string, string, string, string)
var listFunc func() []byte

func Init(af func(string, string, string, string), lf func() []byte) {
	addFunc = af
	listFunc = lf
	Router = mux.NewRouter()
	Router.HandleFunc("/add", add).Methods("POST")
	Router.HandleFunc("/list", list).Methods("GET")
}

func Listen(address string) {
	http.ListenAndServe(address, Router)
}

func add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	addFunc(r.Form["type"][0], r.Form["from"][0], r.Form["to"][0], r.Form["token"][0])
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Write(listFunc())
}
