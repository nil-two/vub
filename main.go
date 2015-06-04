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
  -v, --verbose             display the process
  -h, --help                show this help message
`[1:])
}

func _main() (code int, err error) {
	var filetype string
	flag.StringVar(&filetype, "f", "", "")
	flag.StringVar(&filetype, "filetype", "", "")

	var removeMode, updateMode bool
	flag.BoolVar(&removeMode, "r", false, "")
	flag.BoolVar(&removeMode, "remove", false, "")
	flag.BoolVar(&updateMode, "u", false, "")
	flag.BoolVar(&updateMode, "update", false, "")

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "")
	flag.BoolVar(&verbose, "verbose", false, "")

	var isHelp bool
	flag.BoolVar(&isHelp, "h", false, "")
	flag.BoolVar(&isHelp, "help", false, "")
	flag.Usage = usage
	flag.Parse()
	switch {
	case isHelp:
		usage()
		return 0, nil
	case flag.NArg() < 1:
		return 2, nil
	case removeMode && updateMode:
		return 2, fmt.Errorf("cannot specify multiple mode")
	}
	uri := flag.Arg(0)

	p, err := NewPackage(uri, filetype)
	if err != nil {
		return 1, err
	}
	p.Verbose(verbose)

	switch {
	case removeMode:
		if err := p.Remove(os.Stdout); err != nil {
			return 1, err
		}
	case updateMode:
		if err := p.Update(os.Stdout); err != nil {
			return 1, err
		}
	default:
		if err := p.Install(os.Stdout); err != nil {
			return 1, err
		}
	}
	return 0, nil
}

func main() {
	code, err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "vub:", err)
	}
	if code == 2 {
		shortUsage()
	}
	os.Exit(code)
}
