/*
 * Nobel Prize Master Data
 *
 * Information about the Nobel Prizes and the Nobel Prize Laureates
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/2.1/",
		Index,
	},

	Route{
		"LaureateLaureateIDGet",
		strings.ToUpper("Get"),
		"/2.1/laureate/{laureateID}",
		LaureateLaureateIDGet,
	},

	Route{
		"LaureatesGet",
		strings.ToUpper("Get"),
		"/2.1/laureates",
		LaureatesGet,
	},

	Route{
		"NobelPrizeCategoryYearGet",
		strings.ToUpper("Get"),
		"/2.1/nobelPrize/{category}/{year}",
		NobelPrizeCategoryYearGet,
	},

	Route{
		"NobelPrizesGet",
		strings.ToUpper("Get"),
		"/2.1/nobelPrizes",
		NobelPrizesGet,
	},
}