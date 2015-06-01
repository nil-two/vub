package main

import (
	"path/filepath"
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
			t.Errorf("ToSourceURI(%q) returns %q, want nil", err)
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
	{
		"",
		"https://github.com/sunaku/vim-unbundle",
		filepath.Join(dotvim, "bundle", "vim-unbundle"),
	},
}

func TestDestinationPath(t *testing.T) {
	for _, test := range destinationPathTests {
		expect := test.dst
		actual, err := ToDestinationPath(test.src, test.filetype)
		if err != nil {
			t.Errorf("ToSourceURI(%q) returns %q, want nil", err)
		}
		if actual != expect {
			t.Errorf("(filetype=%q, uri=%q): got %q, want %q",
				test.filetype, test.src, actual, expect)
		}
	}
}
