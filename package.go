package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mitchellh/go-homedir"
)

var (
	ShortGitHubURI = regexp.MustCompile(`^[\w\-.]+/[\w\-.]+$`)
	GitHubURI      = regexp.MustCompile(`^github.com/[\w\-.]+/[\w\-.]+$`)
	BitbucketURI   = regexp.MustCompile(`^bitbucket.org/[\w\-.]+/[\w\-.]+$`)
)

func ToSourceURI(uri string) string {
	switch {
	case ShortGitHubURI.MatchString(uri):
		return "https://github.com/" + uri
	case GitHubURI.MatchString(uri):
		return "https://" + uri
	case BitbucketURI.MatchString(uri):
		return "https://" + uri
	default:
		return uri
	}
}

var (
	home, errInit = homedir.Dir()
	dotvim        = filepath.Join(home, ".vim")
)

func ToDestinationPath(uri, filetype string) string {
	name := filepath.Base(uri)
	if filetype == "" {
		return filepath.Join(dotvim, "bundle", name)
	}
	return filepath.Join(dotvim, "ftbundle", filetype, name)
}

func ListPackages(filetype string) error {
	var path string
	if filetype == "" {
		path = filepath.Join(dotvim, "bundle")
	} else {
		path = filepath.Join(dotvim, "ftbundle", filetype)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

type Package struct {
	src string
	dst string
}

func NewPackage(uri, filetype string) *Package {
	return &Package{
		src: ToSourceURI(uri),
		dst: ToDestinationPath(uri, filetype),
	}
}

func (p *Package) toInstallCommand() *exec.Cmd {
	return exec.Command("git", "clone", p.src, p.dst)
}

func (p *Package) installed() bool {
	_, err := os.Stat(p.dst)
	return err == nil
}

func (p *Package) Install() error {
	if p.installed() {
		return nil
	}
	if _, err := exec.LookPath("git"); err != nil {
		return err
	}

	errMessage := bytes.NewBuffer(make([]byte, 0))

	installcmd := p.toInstallCommand()
	installcmd.Stderr = errMessage
	if err := installcmd.Run(); err != nil {
		return fmt.Errorf("%s", strings.TrimSpace(errMessage.String()))
	}
	return nil
}

func (p *Package) Remove() error {
	if !p.installed() {
		return nil
	}
	return os.RemoveAll(p.dst)
}

func (p *Package) Update() error {
	if p.installed() {
		if err := p.Remove(); err != nil {
			return err
		}
	}
	return p.Install()
}
