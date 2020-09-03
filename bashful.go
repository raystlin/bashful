package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/raystlin/bashful/cmd"
)

const (
	DebugFlagName = "debug"
)

func main() {
	app := &cli.App{
		Name:  "bashful",
		Usage: "A restful server to remotely invoke bash scripts",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    DebugFlagName,
				Aliases: []string{"d"},
				Value:   false,
				Usage:   "Debug level for logs",
			},
			&cli.StringFlag{
				Name:  cmd.StoreFlagName,
				Value: cmd.StoreFlagDefault,
				Usage: "Store type to use: " + cmd.KnownStores(),
			},
		},
		Commands: []*cli.Command{
			cmd.InitCmd,
			cmd.AddCmd,
			cmd.DeleteCmd,
			cmd.SetCmd,
			cmd.ListCmd,
			cmd.ServerCmd,
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				log.SetLevel(log.DebugLevel)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Fatal("Fatal error during the execution")
	}

	log.Info("Done!")
}
