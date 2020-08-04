package main

import (
	"os"

	"github.com/davveo/go-pay/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	cliApp *cli.App
)

func init() {
	cliApp = cli.NewApp()
	cliApp.Name = "goPay"
	cliApp.Version = "0.0.1"
}

func main() {
	cliApp.Commands = []cli.Command{
		{
			Name:  "runserver",
			Usage: "run server",
			Action: func(c *cli.Context) error {
				return cmd.Runserver()
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		logrus.Errorf("Server err: %v", err.Error())
	}
}
