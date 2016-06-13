package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	cmdName = "vub"
	version = "0.3.1"

	flagset    = flag.NewFlagSet(cmdName, flag.ContinueOnError)
	filetype   = flagset.String("filetype", "", "")
	listMode   = flagset.Bool("list", false, "")
	removeMode = flagset.Bool("remove", false, "")
	updateMode = flagset.Bool("update", false, "")
	isHelp     = flagset.Bool("help", false, "")
	isVersion  = flagset.Bool("version", false, "")
)

func init() {
	flagset.SetOutput(ioutil.Discard)
	flagset.StringVar(filetype, "f", "", "")
	flagset.BoolVar(listMode, "l", false, "")
	flagset.BoolVar(removeMode, "r", false, "")
	flagset.BoolVar(updateMode, "u", false, "")
	flagset.BoolVar(isHelp, "h", false, "")
}

func countTrue(bls ...bool) int {
	cnt := 0
	for _, b := range bls {
		if b {
			cnt++
		}
	}
	return cnt
}

func printShortUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %[1]s [OPTION]... URI...
Try '%[1]s --help' for more information.
`[1:], cmdName)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [OPTION]... URI...
Install Vim plugin to under the management of vim-unbundle.

URI:
  sunaku/vim-unbundle                    # short URI
  https://github.com/sunaku/vim-unbundle # full URI

Options:
  -f, --filetype=TYPE       installing under the ftbundle/TYPE
  -l, --list                change the behavior to list packages
  -r, --remove              change the behavior to remove
  -u, --update              change the behavior to update
      --help                show this help message and exit
      --version             output version information and exit
`[1:], cmdName)
}

func printVersion() {
	fmt.Fprintln(os.Stderr, version)
}

func printErr(err interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmdName, err)
}

func _main() int {
	if err := flagset.Parse(os.Args[1:]); err != nil {
		printErr(err)
		return 2
	}
	if *isHelp {
		printUsage()
		return 0
	}
	if *isVersion {
		printVersion()
		return 0
	}

	if !*listMode && flagset.NArg() < 1 {
		printShortUsage()
		return 2
	}
	if countTrue(*listMode, *removeMode, *updateMode) > 1 {
		printErr("cannot specify multiple mode")
		return 2
	}

	if *listMode {
		ListPackages(*filetype)
		return 0
	} else {
		var err error
		for _, uri := range flagset.Args() {
			p := NewPackage(uri, *filetype)
			switch {
			case *removeMode:
				err = p.Remove()
			case *updateMode:
				err = p.Update()
			default:
				err = p.Install()
			}
			if err != nil {
				printErr(err)
				return 1
			}
		}
		return 0
	}
}

func main() {
	e := _main()
	os.Exit(e)
}
