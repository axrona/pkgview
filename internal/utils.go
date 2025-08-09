package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

// getAvailableEditor returns the first available editor from a prioritized list.
// It searches in PATH for vi, nano, vim, nvim, jed, in that order.
// Returns error if none are found.
func getAvailableEditor() (string, error) {
	editors := []string{"vi", "nano", "vim", "nvim", "jed"}

	for _, editor := range editors {
		path, err := exec.LookPath(editor)
		if err == nil {
			return path, nil
		}
	}

	return "", errors.New("no supported editor found in PATH; please set the EDITOR environment variable to your preferred editor")
}

// getEditor returns the editor to use.
// It first tries the EDITOR environment variable and checks if it's executable.
// If EDITOR is unset or invalid, falls back to getAvailableEditor().
func getEditor() (string, error) {
	editorEnv := os.Getenv("EDITOR")
	if _, err := exec.LookPath(editorEnv); err != nil {
		defaultEditor, err := getAvailableEditor()
		if err != nil {
			return "", err
		}

		return defaultEditor, nil
	}

	return editorEnv, nil
}

// GetPKGBUILD downloads the PKGBUILD file for the given AUR package name.
// It URL-escapes the package name and performs an HTTP GET request.
// Returns the PKGBUILD content as bytes or an error if not found or HTTP error occurs.
func GetPKGBUILD(pkgName string) ([]byte, error) {
	pkgName = url.QueryEscape(pkgName)
	pkgbuildURL := fmt.Sprintf("https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=%s", pkgName)
	resp, err := http.Get(pkgbuildURL)
	if err != nil {
		return nil, err
	}

	// Handle package not found (404) error explicitly.
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("%s package not found", pkgName)
	}

	// Handle other HTTP errors.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	pkgbuild, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return pkgbuild, nil
}

// OpenWithEditor writes the PKGBUILD bytes to a temporary file,
// then opens it in the user's preferred editor.
// The editor is run interactively, inheriting stdin/stdout/stderr.
// Returns error if file creation, writing, or editor execution fails.
func OpenWithEditor(pkgbuild []byte) error {
	tmpfile, err := os.CreateTemp("", "pkgbuild-*")
	if err != nil {
		return err
	}

	if _, err := tmpfile.Write(pkgbuild); err != nil {
		return err
	}

	if err := tmpfile.Close(); err != nil {
		return err
	}

	editor, err := getEditor()
	if err != nil {
		return err
	}

	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
