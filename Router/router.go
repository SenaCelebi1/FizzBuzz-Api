package router

import (
	service "fizzbuzzapi/Service"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/fizzbuzz/{count}", service.FizzbuzzFunction).Methods("GET", "OPTIONS")

	return router
}
