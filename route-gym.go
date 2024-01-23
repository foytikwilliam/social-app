package main

import (
	"net/http"
	"social-app/data"
)

func readGymReview(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	gym, err := data.GymByUUID(uuid)
	if err != nil {
		error_message(writer, request, "Cannot read gym")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &gym, "layout", "public.navbar", "public.gym")
		} else {
			generateHTML(writer, &gym, "layout", "private.navbar", "public.gym")
		}
	}
}
