package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	All  bool `short:"a" long:"all" description:"Show all files including hidden ones." `
	Args struct {
		Target string `default:"."`
	} `positional-args:"yes"`
}

func main() {
	args := getArgs()
	args, err := flags.Parse(&opts)
	var targetDir string

	if err != nil {
		panic(err)
	}

	// if no target directory is specified, specify current directory.
	if opts.Args.Target == "" {
		targetDir = "."
	} else {
		targetDir = opts.Args.Target
	}

	// if no command line options provided, just show files excluded hidden ones.
	if len(args) == 0 {
		interpret("default", targetDir)
	} else {
		for _, arg := range args {
			interpret(arg, targetDir)
		}
	}
}

// getArgs returns given args.
func getArgs() []string {
	args := flag.Args()

	return args
}

// interpret args as options.
func interpret(arg string, targetDir string) {
	if opts.All {
		doEachFiles(showAll, targetDir)
	} else {
		doEachFiles(show, targetDir)
	}
}

// doEachFiles repeats process to each files in targetDir.
func doEachFiles(f func(fileName string), targetDir string) {
	files, err := ioutil.ReadDir(targetDir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		f(file.Name())
	}
}

// show files except hidden files.
func show(fileName string) {
	// don't show hidden files
	if !strings.HasPrefix(fileName, ".") {
		fmt.Printf("%s ", fileName)
	}
}

// show all files including hidden files.
func showAll(fileName string) {
	fmt.Printf("%s ", fileName)
}
