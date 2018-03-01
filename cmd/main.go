package main

import (
	"fmt"
	"os"

	"github.com/ianschenck/envflag"
	"github.com/nowenl/kecd/cmd/server"
	"github.com/urfave/cli"
)

func main() {
	envflag.Parse()

	app := cli.NewApp()
	app.Name = "kecd"
	app.Commands = []cli.Command{
		server.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
