package main

import (
	"flag"
	"fmt"
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
	var filetype string
	flag.StringVar(&filetype, "f", "", "")
	flag.StringVar(&filetype, "filetype", "", "")

	var listMode, removeMode, updateMode bool
	flag.BoolVar(&listMode, "l", false, "")
	flag.BoolVar(&listMode, "list", false, "")
	flag.BoolVar(&removeMode, "r", false, "")
	flag.BoolVar(&removeMode, "remove", false, "")
	flag.BoolVar(&updateMode, "u", false, "")
	flag.BoolVar(&updateMode, "update", false, "")

	var isHelp, isVersion bool
	flag.BoolVar(&isHelp, "help", false, "")
	flag.BoolVar(&isVersion, "version", false, "")
	flag.Usage = usage
	flag.Parse()
	switch {
	case isHelp:
		usage()
		os.Exit(0)
	case isVersion:
		version()
		os.Exit(0)
	case !listMode && flag.NArg() < 1:
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
		for _, uri := range flag.Args() {
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
