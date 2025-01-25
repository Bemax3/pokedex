package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Bemax3/pokedex/internal/pokecache"
)

func GetMapData(url string, cache *pokecache.Cache) (MapResult, error) {
	var data MapResult
	cachedData, exists := cache.Get(url)

	if exists {
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return MapResult{}, fmt.Errorf("Error unmarshalling cached data: %v", err)
		}
		return data, nil
	}

	data, err := fetchMaps(url)
	if err != nil {
		return MapResult{}, err
	}

	cachedResult, err := json.Marshal(data)
	if err != nil {
		return MapResult{}, fmt.Errorf("Error marshalling data for cache: %v", err)
	}

	cache.Add(url, cachedResult)
	return data, nil
}

func GetMapDetails(url string, cache *pokecache.Cache) (MapDetailsResult, error) {
	var data MapDetailsResult
	cachedData, exists := cache.Get(url)

	if exists {
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return MapDetailsResult{}, fmt.Errorf("Error unmarshalling cached data: %v", err)
		}
		return data, nil
	}

	data, err := fetchMapDetails(url)
	if err != nil {
		return MapDetailsResult{}, err
	}

	cachedResult, err := json.Marshal(data)
	if err != nil {
		return MapDetailsResult{}, fmt.Errorf("Error marshalling data for cache: %v", err)
	}

	cache.Add(url, cachedResult)
	return data, nil
}

func GetPokemon(url string, cache *pokecache.Cache) (Pokemon, error) {
	var data Pokemon
	cachedData, exists := cache.Get(url)

	if exists {
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return Pokemon{}, fmt.Errorf("Error unmarshalling cached data: %v", err)
		}
		return data, nil
	}

	data, err := fetchPokemon(url)
	if err != nil {
		return Pokemon{}, err
	}

	cachedResult, err := json.Marshal(data)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error marshalling data for cache: %v", err)
	}

	cache.Add(url, cachedResult)
	return data, nil
}

func fetchMaps(url string) (MapResult, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return MapResult{}, fmt.Errorf("Error while creating request: %v", err)
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return MapResult{}, fmt.Errorf("Error while getting resource: %v", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var data MapResult

	if err := decoder.Decode(&data); err != nil {
		return MapResult{}, fmt.Errorf("Error while parsing resource: %v", err)
	}

	return data, nil
}

func fetchMapDetails(url string) (MapDetailsResult, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return MapDetailsResult{}, fmt.Errorf("Error while creating request: %v", err)
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return MapDetailsResult{}, fmt.Errorf("Error while getting resource: %v", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var data MapDetailsResult

	if err := decoder.Decode(&data); err != nil {
		return MapDetailsResult{}, fmt.Errorf("Error while parsing resource: %v", err)
	}

	return data, nil
}

func fetchPokemon(url string) (Pokemon, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, fmt.Errorf("Error while creating request: %v", err)
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return Pokemon{}, fmt.Errorf("Error while getting resource: %v", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var data Pokemon

	if err := decoder.Decode(&data); err != nil {
		return Pokemon{}, fmt.Errorf("Error while parsing resource: %v", err)
	}

	return data, nil
}
