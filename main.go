package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/mitchellh/go-homedir"
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
`[1:])
}

var (
	ShortGitHubURI = regexp.MustCompile(`^[\w\-.]+/[\w\-.]+$`)
)

func ToSourceURI(uri string) (string, error) {
	switch {
	case ShortGitHubURI.MatchString(uri):
		return "https://github.com/" + uri, nil
	default:
		return uri, nil
	}
}

var (
	home, errInit = homedir.Dir()
	dotvim        = filepath.Join(home, ".vim")
)

func ToDestinationPath(uri, filetype string) (string, error) {
	name := filepath.Base(uri)
	if filetype == "" {
		return filepath.Join(dotvim, "bundle", name), nil
	}
	return filepath.Join(dotvim, "ftbundle", filetype, name), nil
}

type Package struct {
	src string
	dst string
}

func NewPackage(uri, filetype string) (*Package, error) {
	src, err := ToSourceURI(uri)
	if err != nil {
		return nil, err
	}
	dst, err := ToDestinationPath(uri, filetype)
	if err != nil {
		return nil, err
	}
	return &Package{
		src: src,
		dst: dst,
	}, nil
}

func (p *Package) ToCommand() *exec.Cmd {
	return exec.Command("git", "clone", p.src, p.dst)
}

func main() {
	var filetype string
	flag.StringVar(&filetype, "f", "", "")
	flag.StringVar(&filetype, "filetype", "", "")

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
	}
	uri := flag.Arg(0)

	p, err := NewPackage(uri, filetype)
	if err != nil {
		fmt.Fprintln(os.Stderr, "vub:", err)
		os.Exit(2)
	}

	c := p.ToCommand()
	if err := c.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "vub:", err)
		os.Exit(1)
	}
}
