package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/i-bielik/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreas, error) {
	// Simulate fetching location areas from an API
	// In a real implementation, this would involve making an HTTP request to the PokeAPI
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, found := c.cache.Get(url); found {
		var locations LocationAreas
		err := json.Unmarshal(cachedData, &locations)
		if err != nil {
			log.Fatal(err)
		}
		return locations, nil
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

	c.cache.Add(url, body)
	return locations, nil
}

func (c *Client) GetLocationArea(locationName string) (LocationArea, error) {
	// Simulate fetching location areas from an API
	// In a real implementation, this would involve making an HTTP request to the PokeAPI
	if locationName == "" {
		return LocationArea{}, errors.New("location cannot be empty")
	}

	url := baseURL + "/location-area/" + locationName

	if cachedData, found := c.cache.Get(url); found {
		var location LocationArea
		err := json.Unmarshal(cachedData, &location)
		if err != nil {
			log.Fatal(err)
		}
		return location, nil
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
	var location LocationArea
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatal(err)
	}

	c.cache.Add(url, body)
	return location, nil
}

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	// Simulate fetching location areas from an API
	// In a real implementation, this would involve making an HTTP request to the PokeAPI
	if pokemonName == "" {
		return Pokemon{}, errors.New("pokemon name cannot be empty")
	}

	url := baseURL + "/pokemon/" + pokemonName

	if cachedData, found := c.cache.Get(url); found {
		var pokemon Pokemon
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			log.Fatal(err)
		}
		return pokemon, nil
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
	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatal(err)
	}

	c.cache.Add(url, body)
	return pokemon, nil
}
