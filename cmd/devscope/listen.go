package main

import (
	"errors"
	"io"
	"net"
	"os"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

var listenCommand = &cli.Command{
	Name:      "listen",
	UsageText: "devscope listen [command options] <address>",
	Action:    runListen,
}

func runListen(ctx *cli.Context) error {
	address := ctx.Args().First()
	if address == "" {
		return errors.New("missing argument: address")
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()
	log.Infof("Listening at %s", address)

	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Infof("Accepted from %v", conn.RemoteAddr())

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			log.Errorf("Fail to send: %v", err)
		}
	}()
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		return err
	}
	log.Info("Connection closed by remote host")
	return nil
}
