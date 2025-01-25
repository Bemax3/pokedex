package internal

import (
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func IsTerminalInput() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error detecting input mode:", err)
		return false
	}
	// Check if stdin is a terminal
	return (info.Mode() & os.ModeCharDevice) != 0
}
