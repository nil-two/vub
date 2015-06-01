package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

var (
	home, errInit = homedir.Dir()
	dotvim        = filepath.Join(home, ".vim")
)

func TestMain(m *testing.M) {
	if errInit != nil {
		fmt.Fprintln(os.Stderr, "vub:", errInit)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

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
