package main

import (
	"flag"
	"fmt"
	"os"
)

func shortUsage() {
	os.Stderr.WriteString(`
Usage: vub [OPTION]... URI
Try 'vub --help' for more information.
`[1:])
}

func usage() {
	os.Stderr.WriteString(`
Usage: vub [OPTION]... URI
Install Vim plugin to under the management of vim-unbundle.

URI:
  sunaku/vim-unbundle                    # short URI
  https://github.com/sunaku/vim-unbundle # full URI

Options:
  -f, --filetype=TYPE       installing under the ftbundle/TYPE
  -r, --remove              change the behavior to remove
  -u, --update              change the behavior to clean update
  -h, --help                show this help message
`[1:])
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, "vub:", err)
}

func main() {
	var filetype string
	flag.StringVar(&filetype, "f", "", "")
	flag.StringVar(&filetype, "filetype", "", "")

	var removeMode, updateMode bool
	flag.BoolVar(&removeMode, "r", false, "")
	flag.BoolVar(&removeMode, "remove", false, "")
	flag.BoolVar(&updateMode, "u", false, "")
	flag.BoolVar(&updateMode, "update", false, "")

	var isHelp bool
	flag.BoolVar(&isHelp, "h", false, "")
	flag.BoolVar(&isHelp, "help", false, "")
	flag.Usage = usage
	flag.Parse()
	switch {
	case isHelp:
		usage()
		os.Exit(0)
	case flag.NArg() < 1:
		shortUsage()
		os.Exit(2)
	case removeMode && updateMode:
		printError(fmt.Errorf("cannot specify multiple mode"))
		os.Exit(2)
	}
	uri := flag.Arg(0)

	p := NewPackage(uri, filetype)

	var err error
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
