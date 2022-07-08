package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

var addFunc func(string, string, string)

func Init(af func(string, string, string)) {
	addFunc = af
	Router = mux.NewRouter()
	Router.HandleFunc("/add", add).Methods("POST")
}

func Listen() {
	http.ListenAndServe(":5555", Router)
}

func add(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for a, b := range vars {
		fmt.Println(a, b)
	}
	if t, ok := vars["type"]; ok {
		fmt.Println(t)
		if d, ok := vars["date"]; ok {
			fmt.Println(d)
			if k, ok := vars["token"]; ok {
				fmt.Println(k)
				addFunc(t, d, k)
			}
		}
	}
}
