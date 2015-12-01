package ftcpops

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/airline",
		Index,
	},

	Route{
		"Show",
		"GET",
		"/airline/{airlineCode}",
		Show,
	},

	Route{
		"Country",
		"GET",
		"/country/{country}",
		Country,
	},

	Route{
		"Remove",
		"DELETE",
		"/airline/{airlineCode}",
		Remove,
	},

	Route{
		"Create",
		"POST",
		"/airline",
		Create,
	},
}
