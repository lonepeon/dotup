package fileutil_test

import (
	"testing"

	"github.com/lonepeon/dotup/internal/dotfile/fileutil"
)

func TestMakeHidden(t *testing.T) {
	tcs := map[string]struct {
		Input          string
		ExpectedOutput string
	}{
		"file with no dot and file at the root": {
			Input:          "gitconfig",
			ExpectedOutput: ".gitconfig",
		},
		"file with no dot and file at the root of a folder": {
			Input:          "./gitconfig",
			ExpectedOutput: ".gitconfig",
		},
		"file with no dot and file in a folder": {
			Input:          "config/nvim/init.nvim",
			ExpectedOutput: ".config/nvim/init.nvim",
		},
		"file with no dot and file in a relative folder": {
			Input:          "./config/nvim/init.nvim",
			ExpectedOutput: ".config/nvim/init.nvim",
		},
		"file with dot and at the root": {
			Input:          ".gitconfig",
			ExpectedOutput: ".gitconfig",
		},
		"file with dot and at the root of a folder": {
			Input:          "./.gitconfig",
			ExpectedOutput: ".gitconfig",
		},
		"file with dot and in a folder with no dot": {
			Input:          "config/nvim/.init.nvim",
			ExpectedOutput: ".config/nvim/.init.nvim",
		},
		"file with dot and in a folder with a dot": {
			Input:          ".config/nvim/.init.nvim",
			ExpectedOutput: ".config/nvim/.init.nvim",
		},
		"file with no dot and in a folder with no dot": {
			Input:          "config/nvim/init.nvim",
			ExpectedOutput: ".config/nvim/init.nvim",
		},
		"file with no dot and in a folder with a dot": {
			Input:          ".config/nvim/init.nvim",
			ExpectedOutput: ".config/nvim/init.nvim",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			output := fileutil.MakeHidden(tc.Input)
			if tc.ExpectedOutput != output {
				t.Fatalf("unepxected ouput. want: %s; got: %s", tc.ExpectedOutput, output)
			}
		})
	}
}
