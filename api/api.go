// Provides the client API. Possible extensions include middleware and API versioning.

package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/swayne275/joke-web-server/jokes"

	// If the original names API starts working again, swap these two imports
	//names "github.com/swayne275/joke-web-server/uinames"
	names "github.com/swayne275/joke-web-server/randomuser"
)

const (
	genericErrorMsg = "Unable to generate custom joke"
)

// use standard logger with the package name prepended
func logErr(err error) {
	log.Printf("api pkg: %v\n", err)
}

// setup Cross-Origin Resource Sharing in handler for browser clients
func setupCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding")
}

func handleGetJoke(w http.ResponseWriter, r *http.Request) {
	setupCors(w)
	if strings.ToLower((*r).Method) == "options" {
		// handle browser cors preflight request
		return
	}

	// use male gender so the joke (if gendered) makes sense with the name
	firstName, lastName, err := names.GetNew("male")
	if err != nil {
		logErr(errors.Wrap(err, "error getting new name"))
		http.Error(w, genericErrorMsg, http.StatusInternalServerError)
		return
	}

	joke, err := jokes.GetNew(firstName, lastName)
	if err != nil {
		logErr(errors.Wrap(err, "error getting new joke"))
		http.Error(w, genericErrorMsg, http.StatusInternalServerError)
		return
	}

	// Content-Type header is set automatically
	if _, err = w.Write([]byte(joke)); err != nil {
		logErr(errors.Wrap(err, "error writing response to client"))
		http.Error(w, genericErrorMsg, http.StatusInternalServerError)
		return
	}
}

// StartServer starts the webserver to deliver jokes at /
func StartServer(port string) error {
	// handler is run in a separate routine for each request
	http.HandleFunc("/", handleGetJoke)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return errors.Wrapf(err, fmt.Sprintf("Could not start client API server on port %s", port))
	}

	return nil
}
