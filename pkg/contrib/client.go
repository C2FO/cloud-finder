package contrib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client wraps the standard library's http.Client and makes it easier to work
// with for the types of requests we are making.
type Client struct {
	client  *http.Client
	headers map[string]string
	baseURL string
}

// NewClient creates and initializes a new client from an existing http.Client
func NewClient(c *http.Client) *Client {
	return &Client{
		client:  c,
		headers: make(map[string]string),
		baseURL: "",
	}
}

// SetHeader will set the header for all subsequent requests
func (c *Client) SetHeader(key, value string) {
	c.headers[key] = value
}

// SetBaseURL sets the base URL that will be prepended to every request
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) getURL(url string) string {
	return fmt.Sprintf("%s%s", c.baseURL, url)
}

// Get returns the body of the response as a string. It returns an error
// if it encounters one as well as if the HTTP response status code does not
// equal 200.
func (c *Client) Get(url string) (string, error) {
	fullURL := c.getURL(url)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("unable to create http request: %s", err)
	}

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received bad status code for response: %s", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body: %s", err)
	}

	return string(bytes), nil
}

// GetAll takes a map of named endpoints and processes all of them.
func (c *Client) GetAll(urls map[string]string) (map[string]string, error) {
	var err error
	responses := make(map[string]string)

	for key, url := range urls {
		responses[key], err = c.Get(url)
		if err != nil {
			return responses, err
		}
	}
	return responses, nil
}
