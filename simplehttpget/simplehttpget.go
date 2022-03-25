// simplehttpget is a very minimal wrapper for simple queries. It adds a timeout
// and unifies http calls used for various supporting APIs

package simplehttpget

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Get returns the server's response, or an error if the process failed
func Get(url string) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get failed for url %q: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get failed (status code %d) for url %q", resp.StatusCode, url)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("get failed reading body for url %q: %w", url, err)
	}

	return body, nil
}
