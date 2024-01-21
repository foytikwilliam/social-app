package main

import (
	"fmt"
	"net/http"
	"social-app/data"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// Handle parse form error
		fmt.Println("Error parsing form:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user, err := data.UserByEmail(email)
	if err != nil {
		// Handle user retrieval error
		fmt.Println("Error retrieving user:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Check if the user exists before accessing its properties
	if user == (data.User{}) || user.Password != data.Encrypt(password) {
		fmt.Println("Invalid email or password")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	session, err := user.CreateSession()
	if err != nil {
		// Handle session creation error
		fmt.Println("Error creating session:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "layout", "public.navbar", "login")
}
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "signup")
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println("Cannot parse form", err)
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		fmt.Println("Cannot create user", err)
	} else {
		fmt.Println("user succesfully created", err)
	}
	http.Redirect(writer, request, "/", http.StatusFound)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != nil && err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
	}

	if err == nil {
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}
