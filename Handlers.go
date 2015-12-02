package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	allAirlines := GetAllRows("airline")

	if err := json.NewEncoder(w).Encode(allAirlines); err != nil {
		panic(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airlineCode := strings.Trim(vars["airlineCode"], " ")
	airlineInfo := GetRows("airline", "IATA", airlineCode)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airlineInfo); err != nil {
		panic(err)
	}
}

func ShowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := strings.Trim(vars["Id"], " ")
	airlineInfo := GetRows("airline", "Id", id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airlineInfo); err != nil {
		panic(err)
	}
}

func Country(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := strings.Trim(vars["country"], " ")
	airlineInfo := GetRows("airline", "Country", country)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airlineInfo); err != nil {
		panic(err)
	}
}
