package main

import (
	"github.com/bboortz/goborg/pkg/appcontext"
	"github.com/bboortz/goborg/pkg/server"
	"github.com/docopt/docopt-go"
	"strconv"
)

func cmdServer(argv []string) (err error) {
	usage := `usage: goborg server [options]

options:
	-h, --help             show this help
	-p, --port <port>      listen on port [default:5555]
`
	args, _ := docopt.ParseArgs(usage, argv, "")
	printArgs(args)

	// parse port
	port, argErr := args.Int("--port")
	if argErr != nil {
		port = 5555
	}

	// start the server
	addr := ":" + strconv.Itoa(port)
	var srv server.Server
	srv.ListenAndServe(appcontext.Ctx(), addr)

	return err
}
