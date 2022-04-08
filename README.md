# Dotup

This command line is in charge of maintaining a local copy of you dot files from a `<source>` folder.

It copies all the files in the `<source>` folder to the `$HOME` folder, prefixing them by `.` if necessary.
When the `<source>` folder contains sub-folders, it will prefix the folders with a `.` but won't prefix the files and folders it contains.

## Usage

```shell
dotup  [flags] <source>

This command line is in charge of maintaining a local copy of you dot files from
a `<source>` folder.

It copies all the files in the `<source>` folder to the `$HOME` folder, prefixing
them by `.` if necessary. When the `<source>` folder contains sub-folders, it
will prefix the folders with a `.` but won't prefix the files and folders it
contains.

Flags

  -home string
    	if sets, overrides the default $HOME directory
  -mode string
    	define the kind of copy (symlink / cp) used for each dot file (default "symlink")
  -v	displays current program version
  -verbose
    	prints out all the operations and their status

Arguments

  <source>
    	is the path to the dotfiles folder (required)

Examples

  $ dotup ~/Workspace/dotfiles
    	symlinks all the files from ~/Workspace/dotfiles to $HOME

  $ dotup -mode cp ~/Workspace/dotfiles
    	cp all the files from ~/Workspace/dotfiles to $HOME

  $ dotup -home /tmp/test ~/Workspace/dotfiles
    	symlinks all the files from ~/Workspace/dotfiles to /tmp/test
```

## Installation

You can download the latest version from the [GitHub release](https://github.com/lonepeon/dotup/releases) page.
