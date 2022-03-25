// jokes provides a new joke specific to a given name (from icndb, in this case)
// api doc: http://www.icndb.com/api/

package jokes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/swayne275/joke-web-server/simplehttpget"
)

const (
	// use the default html encoding of the joke
	base            = "http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=[nerdy]&escape=javascript"
	genericErrorMsg = "could not generate a new joke"
)

// response represents the response from the joke service
type response struct {
	ResponseType  string `json:"type"`
	ResponseValue value  `json:"value"`
}

// value represents the value object in the joke service response with bits we care about
type value struct {
	// we don't need the id or categories fields for this use case
	Joke string `json:"joke"`
}

// use standard logger with the package name prepended
func logErr(err error) {
	log.Printf("jokes pkg: %v\n", err)
}

// return the query url for the given first/last name. Error if name is blank
func getNameQueryURL(first, last string) (string, error) {
	if first == "" {
		return "", errors.New("empty first name")
	}
	if last == "" {
		return "", errors.New("empty last name")
	}

	firstEscaped := url.QueryEscape(first)
	lastEscaped := url.QueryEscape(last)

	return fmt.Sprintf(base, firstEscaped, lastEscaped), nil
}

// parseJoke returns the joke from the joke server, error if anything went wrong
func parseJoke(data []byte) (string, error) {
	response := response{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		logErr(errors.Wrap(err, "couldn't parse joke"))
		return "", errors.New(genericErrorMsg)
	}

	if strings.ToLower(response.ResponseType) != "success" {
		// API doc doesn't cover what happens if !"success",
		// so I'll assume the structure is similar enough that this won't cause problems
		logErr(errors.New(fmt.Sprintf("unsuccessful API response: %s", response.ResponseType)))
		return "", errors.New(genericErrorMsg)
	}

	return response.ResponseValue.Joke, nil
}

// GetNew gets a new joke using the supplied first and last name and returns
// an error if invalid data is provided, or if the backing service is down
func GetNew(first, last string) (string, error) {
	url, err := getNameQueryURL(first, last)
	if err != nil {
		logErr(errors.Wrap(err, "invalid name query parameters"))
		return "", errors.New(genericErrorMsg)
	}

	body, err := simplehttpget.Get(url)
	if err != nil {
		logErr(errors.Wrap(err, "jokes api get failed"))
		return "", errors.New(genericErrorMsg)
	}

	return parseJoke(body)
}
