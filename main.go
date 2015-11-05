package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func shortUsage() {
	os.Stderr.WriteString(`
Usage: vub [OPTION]... URI...
Try 'vub --help' for more information.
`[1:])
}

func usage() {
	os.Stderr.WriteString(`
Usage: vub [OPTION]... URI...
Install Vim plugin to under the management of vim-unbundle.

URI:
  sunaku/vim-unbundle                    # short URI
  https://github.com/sunaku/vim-unbundle # full URI

Options:
  -f, --filetype=TYPE       installing under the ftbundle/TYPE
  -l, --list                change the behavior to list packages
  -r, --remove              change the behavior to remove
  -u, --update              change the behavior to clean update
      --help                show this help message
      --version             output version information
`[1:])
}

func version() {
	os.Stderr.WriteString(`
v0.2.0
`[1:])
}

func countTrue(bls ...bool) int {
	cnt := 0
	for _, b := range bls {
		if b {
			cnt++
		}
	}
	return cnt
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, "vub:", err)
}

func main() {
	f := flag.NewFlagSet("vub", flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)

	var filetype string
	f.StringVar(&filetype, "f", "", "")
	f.StringVar(&filetype, "filetype", "", "")

	var listMode, removeMode, updateMode bool
	f.BoolVar(&listMode, "l", false, "")
	f.BoolVar(&listMode, "list", false, "")
	f.BoolVar(&removeMode, "r", false, "")
	f.BoolVar(&removeMode, "remove", false, "")
	f.BoolVar(&updateMode, "u", false, "")
	f.BoolVar(&updateMode, "update", false, "")

	var isHelp, isVersion bool
	f.BoolVar(&isHelp, "help", false, "")
	f.BoolVar(&isVersion, "version", false, "")

	if err := f.Parse(os.Args[1:]); err != nil {
		printError(err)
		os.Exit(2)
	}
	switch {
	case isHelp:
		usage()
		os.Exit(0)
	case isVersion:
		version()
		os.Exit(0)
	case !listMode && f.NArg() < 1:
		shortUsage()
		os.Exit(2)
	case countTrue(listMode, removeMode, updateMode) > 1:
		printError(fmt.Errorf("cannot specify multiple mode"))
		os.Exit(2)
	}

	switch {
	case listMode:
		ListPackages(filetype)
	default:
		var err error
		for _, uri := range f.Args() {
			p := NewPackage(uri, filetype)
			switch {
			case removeMode:
				err = p.Remove()
			case updateMode:
				err = p.Update()
			default:
				err = p.Install()
			}
			if err != nil {
				printError(err)
				os.Exit(1)
			}
		}
	}
}
