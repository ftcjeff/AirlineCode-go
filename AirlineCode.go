package main

type AirlineCode struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IATA     string `json:"iata"`
	ICAO     string `json:"icao"`
	CallSign string `json:"callsign"`
	Country  string `json:"country"`
	Comments string `json:"comments"`
}

type AirlineCodes []AirlineCode
