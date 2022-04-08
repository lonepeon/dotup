package dotfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFunc represents a function in charge of copying the content of src to dest
type CopyFunc func(src File, dest File) error

// SymlinkFile creates a symbolic link between the src file and the dest, creating all necessary folders
func SymlinkFile(src File, dest File) error {
	if _, err := os.Stat(src.Path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("src file doesn't exist (path=%s) : %w", src.Path, ErrNotExist)
		}

		return fmt.Errorf("can't get src file stat (path=%s): %v", src.Path, err)
	}

	srcPath, err := filepath.Abs(src.Path)
	if err != nil {
		return fmt.Errorf("can't build absolute path from src file (path=%s) : %w: %v", src.Path, ErrNotExist, err)
	}

	if err := os.MkdirAll(dest.RelativeFolder(), 0755); err != nil {
		return fmt.Errorf("can't create dest folder (folder=%s path=%s) : %v", dest.RelativeFolder(), dest.Path, err)
	}

	if _, err := os.Stat(dest.Path); err == nil {
		if err := os.Remove(dest.Path); err != nil {
			return fmt.Errorf("can't delete existing dest file (path=%s): %v", dest.Path, err)
		}
	}

	return os.Symlink(srcPath, dest.Path)
}

// DuplicateFile copies the src file content to its dest file, creating all necessary folders
func DuplicateFile(src File, dest File) error {
	statSrc, err := os.Stat(src.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("src file doesn't exist (path=%s) : %w", src.Path, ErrNotExist)
		}

		return fmt.Errorf("can't get src file stat (path=%s): %v", src.Path, err)
	}

	if !statSrc.Mode().IsRegular() {
		return fmt.Errorf("can't copy because src file is not a regular file (path=%s)", src.Path)
	}

	fSrc, err := os.Open(src.Path)
	if err != nil {
		return fmt.Errorf("can't open src file (path=%s): %v", src.Path, err)
	}
	defer fSrc.Close()

	if err := os.MkdirAll(dest.RelativeFolder(), 0755); err != nil {
		return fmt.Errorf("can't create dest folder (folder=%s path=%s) : %v", dest.RelativeFolder(), dest.Path, err)
	}

	if _, err := os.Stat(dest.Path); err == nil {
		if err := os.Remove(dest.Path); err != nil {
			return fmt.Errorf("can't delete existing dest file (path=%s): %v", dest.Path, err)
		}
	}

	fDest, err := os.OpenFile(dest.Path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, statSrc.Mode().Perm())
	if err != nil {
		return fmt.Errorf("can't open dest file (path=%s): %v", dest.Path, err)
	}
	defer fDest.Close()

	if _, err := io.Copy(fDest, fSrc); err != nil {
		return fmt.Errorf("can't copy src file to dest (src=%s, dest=%s): %v", src.Path, dest.Path, err)
	}

	return nil
}
