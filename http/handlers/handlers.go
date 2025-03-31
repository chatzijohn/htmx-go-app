package handlers

import (
	"net/http"
)

func CountryList(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func CountryDetail(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
