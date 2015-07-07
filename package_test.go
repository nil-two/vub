package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

var sourceURITests = []struct {
	src string
	dst string
}{
	//Full URI
	{
		src: "https://github.com/sunaku/vim-unbundle",
		dst: "https://github.com/sunaku/vim-unbundle",
	},

	//Short GitHub URI
	{
		src: "Shougo/neobundle.vim",
		dst: "https://github.com/Shougo/neobundle.vim",
	},
	{
		src: "thinca/vim-quickrun",
		dst: "https://github.com/thinca/vim-quickrun",
	},

	//GitHub URI
	{
		src: "github.com/Shougo/neobundle.vim",
		dst: "https://github.com/Shougo/neobundle.vim",
	},
	{
		src: "github.com/thinca/vim-quickrun",
		dst: "https://github.com/thinca/vim-quickrun",
	},
}

func TestSourceURI(t *testing.T) {
	for _, test := range sourceURITests {
		expect := test.dst
		actual := ToSourceURI(test.src)
		if actual != expect {
			t.Errorf("%q: got %q, want %q",
				test.src, actual, expect)
		}
	}
}

var destinationPathTests = []struct {
	filetype string
	src      string
	dst      string
}{
	//No filetype
	{
		filetype: "",
		src:      "https://github.com/sunaku/vim-unbundle",
		dst:      filepath.Join(dotvim, "bundle", "vim-unbundle"),
	},
	{
		filetype: "",
		src:      "sunaku/vim-unbundle",
		dst:      filepath.Join(dotvim, "bundle", "vim-unbundle"),
	},

	//Filetype specified
	{
		filetype: "go",
		src:      "https://github.com/fatih/vim-go",
		dst:      filepath.Join(dotvim, "ftbundle", "go", "vim-go"),
	},
	{
		filetype: "perl",
		src:      "https://github.com/hotchpotch/perldoc-vim",
		dst:      filepath.Join(dotvim, "ftbundle", "perl", "perldoc-vim"),
	},
}

func TestDestinationPath(t *testing.T) {
	for _, test := range destinationPathTests {
		expect := test.dst
		actual := ToDestinationPath(test.src, test.filetype)
		if actual != expect {
			t.Errorf("(uri=%q, filetype=%q): got %q, want %q",
				test.filetype, test.src, actual, expect)
		}
	}
}

func TestPackage(t *testing.T) {
	src, filetype := "sunaku/vim-unbundle", ""

	expect := &Package{
		src: "https://github.com/sunaku/vim-unbundle",
		dst: filepath.Join(dotvim, "bundle", "vim-unbundle"),
	}
	actual := NewPackage(src, filetype)
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("(uri=%q, filetype=%q): got %q, want %q",
			filetype, src, actual, expect)
	}
}

func TestPackageToCommnad(t *testing.T) {
	src, filetype := "sunaku/vim-unbundle", ""
	p := NewPackage(src, filetype)

	expect := exec.Command("git", "clone",
		"https://github.com/sunaku/vim-unbundle",
		filepath.Join(dotvim, "bundle", "vim-unbundle"))
	actual := p.toInstallCommand()
	if !reflect.DeepEqual(actual.Args, expect.Args) {
		t.Errorf("(filetype=%q, uri=%q): got %q, want %q",
			filetype, src, actual, expect)
	}
}

func TestInstalled(t *testing.T) {
	src, filetype := "sunaku/vim-unbundle", ""
	p := NewPackage(src, filetype)

	_, err := os.Stat(p.dst)
	expect := err == nil
	actual := p.installed()
	if actual != expect {
		t.Errorf("got %v, want %v",
			actual, expect)
	}
}
