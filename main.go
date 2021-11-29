package main

import (
	router "fizzbuzzapi/Router"
	"net/http"
)

func main() {
	router := router.Router()
	http.ListenAndServe(":8080", router)
}
