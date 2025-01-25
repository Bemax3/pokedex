package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Bemax3/pokedex/internal/pokecache"
)

func GetMapData(url string, cache *pokecache.Cache) (MapResult, error) {
	return GetCachedData[MapResult](url, cache)
}

func GetMapDetails(url string, cache *pokecache.Cache) (MapDetailsResult, error) {
	return GetCachedData[MapDetailsResult](url, cache)
}

func GetPokemon(url string, cache *pokecache.Cache) (Pokemon, error) {
	return GetCachedData[Pokemon](url, cache)
}

func GetCachedData[T any](url string, cache *pokecache.Cache) (T, error) {
	var data T
	cachedData, exists := cache.Get(url)

	if exists {
		if err := json.Unmarshal(cachedData, &data); err != nil {
			return data, fmt.Errorf("Error unmarshalling cached data: %v", err)
		}
		return data, nil
	}

	data, err := fetchResource[T](url)
	if err != nil {
		return data, err
	}

	cachedResult, err := json.Marshal(data)
	if err != nil {
		return data, fmt.Errorf("Error marshalling data for cache: %v", err)
	}

	cache.Add(url, cachedResult)
	return data, nil
}

func fetchResource[T any](url string) (T, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return *new(T), fmt.Errorf("Error while creating request: %v", err)
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return *new(T), fmt.Errorf("Error while getting resource: %v", err)
	}
	return decodeJSON[T](res.Body)
}

func decodeJSON[T any](body io.ReadCloser) (T, error) {
	defer body.Close()
	var data T
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&data); err != nil {
		return data, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return data, nil
}
