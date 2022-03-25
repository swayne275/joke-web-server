// randomuser provides a first and last name (in this case, from randomuser.me)
// api doc: https://randomuser.me
// this is in place because uinmaes.com was down (403 status code) when I wrote this

package randomuser

import (
	"fmt"
	"log"
	"net/url"

	"github.com/pkg/errors"
	"github.com/swayne275/joke-web-server/simplehttpget"
	"github.com/tidwall/gjson"
)

const (
	// lock api version to 1.3 to avoid breaking changes. Use default JSON format
	base            = "https://randomuser.me/api/1.3"
	genericErrorMsg = "could not generate a new name"
)

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
	firstNameResult := gjson.GetBytes(data, "results.0.name.first")
	if !firstNameResult.Exists() {
		logErr(errors.New("no name.first given by API"))
		return "", "", errors.New(genericErrorMsg)
	}

	lastNameResult := gjson.GetBytes(data, "results.0.name.last")
	if !lastNameResult.Exists() {
		logErr(errors.New("no name.last given by API"))
		return "", "", errors.New(genericErrorMsg)
	}

	return firstNameResult.String(), lastNameResult.String(), nil
}

// GetNew returns a new {firstName, lastName}, or error if anything went wrong
func GetNew(gender string) (string, string, error) {
	body, err := simplehttpget.Get(getGenderedURL(gender))
	if err != nil {
		logErr(errors.Wrap(err, "randomuser api get failed"))
		return "", "", errors.New(genericErrorMsg)
	}

	return parseNames(body)
}
