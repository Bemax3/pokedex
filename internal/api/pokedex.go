package api

import "fmt"

type Pokemap struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type MapResult struct {
	Count    int
	Next     string
	Previous string
	Results  []Pokemap
}

type PokemonDetails struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Pokemon PokemonDetails `json:"pokemon"`
}

type MapDetailsResult struct {
	PokemonEncounters []PokemonList `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (pokemon *Pokemon) Print() {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Name)
	fmt.Printf("Weight: %v\n", pokemon.Name)
	fmt.Println("Stats:")

	for _, v := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", v.Stat.Name, v.BaseStat)
	}

	fmt.Println("Types:")

	for _, v := range pokemon.Types {
		fmt.Printf("  - %v\n", v.Type.Name)
	}
}

type Pokedex map[string]Pokemon
