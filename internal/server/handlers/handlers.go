package handlers

import (
	"net/http"
)

func GetCountries(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func GetCountry(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
