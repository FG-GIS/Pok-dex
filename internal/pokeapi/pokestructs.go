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

type RespPokemonDetail struct {
	Abilties []struct {
		Ability  pokeElement `json:"ability"`
		IsHidden bool        `json:"is_hidden"`
		Slot     int         `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms       []pokeElement `json:"forms"`
	GameIndices []struct {
		GameIndex int         `json:"game_index"`
		Version   pokeElement `json:"version"`
	} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []pokeElement `json:"held_items"`
	Id                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move                pokeElement `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int         `json:"level_learned_at"`
			MoveLearnMethod pokeElement `json:"move_learn_method"`
			VersionGroup    pokeElement `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string        `json:"name"`
	Order         int           `json:"order"`
	PastAbilities []pokeElement `json:"past_abilities"`
	PastTypes     []pokeElement `json:"past_types"`
	Species       pokeElement   `json:"species"`
	Sprites       struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int         `json:"base_stat"`
		Effort   int         `json:"effort"`
		Stat     pokeElement `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int         `json:"slot"`
		Type pokeElement `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type pokeElement struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Client struct {
	cache pokecache.Cache
	hC    http.Client
}
