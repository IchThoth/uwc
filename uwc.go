package uwc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Unsplash Client version
const version = "1.0.0"

const (
	//Maximum amount of requests that can e directed to the Unsplash Api at once
	MaxAmountOfRequests = 32
	// specifies the duration of wait before you can make a new request to the
	//unsplash api
	defaultRetryDuration = time.Second * 10
)

// client options using optional function pattern
type ClientOptions func(client *Client)

type Client struct {
	httpclient    *http.Client
	BaseUrl       string
	RetryRequests bool
}

// Errors represents errors sent by the Unsplash Web API
type Error struct {
	// A short description of the error.
	Message string `json:"message"`
	// The HTTP status code.
	Status int `json:"status"`
}

func (e Error) Error() string {
	return e.Message
}

// Accepts a retry request
func AcceptRetry(retry bool) ClientOptions {
	return func(client *Client) {
		client.RetryRequests = retry
	}
}

func (c *Client) decodeError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return fmt.Errorf("unsplash: HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	// buffers create a seperates storage location in memory
	//for the response body
	buf := bytes.NewBuffer(body)

	var e struct {
		E Error `json:"error"`
	}
	err = json.NewDecoder(buf).Decode(e)
	if err != nil {
		return fmt.Errorf("unsplash: couldnt find any errors (%d), [%v]")
	}

	return e.E
}

// configures the original url for making a request
func ConfigureBaseUrl(url string) ClientOptions {
	return func(client *Client) {
		client.BaseUrl = url
	}
}

// New creates a new client to interact with the Unsplash API
func NewClient(client *http.Client, options ...ClientOptions) *Client {
	c := &Client{
		httpclient: client,
		BaseUrl:    "https://api.unsplash.com/",
	}
	for _, opt := range options {
		opt(c)
	}

	return c

}

// Creates a get request to the unsplash api
func (c *Client) Get(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// NonGetRequest to mae other requests other than a get request to the unsplash web ap
func (c *Client) NonGetRequest(ctx context.Context, url string, result interface{}) error {
	// dont no what to do here :)
}
