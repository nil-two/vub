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
  -v, --verbose             display the process
  -h, --help                show this help message
`[1:])
}

func main() {
	filetype, removeMode, updateMode, verbose := "", false, false, false
	flag.StringVar(&filetype, "f", "", "")
	flag.StringVar(&filetype, "filetype", "", "")
	flag.BoolVar(&removeMode, "r", false, "")
	flag.BoolVar(&removeMode, "remove", false, "")
	flag.BoolVar(&updateMode, "u", false, "")
	flag.BoolVar(&updateMode, "update", false, "")
	flag.BoolVar(&verbose, "v", false, "")
	flag.BoolVar(&verbose, "verbose", false, "")

	isHelp := false
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
	}
	uri := flag.Arg(0)

	p, err := NewPackage(uri, filetype)
	if err != nil {
		fmt.Fprintln(os.Stderr, "vub:", err)
		os.Exit(2)
	}
	p.Verbose(verbose)

	switch {
	case removeMode:
		if err := p.Remove(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "vub:", err)
			os.Exit(1)
		}
	case updateMode:
		if err := p.Update(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "vub:", err)
			os.Exit(1)
		}
	default:
		if err := p.Install(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "vub:", err)
			os.Exit(1)
		}
	}
}
