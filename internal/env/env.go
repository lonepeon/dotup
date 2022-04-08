// Package env contains function in charge of reading / parsing environment variables
package env

import (
	"fmt"
	"os"
	"path/filepath"
)

// Homedir returns the absolute path to the home directory based on environment variables
func Homedir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("can't get $HOME environment variable")
	}

	return filepath.Abs(home)
}
