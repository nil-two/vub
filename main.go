package main

import (
	"flag"
	"fmt"
	"os"
)

func shortUsage() {
	os.Stderr.WriteString(`
usage: vub [option(s)] <repository-uri>
try 'vub --help' for more information.
`[1:])
}

func usage() {
	os.Stderr.WriteString(`
usage: vub [option(s)] <repository-uri>
install Vim plugin to under the management of vim-unbundle.

options:
  -h, --help                show this help message
  -f, --filetype=TYPE       installing under the ftbundle/TYPE
  -v, --verbose             display the process
`[1:])
}

func main() {
	filetype, verbose := "", false
	flag.StringVar(&filetype, "f", "", "")
	flag.StringVar(&filetype, "filetype", "", "")

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

	if err := p.Install(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "vub:", err)
		os.Exit(1)
	}
}
