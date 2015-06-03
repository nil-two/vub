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
		"https://github.com/sunaku/vim-unbundle",
		"https://github.com/sunaku/vim-unbundle",
	},

	//Short GitHub URI
	{
		"Shougo/neobundle.vim",
		"https://github.com/Shougo/neobundle.vim",
	},
	{
		"thinca/vim-quickrun",
		"https://github.com/thinca/vim-quickrun",
	},
}

func TestSourceURI(t *testing.T) {
	for _, test := range sourceURITests {
		expect := test.dst
		actual, err := ToSourceURI(test.src)
		if err != nil {
			t.Errorf("ToSourceURI(%q) returns %q, want nil",
				test.src, err)
		}
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
		"",
		"https://github.com/sunaku/vim-unbundle",
		filepath.Join(dotvim, "bundle", "vim-unbundle"),
	},
	{
		"",
		"sunaku/vim-unbundle",
		filepath.Join(dotvim, "bundle", "vim-unbundle"),
	},

	//Filetype specified
	{
		"go",
		"https://github.com/fatih/vim-go",
		filepath.Join(dotvim, "ftbundle", "go", "vim-go"),
	},
	{
		"perl",
		"https://github.com/hotchpotch/perldoc-vim",
		filepath.Join(dotvim, "ftbundle", "perl", "perldoc-vim"),
	},
}

func TestDestinationPath(t *testing.T) {
	for _, test := range destinationPathTests {
		expect := test.dst
		actual, err := ToDestinationPath(test.src, test.filetype)
		if err != nil {
			t.Errorf("ToSourceURI(%q) returns %q, want nil",
				test.src, err)
		}
		if actual != expect {
			t.Errorf("(uri=%q, filetype=%q): got %q, want %q",
				test.filetype, test.src, actual, expect)
		}
	}
}

func TestPackage(t *testing.T) {
	src, filetype := "sunaku/vim-unbundle", ""
	expect := &Package{
		verbose: false,
		src:     "https://github.com/sunaku/vim-unbundle",
		dst:     filepath.Join(dotvim, "bundle", "vim-unbundle"),
	}
	actual, err := NewPackage(src, filetype)
	if err != nil {
		t.Errorf("NewPackage(%q, %q) returns %q, want nil",
			src, filetype, err)
	}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("(uri=%q, filetype=%q): got %q, want %q",
			filetype, src, actual, expect)
	}
}

func TestPackageToCommnad(t *testing.T) {
	src, filetype := "sunaku/vim-unbundle", ""
	p, err := NewPackage(src, filetype)
	if err != nil {
		t.Errorf("NewPackage(%q, %q) returns %q, want nil",
			src, filetype, err)
	}
	expect := exec.Command("git", "clone",
		"https://github.com/sunaku/vim-unbundle",
		filepath.Join(dotvim, "bundle", "vim-unbundle"))
	actual := p.toCommand()
	if !reflect.DeepEqual(actual.Args, expect.Args) {
		t.Errorf("(filetype=%q, uri=%q): got %q, want %q",
			filetype, src, actual, expect)
	}
}

func TestVerbose(t *testing.T) {
	src, filetype := "sunaku/vim-unbundle", ""
	p, err := NewPackage(src, filetype)
	if err != nil {
		t.Errorf("NewPackage(%q, %q) returns %q, want nil",
			src, filetype, err)
	}
	p.Verbose(true)
	expect := true
	actual := p.verbose
	if actual != expect {
		t.Errorf("got %v, want %v",
			actual, expect)
	}
}

func TestInstalled(t *testing.T) {
	src, filetype := "sunaku/vim-unbundle", ""
	p, err := NewPackage(src, filetype)
	if err != nil {
		t.Errorf("NewPackage(%q, %q) returns %q, want nil",
			src, filetype, err)
	}
	_, err = os.Stat(p.dst)
	expect := err == nil
	actual := p.installed()
	if actual != expect {
		t.Errorf("got %v, want %v",
			actual, expect)
	}
}
