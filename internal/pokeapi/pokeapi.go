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
	if data, ok, err := GetCachedData[RespLocationAreas](c, url); err != nil {
		return RespLocationAreas{}, err
	} else if ok {
		return data, nil
	}

	// non cached data procedure
	data, err := GetRequest[RespLocationAreas](c, url)
	if err != nil {
		return data, err
	}

	return data, nil
}
func (c *Client) ExploreLocation(location string) (RespLocationDetail, error) {
	url := baseURL + "/" + areas + "/" + location

	// cached data test
	if data, ok, err := GetCachedData[RespLocationDetail](c, url); err != nil {
		return data, err
	} else if ok {
		return data, nil
	}

	// non cached data procedure
	data, err := GetRequest[RespLocationDetail](c, url)
	if err != nil {
		return data, err
	}
	return data, nil

}

func GetCachedData[T any](c *Client, key string) (T, bool, error) {
	var result T

	data, ok := c.cache.Get(key)
	if !ok {
		return result, false, nil
	}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, false, err
	}
	return result, true, nil
}

func GetRequest[T any](c *Client, url string) (T, error) {
	var result T

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	res, err := c.hC.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	d, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(d, &result)
	if err != nil {
		return result, err
	}
	c.cache.Add(url, d)
	return result, nil
}

func GetPokemons(det RespLocationDetail) []string {
	pkms := []string{}
	for _, p := range det.PokemonEncounters {
		pkms = append(pkms, p.Pokemon.Name)
	}
	return pkms
}
