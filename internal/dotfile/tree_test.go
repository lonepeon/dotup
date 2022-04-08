package dotfile_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/lonepeon/dotup/internal/dotfile"
	"github.com/lonepeon/dotup/internal/dotfile/dotfiletest"
)

func TestWalkFolder(t *testing.T) {
	srcWorkdir := "testdata/sources"
	tmpFolder, err := ioutil.TempDir("", "walk-*")
	if err != nil {
		t.Fatalf("can't create temporary folder: %v", err)
	}
	defer os.RemoveAll(tmpFolder)

	var operations []dotfiletest.CopyOperation
	if err := dotfile.CopyTree(srcWorkdir, tmpFolder, dotfiletest.NoOpCopy, dotfiletest.WatchFileCopy(&operations)); err != nil {
		t.Fatalf("didn't expect failures while copying tree (src=%s, dest=%s): %v", "testdata/sources", tmpFolder, err)

	}

	expectedSources := []dotfile.File{
		dotfile.NewSource(srcWorkdir, "config/nvim/init.vim"),
		dotfile.NewSource(srcWorkdir, "gitconfig"),
	}

	if len(expectedSources) != len(operations) {
		var expectedFilenames []string
		for _, src := range expectedSources {
			expectedFilenames = append(expectedFilenames, src.Path)
		}

		var actualFilenames []string
		for _, operation := range operations {
			expectedFilenames = append(expectedFilenames, operation.Source.Path)
		}

		t.Fatalf("didn't copy the right number of files. want: %d; got: %d\nwant:\n%s\ngot:\n%s", len(expectedSources), len(operations), strings.Join(expectedFilenames, ", "), strings.Join(actualFilenames, ", "))
	}

	for i, expectedSource := range expectedSources {
		if expectedSource.Path != operations[i].Source.Path {
			t.Errorf("unexpected source at index %d: want:%s; got: %s", i, expectedSource.Path, operations[i].Source.Path)
			continue
		}

		expectedDestination := dotfile.NewDestination(tmpFolder, expectedSource)
		if expectedDestination.Path != operations[i].Destination.Path {
			t.Errorf("unexpected destination at index %d: want:%s; got: %s", i, expectedDestination.Path, operations[i].Destination.Path)
			continue
		}
	}
}
