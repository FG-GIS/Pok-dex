package pokeapi

import (
	"net/http"

	"github.com/FG-GIS/bootpokedex/internal/pokecache"
)

type RespLocationAreas struct {
	Count    int           `json:"count"`
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []pokeElement `json:"results"`
}

type RespLocationDetail struct {
	EncounterMethodRates []struct {
		EncounterMethod pokeElement `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int         `json:"rate"`
			Version pokeElement `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int         `json:"game_index"`
	Id        int         `json:"id"`
	Location  pokeElement `json:"location"`
	Name      string      `json:"name"`
	Names     []struct {
		Language pokeElement `json:"language"`
		Name     string      `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        pokeElement `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int         `json:"chance"`
				ConditionValues []struct{}  `json:"condition_values"`
				MinLevel        int         `json:"min_level"`
				MaxLevel        int         `json:"max_level"`
				Method          pokeElement `json:"method"`
			} `json:"encounter_details"`
			MaxChance int         `json:"max_chance"`
			Version   pokeElement `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type pokeElement struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Client struct {
	cache pokecache.Cache
	hC    http.Client
}
