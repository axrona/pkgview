package main

import (
	"fmt"

	"github.com/axrona/pkgview/internal/cmd"
)

// main is the entry point of the application.
// It runs the command logic from cmd.Run and prints any error returned.
func main() {
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
	}
}
