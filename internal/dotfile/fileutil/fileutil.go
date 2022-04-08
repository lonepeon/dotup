// Package fileutil contains helper to handle dotfiles specifity like where to
// add the "." in a path
package fileutil

import (
	"path"
	"strings"
)

// MakeHidden adds a dot (.) prefix to the root folder or to the file if it's missing
func MakeHidden(fpath string) string {
	dir := path.Dir(fpath)
	filename := path.Base(fpath)

	if dir == "." {
		if strings.HasPrefix(filename, ".") {
			return filename
		}

		return "." + filename
	}

	if strings.HasPrefix(dir, ".") {
		return path.Join(dir, filename)
	}

	return path.Join("."+dir, filename)
}
