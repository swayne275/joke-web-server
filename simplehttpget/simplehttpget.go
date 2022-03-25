// simplehttpget is a very minimal wrapper for simple queries. It adds a timeout
// and unifies http calls used for various supporting APIs

package simplehttpget

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// Get returns the server's response, or an error if the process failed
func Get(url string) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("get failed for url %s", url))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("get failed (status code %d) for url %s",
											resp.StatusCode, url))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("get failed reading body for url %s", url))
	}

	return body, nil
}