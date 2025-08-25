package cmd

import (
	"fmt"
	"os"
	"strings"

	utils "github.com/axrona/pkgview/internal"
)

// Run parses CLI arguments and dispatches the corresponding actions.
func Run() error {
	// If no arguments provided, show help and exit.
	if len(os.Args) < 2 {
		printHelp()
		return nil
	}

	// Handle the first argument as command or package name.
	switch os.Args[1] {
	case "-h", "--help":
		// Show help message and exit.
		printHelp()
		return nil
	case "completion", "completions":
		// Check that shell argument is provided.
		if len(os.Args) < 3 {
			return fmt.Errorf("completion: shell arg missing (bash/zsh/fish)")
		}

		// Print shell completion script for the given shell.
		shell := os.Args[2]
		printCompletion(shell)
		return nil
	default:
		// Treat the argument as a package name and attempt to view it.
		pkg := os.Args[1]
		err := viewPackage(pkg)
		if err != nil {
			return err
		}

		return nil
	}
}

// printHelp outputs the usage instructions and available commands.
func printHelp() {
	fmt.Println(`
Usage: pkgview [OPTIONS] <package_name>
    pkgview completion <shell>

Options:
  -h, --help          Show this help message and exit

Subcommands:
  completion <shell>  Output shell completion script for one of: bash, zsh, fish

Arguments:
  <package_name>      Name of the AUR package to view its PKGBUILD file

Examples:
  pkgview yay
  pkgview --help
  pkgview completion bash`)
}

// printCompletion prints the completion script for the specified shell.
func printCompletion(shell string) {
	fmt.Println(completions[strings.ToLower(shell)])
}

// viewPackage fetches the PKGBUILD file of the given package and opens it in editor.
func viewPackage(pkg string) error {
	// Retrieve PKGBUILD content.
	pkgbuild, err := utils.GetPKGBUILD(pkg)
	if err != nil {
		return err
	}

	// Open the content in the preferred editor.
	err = utils.OpenWithEditor(pkgbuild)
	if err != nil {
		return err
	}

	return nil
}
