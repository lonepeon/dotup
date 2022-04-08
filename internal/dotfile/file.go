package dotfile

import (
	"errors"
	"path"
	"path/filepath"

	"github.com/lonepeon/dotup/internal/dotfile/fileutil"
)

var (
	// ErrNotExist is an error when the file doesn't exist
	ErrNotExist = errors.New("file doesn't exist")
)

// File represents a dotfile
type File struct {
	// Workdir is the root folder where dotfile are stored
	Workdir string
	// Filepath is the path to a specific dotfile, relative to the Workdir
	Filepath string
	// Path is the relative path to a specific dotfile. It includes the Workdir and Filepath
	Path string
}

// AbsolutePath is the path to the system's root folder
func (f File) AbsolutePath() (string, error) {
	return filepath.Abs(f.Path)
}

// RelativeFolder represents the folder where the file is
func (f File) RelativeFolder() string {
	return path.Join(f.Workdir, path.Dir(f.Filepath))
}

// NewSource represents a dotfile to copy
func NewSource(workdir string, fpath string) File {
	return newFile(workdir, fpath)
}

// NewDestination represents where a dotfile should be copied
func NewDestination(workdir string, src File) File {
	fpath := fileutil.MakeHidden(src.Filepath)
	return newFile(workdir, fpath)
}

func newFile(workdir, fpath string) File {
	return File{
		Workdir:  workdir,
		Filepath: fpath,
		Path:     path.Join(workdir, fpath),
	}
}
