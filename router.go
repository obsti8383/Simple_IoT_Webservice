package main

import (
	mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Method string
	Path   string
	Handle mux.Handle // httprouter package as mux
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/device",
		GetDevices,
	},
	Route{
		"GET",
		"/device/:id",
		GetDeviceId,
	},
	Route{
		"POST",
		"/device",
		PostDevice,
	},
	Route{
		"DELETE",
		"/device/:id",
		DeleteDeviceId,
	},
}

func NewRouter() *mux.Router {
	router := mux.New()
	for _, route := range routes {

		router.Handle(route.Method, route.Path, route.Handle)

	}

	return router
}
