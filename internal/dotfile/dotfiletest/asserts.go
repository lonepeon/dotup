package dotfiletest

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/lonepeon/dotup/internal/dotfile"
)

// AssertFileContentEqual opens the two files and check if their content match
func AssertFileContentEqual(t *testing.T, src dotfile.File, dest dotfile.File) {
	t.Helper()

	srcPath, err := src.AbsolutePath()
	if err != nil {
		t.Fatalf("cannot generate src absolute path from '%s': %v", src.Path, err)
	}

	destPath, err := dest.AbsolutePath()
	if err != nil {
		t.Fatalf("cannot generate dest absolute path from '%s': %v", dest.Path, err)
	}

	fSrc, err := os.Open(srcPath)
	if err != nil {
		t.Fatalf("can't open src file '%s' for comparison: %v", src, err)
	}
	defer fSrc.Close()

	fDest, err := os.Open(destPath)
	if err != nil {
		t.Fatalf("can't open dest file '%s' for comparison: %v", dest, err)
	}
	defer fDest.Close()

	chunkSize := 4096
	for {
		bSrc := make([]byte, chunkSize)
		_, errSrc := fSrc.Read(bSrc)

		bDest := make([]byte, chunkSize)
		_, errDest := fDest.Read(bDest)

		if errSrc != nil || errDest != nil {
			if errSrc == io.EOF && errDest == io.EOF {
				return
			}

			if errSrc == io.EOF || errDest == io.EOF {
				t.Errorf("didn't expect files to have different sizes (src=%s, dest=%s)", src, dest)
			}

			if errSrc != nil {
				t.Errorf("didn't expect an error while reading source '%s' file content: %v", src, errSrc)
			}

			if errDest != nil {
				t.Errorf("didn't expect an error while reading dest '%s' file content: %v", src, errSrc)
			}

			return
		}

		if !bytes.Equal(bSrc, bDest) {
			t.Errorf("didn't expect files content to be different (src=%s, dest=%s)", src, dest)
			return
		}
	}
}
