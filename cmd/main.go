package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	docopt "github.com/docopt/docopt-go"
	"os"
)

const PROG_NAME = "goborg"
const PROG_VERSION = "0.0.1"
const DEBUG = false

func printArgs(args docopt.Opts) (err error) {
	if !DEBUG {
		return err
	}

	if args == nil {
		fmt.Println("\n*** Unparse Program Arguments ***")
		spew.Dump(os.Args)
	} else {
		fmt.Println("\n*** Parsed Program Arguments ***")
		spew.Dump(args)
	}
	return err
}

func setupLogging() {
}

func run() (err error) {
	// define the program arguments/help
	usage := `usage: goborg [--version] [-h|--help] <command> [<args>...]
Options:
  -h --help     Show this screen.
  --version     Show version.

The commands are:
   client     starts goborg in client mode and connects to a server.
   server     starts goborg in server moe and wait for clients connecting.

Use'go help <command>' for more information on a specific command.
`

	// parsing the program arguments
	printArgs(nil)
	parser := &docopt.Parser{OptionsFirst: true}
	args, _ := parser.ParseArgs(usage, nil, PROG_NAME+" version "+PROG_VERSION)
	printArgs(args)
	cmd := args["<command>"].(string)
	cmdArgs := args["<args>"].([]string)

	// setup the new cmd for the sub-command
	newCmd := []string{cmd}
	newCmd = append(newCmd, cmdArgs...)

	switch cmd {
	case "server":
		return cmdServer(newCmd)
	case "client":
		return cmdClient(newCmd)
	case "help", "":
		return cmdHelp(newCmd)
	}

	err = fmt.Errorf("<%s> is not a git command. See 'goborg --help'", cmd)
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
