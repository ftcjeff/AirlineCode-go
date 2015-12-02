package main

import (
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
	tableName := "airline"
	tableFields := []string{"Id", "Name", "IATA", "ICAO", "CallSign", "Country", "Comments"}

	mysql := GetServiceURI("mysql")
	CreateTable(mysql, "picasso", "picasso", "picasso", tableName)

	dat, err := ioutil.ReadFile("airline_codes.csv")
	check(err)

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		if strings.Contains(line, ",") {
			tokens := strings.Split(line, ",")

			if len(tokens[3]) == 0 {
				continue
			}

			AddRow(tableName, tableFields, append(tokens[:2], tokens[3:]...))
		}
	}
}
