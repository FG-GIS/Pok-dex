package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/FG-GIS/bootpokedex/internal/pokecache"
)

const (
	baseURL        = "https://pokeapi.co/api/v2"
	areas   string = "location-area"
)

type Client struct {
	cache pokecache.Cache
	hC    http.Client
}

type RespLocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(timeout, cacheLimit time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheLimit),
		hC: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocations(page *string) (RespLocationAreas, error) {
	url := baseURL + "/" + areas
	if page != nil {
		url = *page
	}

	// cached data test
	if data, ok := c.cache.Get(url); ok {
		locData := RespLocationAreas{}
		err := json.Unmarshal(data, &locData)
		if err != nil {
			return RespLocationAreas{}, err
		}
		return locData, nil
	}

	// non cached data procedure
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, nil
	}

	res, err := c.hC.Do(req)
	if err != nil {
		return RespLocationAreas{}, nil
	}
	defer res.Body.Close()

	d, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationAreas{}, nil
	}
	locs := RespLocationAreas{}
	err = json.Unmarshal(d, &locs)
	if err != nil {
		return RespLocationAreas{}, nil
	}
	c.cache.Add(url, d)

	return locs, nil
}
