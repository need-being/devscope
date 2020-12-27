package main

import (
	"errors"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

var connectCommand = &cli.Command{
	Name:      "connect",
	UsageText: "devscope connect [command options] <address>",
	Action:    runConnect,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Value:   cli.NewStringSlice("PS1=$ "),
		},
		&cli.StringSliceFlag{
			Name:    "shell",
			Aliases: []string{"s"},
			Value:   cli.NewStringSlice("/bin/sh", "-i"),
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

	shell := ctx.StringSlice("shell")
	if len(shell) == 0 {
		return errors.New("missing shell")
	}
	cmd := exec.Command(shell[0], shell[1:]...)
	cmd.Env = append(os.Environ(), ctx.StringSlice("env")...)
	cmd.Stdin = io.TeeReader(conn, os.Stdout)
	cmd.Stdout = io.MultiWriter(conn, os.Stdout)
	cmd.Stderr = io.MultiWriter(conn, os.Stderr)
	if err := cmd.Start(); err != nil {
		return err
	}
	log.Infof("Running %v", strings.Join(shell, " "))
	log.Info("Control transferred to Server")

	return cmd.Wait()
}
