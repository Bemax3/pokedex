package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Bemax3/pokedex/internal"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	isInteractive := internal.IsTerminalInput()
	cfg := internal.NewConfig()

	for {

		if isInteractive {
			fmt.Print("Pokedex > ")
		}

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		cleaned := internal.CleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		matched := false
		for _, command := range internal.GetCommands() {
			if cleaned[0] == command.Name {

				if len(cleaned) > 1 {
					cfg.Arguments = cleaned[1:]
				}

				err := command.Callback(cfg)

				if err != nil {
					fmt.Println("An error occured while executing command :", err)
				}

				matched = true
				break
			}
		}

		if !matched {
			fmt.Println("Unknown command:", cleaned[0])
		}
	}
}
