// utils.go

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"social-app/data"
	"strconv"
	"strings"
	"text/template"
)

var logger *log.Logger

// renderTemplate is a helper function to render HTML templates
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	fmt.Println("Hello from session!", nil)
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), http.StatusFound)
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func parseRating(value string) int {
	// Check if the string is empty
	if value == "" {
		// Handle the empty string case; return a default value, or log it
		return 0
	}

	// Parse the value to an integer; handle errors if needed
	rating, err := strconv.Atoi(value)
	if err != nil {
		// Handle the error, return a default value, or log it
		return 0
	}
	return rating
}
