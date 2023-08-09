package uwc

import (
	"net/http"
	"time"
)

// Unsplash Client version
const version = "1.0.0"

const (
	//Maximum amount of requests that can e directed to the Unsplash Api at once
	MaxAmountOfRequests = 32
	// specifies the duration of wait before yyou can make a new reqquest to the
	//unsplash api
	defaultRetryDuration = time.Second * 10
)

// client options
type ClientOpts func(client *Client)

type Client struct {
	httpclient    *http.Client
	BaseUrl       string
	RetryRequests bool
}

func AcceptRetry(retry bool) ClientOpts {
	return func(client *Client) {
		client.RetryRequests = true
	}
}

func ConfigureBaseUrl(url string) ClientOpts {
	return func(client *Client) {
		client.BaseUrl = url
	}
}

// New creates a new client to interact with the Unsplash API
func New(client *http.Client, opts ...ClientOpts) *Client {
	c := &Client{
		httpclient: client,
		BaseUrl:    "https://api.unsplash.com/",
	}
	for _, opt := range opts {
		opt(c)
	}

	return c

}
