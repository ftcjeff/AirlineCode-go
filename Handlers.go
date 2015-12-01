package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airlineCodes); err != nil {
		panic(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airlineCode := strings.Trim(vars["airlineCode"], " ")
	airlineInfo := RepoFindAirlineCode(airlineCode)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airlineInfo); err != nil {
		panic(err)
	}
}

func ShowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := strings.Trim(vars["Id"], " ")
	airlineInfo := RepoFindAirlineId(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airlineInfo); err != nil {
		panic(err)
	}
}

func Country(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := strings.Trim(vars["country"], " ")
	airlineInfo := RepoFindAirlineCodesByCountry(country)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(airlineInfo); err != nil {
		panic(err)
	}
}

func Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airlineCode := strings.Trim(vars["airlineCode"], " ")
	RepoDestroyAirlineCode(airlineCode)
	fmt.Fprintln(w, "AirlineCode Removed [", airlineCode, "]")
}

func Create(w http.ResponseWriter, r *http.Request) {
	var airlineCode AirlineCode
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &airlineCode); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateAirlineCode(airlineCode)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
