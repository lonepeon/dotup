package dotfile

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// WatcherFunc is used in order to watch each operation
type WatcherFunc func(src File, dest File, err error)

// NoopWatcher is a  watcher function doing nothing
func NoopWatcher(src File, dest File, err error) {}

// CopyTree goes through all file and subfolders of the src working directory and copies all files to the dest working directory using the copy fn
func CopyTree(srcWorkdir string, destWorkdir string, fn CopyFunc, watcher WatcherFunc) error {
	return filepath.Walk(srcWorkdir, func(fpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(path.Base(fpath), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if info.IsDir() {
			return nil
		}

		relfpath := strings.TrimPrefix(strings.TrimPrefix(fpath, srcWorkdir), "/")
		src := NewSource(srcWorkdir, relfpath)
		dest := NewDestination(destWorkdir, src)
		err = fn(src, dest)

		watcher(src, dest, err)

		return err
	})
}
