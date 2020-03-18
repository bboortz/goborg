package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func cmdClient(argv []string) (err error) {
	usage := `usage: goborg client [options]

options:
	-h, --help             show this help
	-t, --target <host>    connect to host
	-p, --port <port>      connect to host
`

	args, _ := docopt.ParseArgs(usage, argv, "")
	fmt.Println(args)
	return err
}
