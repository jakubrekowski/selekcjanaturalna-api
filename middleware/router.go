package middleware

import (
	"github.com/gorilla/mux"
	"github.com/jakubrekowski/selekcjanaturalna-api/endpoints"
	"net/http"
)

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/", endpoints.RouteOpenAPI).Methods("GET")

	router.HandleFunc("/users", endpoints.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", endpoints.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", endpoints.GetUsers).Methods("POST")
	router.HandleFunc("/user/{id}", endpoints.GetUsers).Methods("PUT")
	router.HandleFunc("/user/{id}", endpoints.GetUsers).Methods("DELETE")

	err := http.ListenAndServe(":5050", router)
	if err != nil {
		return
	}
}
