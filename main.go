package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/lonepeon/dotup/internal/build"
	"github.com/lonepeon/dotup/internal/dotfile"
	"github.com/lonepeon/dotup/internal/env"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	var cfg struct {
		Mode        string
		Homedir     string
		Verbose     bool
		ShowVersion bool
	}

	fset := flag.NewFlagSet("dotup", flag.ExitOnError)
	fset.Usage = func() {
		var buf strings.Builder
		fset.SetOutput(&buf)
		fset.PrintDefaults()

		fset.SetOutput(os.Stderr)
		fmt.Fprintf(fset.Output(), summary, fset.Name(), buf.String())
	}

	fset.StringVar(&cfg.Mode, "mode", "symlink", "define the kind of copy (symlink / cp) used for each dot file")
	fset.StringVar(&cfg.Homedir, "home", "", "if sets, overrides the default $HOME directory")
	fset.BoolVar(&cfg.Verbose, "verbose", false, "prints out all the operations and their status")
	fset.BoolVar(&cfg.ShowVersion, "v", false, "displays current program version")
	if err := fset.Parse(os.Args[1:]); err != nil {
		return err
	}

	if cfg.ShowVersion {
		fmt.Println(os.Args[0], build.GetVersionInfo().String())
		return nil
	}

	operation, ok := map[string]dotfile.CopyFunc{
		"symlink": dotfile.SymlinkFile,
		"cp":      dotfile.DuplicateFile,
	}[cfg.Mode]

	if !ok {
		return fmt.Errorf("invalid operation mode. must be one of: symlink, cp")
	}

	watcher := dotfile.NoopWatcher
	if cfg.Verbose {
		output := tabwriter.NewWriter(os.Stdout, 2, 1, 1, ' ', tabwriter.TabIndent)
		defer output.Flush()

		fmt.Println("output mode:", cfg.Mode)
		fmt.Fprintln(output, "src\tdest\tstatus")
		watcher = func(src dotfile.File, dest dotfile.File, err error) {
			status := "OK"
			if err != nil {
				status = err.Error()
			}

			fmt.Fprintf(output, "%s\t%s\t%s\n", src.Path, dest.Path, status)
		}
	}

	if flag.NArg() != 1 {
		flag.Usage()
		return fmt.Errorf("invalid number of arguments. should have exactly one")
	}

	srcWorkdir := flag.Arg(0)
	destWorkdir := cfg.Homedir
	if destWorkdir == "" {
		homedir, err := env.Homedir()
		if err != nil {
			return fmt.Errorf("can't detect user's home directory. the value can be overriden using the -home flag")
		}

		destWorkdir = homedir
	}

	if err := dotfile.CopyTree(srcWorkdir, destWorkdir, operation, watcher); err != nil {
		return fmt.Errorf("can't %s dotfiles from '%s' to '%s': %v", cfg.Mode, srcWorkdir, destWorkdir, err)
	}

	return nil
}

const summary = `%[1]s [flags] <source>

This command line is in charge of maintaining a local copy of you dot files from
a <source> folder.

It copies all the files in the <source> folder to the $HOME folder, prefixing
them by "." if necessary. When the <source> folder contains sub-folders, it
will prefix the folders with a "." but won't prefix the files and folders it
contains.

Flags
%[2]s
Arguments
  <source>
    \tis the path to the dotfiles folder (required)

Examples
  $> %[1]s ~/Workspace/dotfiles
    \tsymlinks all the files from ~/Workspace/dotfiles to $HOME

  $> %[1]s -mode cp ~/Workspace/dotfiles
    \tcp all the files from ~/Workspace/dotfiles to $HOME

  $> %[1]s -home /tmp/test ~/Workspace/dotfiles
    \tsymlinks all the files from ~/Workspace/dotfiles to /tmp/test
`
