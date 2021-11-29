package router

import (
	middleware "fizzbuzzapi/Middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/fizzbuzz/{count}", middleware.FizzbuzzFunction).Methods("GET", "OPTIONS")

	return router
}
