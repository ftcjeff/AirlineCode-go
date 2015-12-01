package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var airlineCodes AirlineCodes

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	dat, err := ioutil.ReadFile("airline_codes.csv")
	check(err)

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		if strings.Contains(line, ",") {
			tokens := strings.Split(line, ",")

			go RepoCreateAirlineCode(
				AirlineCode{Id: tokens[0],
					Name:     tokens[1],
					IATA:     tokens[3],
					ICAO:     tokens[4],
					CallSign: tokens[5],
					Country:  tokens[6],
					Comments: tokens[7]})
		}
	}
}

func RepoFindAirlineCodesByCountry(country string) AirlineCodes {
	var rv AirlineCodes

	for _, t := range airlineCodes {
		if t.Country == country {
			rv = append(rv, t)
		}
	}

	return rv
}

func RepoFindAirlineId(id string) AirlineCode {
	for _, t := range airlineCodes {
		if t.Id == id {
			return t
		}
	}

	return AirlineCode{}
}

func RepoFindAirlineCode(iata string) AirlineCode {
	for _, t := range airlineCodes {
		if t.IATA == iata {
			return t
		}
	}

	return AirlineCode{}
}

func RepoCreateAirlineCode(t AirlineCode) AirlineCode {
	fmt.Println("Creating ", t)
	airlineCodes = append(airlineCodes, t)
	return t
}

func RepoDestroyAirlineCode(iata string) error {
	for i, t := range airlineCodes {
		if t.IATA == iata {
			airlineCodes = append(airlineCodes[:i], airlineCodes[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find AirlineCode with id of %s to delete", iata)
}
