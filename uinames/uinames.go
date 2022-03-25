// uinames provides a first and last name (in this case, from uinames.com)
// api doc: https://github.com/thm/uinames

package uinames

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/swayne275/joke-web-server/simplehttpget"
)

const (
	base            = "http://uinames.com/api"
	genericErrorMsg = "could not generate a new name"
)

// fullName helps unmarshal a successful request into a struct with bits we care about
type fullName struct {
	// we don't care about the gender or region fields
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
}

// use standard logger with the package name prepended
func logErr(err error) {
	log.Printf("names pkg: %v\n", err)
}

// optionally append gender to the URL
func getGenderedURL(gender string) string {
	if gender != "" {
		return base + fmt.Sprintf("?gender=%s", url.QueryEscape(gender))
	}

	return base
}

// parseNames returns the {firstName, lastName} from the name server, or error if anything went wrong
func parseNames(data []byte) (string, string, error) {
	name := fullName{}
	if err := json.Unmarshal(data, &name); err != nil {
		// we likely got an error message, as per the API docs
		logErr(fmt.Errorf("error marshalling name from api: %w", err))
		return "", "", fmt.Errorf(genericErrorMsg)
	}

	return name.FirstName, name.LastName, nil
}

// GetNew returns a new {firstName, lastName}, or error if anything went wrong
func GetNew(gender string) (string, string, error) {
	body, err := simplehttpget.Get(getGenderedURL(gender))
	if err != nil {
		logErr(fmt.Errorf("uinames api get failed: %w", err))
		return "", "", fmt.Errorf(genericErrorMsg)
	}

	return parseNames(body)
}
