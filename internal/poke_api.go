package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreas, error) {
	// Simulate fetching location areas from an API
	// In a real implementation, this would involve making an HTTP request to the PokeAPI
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	locations := LocationAreas{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}

	return locations, nil
}
