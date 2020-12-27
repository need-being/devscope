package main

import (
	"errors"
	"io"
	"net"
	"os"
	"os/exec"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

var connectCommand = &cli.Command{
	Name:      "connect",
	UsageText: "devscope connect [command options] <address>",
	Action:    runConnect,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "shell",
			Aliases: []string{"s"},
			Value:   "/bin/sh",
		},
	},
}

func runConnect(ctx *cli.Context) error {
	address := ctx.Args().First()
	if address == "" {
		return errors.New("missing argument: address")
	}

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Infof("Connected to %v", address)

	cmd := exec.Command(ctx.String("shell"))
	cmd.Stdin = io.TeeReader(conn, os.Stdout)
	cmd.Stdout = io.MultiWriter(conn, os.Stdout)
	cmd.Stderr = io.MultiWriter(conn, os.Stderr)
	return cmd.Run()
}
