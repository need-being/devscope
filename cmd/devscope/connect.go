package main

import (
	"github.com/urfave/cli/v2"
)

var connectCommand = &cli.Command{
	Name:   "connect",
	Action: runConnect,
}

func runConnect(ctx *cli.Context) error {
	return nil
}
