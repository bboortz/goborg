package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/docopt/docopt-go"
	"os"
)

func cmdHelp(argv []string) (err error) {
	usage := `usage: goborg help <command> 
`
	spew.Dump(argv)
	args, _ := docopt.ParseDoc(usage)
	printArgs(args)
	cmd := args["<command>"].(string)
	helpArgs := []string{"-h"}

	switch cmd {
	case "server":
		return cmdServer(helpArgs)
	case "client":
		fmt.Println("******* callback client")
		return cmdClient(helpArgs)
	}

	fmt.Fprintln(os.Stderr, usage)
	return err
}
