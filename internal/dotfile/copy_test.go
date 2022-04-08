package dotfile_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/lonepeon/dotup/internal/dotfile"
	"github.com/lonepeon/dotup/internal/dotfile/dotfiletest"
)

func TestDuplicateFileWhenSourceExists(t *testing.T) {
	tmpFolder, err := ioutil.TempDir("", "duplicate-*")
	if err != nil {
		t.Fatalf("can't create temporary folder: %v", err)
	}
	defer os.RemoveAll(tmpFolder)

	tcs := map[string]struct {
		SrcFilepath          string
		ExpectedDestFilepath string
	}{
		"when file is at the root": {
			SrcFilepath:          "gitconfig",
			ExpectedDestFilepath: ".gitconfig",
		},
		"when file is in a sub folder": {
			SrcFilepath:          "config/nvim/init.vim",
			ExpectedDestFilepath: ".config/nvim/init.vim",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			src := dotfile.NewSource("testdata/sources", tc.SrcFilepath)
			dest := dotfile.NewDestination(tmpFolder, src)

			if dest.Filepath != tc.ExpectedDestFilepath {
				t.Fatalf("didn't expect this dest filepath: want: %s; got: %s", tc.ExpectedDestFilepath, dest.Filepath)
			}

			if err := dotfile.DuplicateFile(src, dest); err != nil {
				t.Fatalf("didn't expect an error on copy: %v", err)
			}

			if _, err := os.Stat(path.Join(tmpFolder, tc.ExpectedDestFilepath)); os.IsNotExist(err) {
				t.Fatalf("didn't expect destination file to not be present (dest=%s): %v", tc.ExpectedDestFilepath, err)
			}

			dotfiletest.AssertFileContentEqual(t, src, dest)
		})
	}
}

func TestSymlinkFileWhenSourceExists(t *testing.T) {
	tmpFolder, err := ioutil.TempDir("", "symlink-*")
	if err != nil {
		t.Fatalf("can't create temporary folder: %v", err)
	}
	defer os.RemoveAll(tmpFolder)

	tcs := map[string]struct {
		SrcFilepath          string
		ExpectedDestFilepath string
	}{
		"when file is at the root": {
			SrcFilepath:          "gitconfig",
			ExpectedDestFilepath: ".gitconfig",
		},
		"when file is in a sub folder": {
			SrcFilepath:          "config/nvim/init.vim",
			ExpectedDestFilepath: ".config/nvim/init.vim",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			src := dotfile.NewSource("testdata/sources", tc.SrcFilepath)
			dest := dotfile.NewDestination(tmpFolder, src)

			if dest.Filepath != tc.ExpectedDestFilepath {
				t.Fatalf("didn't expect this dest filepath: want: %s; got: %s", tc.ExpectedDestFilepath, dest.Filepath)
			}

			if err := dotfile.SymlinkFile(src, dest); err != nil {
				t.Fatalf("didn't expect an error on copy: %v", err)
			}

			if _, err := os.Lstat(dest.Path); os.IsNotExist(err) {
				t.Fatalf("didn't expect destination file to not be present (dest=%s): %v", dest.Path, err)
			}

			target, err := os.Readlink(dest.Path)
			if err != nil {
				t.Fatalf("didn't expect failure while reading dest symlink (dest=%s): %v", dest.Path, err)
			}

			absSrcPath, err := src.AbsolutePath()
			if err != nil {
				t.Fatalf("didn't expect to fail while building source absolute path (path=%s): %v", src.Path, err)
			}

			if target != absSrcPath {
				t.Fatalf("didn't expect difference between dest symlink target and src. want:%s; got:%s)", absSrcPath, target)
			}

			dotfiletest.AssertFileContentEqual(t, src, dest)
		})
	}
}
