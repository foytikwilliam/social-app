package main

import (
	"fmt"
	"net/http"
	"social-app/data"
)

type PageData struct {
	Gyms    []data.Gym
	Reviews []data.Review
}

func createReview(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusFound)
		return
	}

	// Parse the form
	err = request.ParseForm()
	if err != nil {
		fmt.Println(err, "Cannot parse form")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Get the current user from the session
	user, err := sess.User()
	if err != nil {
		fmt.Println(err, "Cannot get user from session")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Get the review body from the form
	body := request.PostFormValue("body")

	// Assuming you have the gym ID or gym UUID from the request
	gymUUID := request.PostFormValue("uuid")
	// Assuming you have a function to get the Gym by ID or UUID
	gym, err := data.GymByUUID(gymUUID)
	if err != nil {
		fmt.Println(err, "Cannot get gym by ID or UUID")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Create the review for the user and the specified gym
	if _, err := user.CreateReview(gym, body); err != nil {
		fmt.Println(err, "Cannot create review")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Redirect the user to the home page or the gym details page
	http.Redirect(writer, request, "/", http.StatusFound)
}

func readGymReview(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	gymUUID := vals.Get("id")

	// Check if gymUUID parameter is present
	if gymUUID == "" {
		error_message(writer, request, "Missing gym UUID parameter")
		return
	}

	// Assuming you have a function to get the Gym by UUID
	gym, err := data.GymByUUID(gymUUID)
	if err != nil {
		// Handle error (unable to fetch gym by UUID)
		error_message(writer, request, "Invalid gym UUID: "+err.Error())
		return
	}
	println(gym.Id)
	// Assuming ReviewsByGymID returns a slice of reviews for the given gym ID
	reviews, err := data.ReviewsByGymID(gym.Id)

	pageData := PageData{
		Gyms:    []data.Gym{gym}, // Assuming you want to pass a single gym
		Reviews: reviews,
	}
	fmt.Printf("Number of Reviews: %d\n", len(pageData.Reviews))

	fmt.Println("Gym:", gym)
	fmt.Println("Review:", reviews)
	if err != nil {
		error_message(writer, request, "Cannot read gym reviews")
	} else {
		// Assuming you want to render the reviews in the HTML template
		_, err := session(writer, request)

		if err != nil {
			// Assuming you want to display the public navbar for non-authenticated users
			generateHTML(writer, &pageData, "layout", "public.navbar", "public.gym")
		} else {
			// Assuming you want to display the private navbar for authenticated users
			generateHTML(writer, &pageData, "layout", "private.navbar", "public.gym")
		}
	}
}

func postReview(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusFound)
		return
	}

	// Parse the form
	err = request.ParseForm()
	if err != nil {
		fmt.Println(err, "cannot parse form")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Get the current user from the session
	user, err := sess.User()
	if err != nil {
		fmt.Println(err, "cannot get user from session")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Get the review body from the form
	body := request.PostFormValue("body")

	gymUUID := request.FormValue("uuid")
	// Assuming you have a function to get the Gym by UUID
	gym, err := data.GymByUUID(gymUUID)
	if err != nil {
		fmt.Println(err, "Cannot get gym by UUID")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Create the review for the user
	if _, err := user.CreateReview(gym, body); err != nil {
		fmt.Println(err, "Cannot create review")
		// Handle the error accordingly
		// You might want to display an error message or redirect the user
		return
	}

	// Redirect the user to a success page or the gym details page
	// Redirect the user to the gym details page
	http.Redirect(writer, request, "/gym/read?id="+gymUUID, http.StatusFound)

}

func newReview(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusFound)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.review")
	}
}
