package skyroom

import (
	"net/http"
	"net/url"
)

const (
	endpoint = "https://www.skyroom.online/skyroom/api/"
)

type Skyroom struct {
	APIKey  string
	BaseURL *url.URL
	Client  *http.Client
}

func New(apiKey string) *Skyroom {
	u, _ := url.Parse(endpoint)

	return &Skyroom{
		BaseURL: u,
		APIKey:  apiKey,
		Client:  http.DefaultClient,
	}
}
