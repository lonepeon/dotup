package dotfiletest

import "github.com/lonepeon/dotup/internal/dotfile"

// CopyOperation represent a copy of Source file to Destination file
type CopyOperation struct {
	Source      dotfile.File
	Destination dotfile.File
	Err         error
}

// NoOpCopy is a dotfile.CopyFunc doing nothing
func NoOpCopy(src dotfile.File, dest dotfile.File) error {
	return nil
}

// WatchFileCopy is a dotfile.WatcherFunc operation in charge of storing each operation in a slice for later use
func WatchFileCopy(ops *[]CopyOperation) dotfile.WatcherFunc {
	return func(src dotfile.File, dest dotfile.File, err error) {
		*ops = append(*ops, CopyOperation{Source: src, Destination: dest, Err: err})
	}
}
