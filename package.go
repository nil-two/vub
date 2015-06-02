package main

import (
	"io"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"
)

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
	verbose bool
	src     string
	dst     string
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

func (p *Package) toCommand() *exec.Cmd {
	return exec.Command("git", "clone", p.src, p.dst)
}

func (p *Package) Verbose(enable bool) {
	p.verbose = enable
}

func (p *Package) Install(w io.Writer) error {
	c := p.toCommand()
	if w != nil && p.verbose {
		c.Stdout = w
		c.Stderr = w
		_, err := io.WriteString(w, strings.Join(c.Args, " "))
		if err != nil {
			return err
		}
	}
	return c.Run()
}
