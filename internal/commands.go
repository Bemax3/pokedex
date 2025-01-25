package internal

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/Bemax3/pokedex/internal/api"
	"github.com/Bemax3/pokedex/internal/pokecache"
)

type Config struct {
	Next          string
	Previous      string
	MapDetailsUrl string
	PokemonUrl    string
	Cache         *pokecache.Cache
	Pokedex       api.Pokedex
	Arguments     []string
}

func NewConfig() *Config {
	return &Config{
		Next:          "https://pokeapi.co/api/v2/location-area",
		Previous:      "",
		MapDetailsUrl: "https://pokeapi.co/api/v2/location-area/%v",
		PokemonUrl:    "https://pokeapi.co/api/v2/pokemon/%v",
		Cache:         pokecache.NewCache(time.Second * 10),
		Arguments:     []string{},
		Pokedex:       make(api.Pokedex, 10),
	}
}

type Command struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    help,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 maps",
			Callback:    mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous maps",
			Callback:    mapBackCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a map area",
			Callback:    explore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    catch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon",
			Callback:    inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List your pokedex",
			Callback:    pokedex,
		},
	}
}

func exit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func help(cfg *Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage: \n\n")

	commands := GetCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	return nil
}

func mapCommand(cfg *Config) error {

	data, err := api.GetMapData(cfg.Next, cfg.Cache)

	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, pmap := range data.Results {
		fmt.Println(pmap.Name)
	}

	return nil
}

func mapBackCommand(cfg *Config) error {

	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
	}

	data, err := api.GetMapData(cfg.Previous, cfg.Cache)

	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, pmap := range data.Results {
		fmt.Println(pmap.Name)
	}

	return nil
}

func explore(cfg *Config) error {

	if len(cfg.Arguments) == 0 {
		return fmt.Errorf("Please provide a map area to explore")
	}

	data, err := api.GetMapDetails(fmt.Sprintf(cfg.MapDetailsUrl, cfg.Arguments[0]), cfg.Cache)

	if err != nil {
		return err
	}

	for _, encounter := range data.PokemonEncounters {
		fmt.Printf("- %v\n", encounter.Pokemon.Name)
	}

	return nil
}

func catch(cfg *Config) error {

	if len(cfg.Arguments) == 0 {
		return fmt.Errorf("Please provide a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", cfg.Arguments[0])

	data, err := api.GetPokemon(fmt.Sprintf(cfg.PokemonUrl, cfg.Arguments[0]), cfg.Cache)

	if err != nil {
		return err
	}

	difficulty := math.Min(float64(data.BaseExperience)/500, 1.0)

	catchRate := 1.0 - difficulty

	random := rand.Float64()

	if random < catchRate {
		fmt.Printf("%v was caught!\n", cfg.Arguments[0])
		cfg.Pokedex[data.Name] = data
	} else {
		fmt.Printf("%v escaped!\n", cfg.Arguments[0])
	}

	return nil
}

func inspect(cfg *Config) error {

	if len(cfg.Arguments) == 0 {
		return fmt.Errorf("Please provide a pokemon name")
	}

	pokemon, exists := cfg.Pokedex[cfg.Arguments[0]]

	if !exists {
		return fmt.Errorf("You have not caught that pokemon")
	}

	pokemon.Print()

	return nil
}

func pokedex(cfg *Config) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("- %v\n", pokemon.Name)
	}

	return nil
}
