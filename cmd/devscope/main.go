package main

import (
	"os"

	"github.com/apex/log"
	apexcli "github.com/apex/log/handlers/cli"
	"github.com/urfave/cli/v2"
)

func main() {
	log.SetHandler(apexcli.New(os.Stderr))
	app := &cli.App{
		Name: "devscope",
		Commands: []*cli.Command{
			connectCommand,
			listenCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
