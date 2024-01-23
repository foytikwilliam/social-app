package main

import (
	"log"
	"net/http"
	"social-app/data"
)

// IndexHandler is a simple handler for the index route
func index(writer http.ResponseWriter, request *http.Request) {
	gyms, err := data.GetGyms()
	if err != nil {
		error_message(writer, request, "Cannot get gyms")
	} else {
		_, err := session(writer, request)
		if err != nil {
			log.Println("Error checking session:", err)
			generateHTML(writer, nil, "layout", "public.navbar", "index")
		} else {
			log.Println("Gyms data:", gyms)
			generateHTML(writer, gyms, "layout", "private.navbar", "index")

		}
	}
}
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)

	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}

}
