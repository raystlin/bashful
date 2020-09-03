package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/raystlin/bashful/storage"
)

const (
	statusFlagName      = "status"
	commandNameFlagName = "command"
)

var SetCmd = &cli.Command{
	Name:   "set",
	Usage:  "Sets the current status of a command, triggering its action",
	Action: setAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     statusFlagName,
			Aliases:  []string{"s"},
			Usage:    "New status to set",
			Required: true,
		},
		&cli.StringFlag{
			Name:     commandNameFlagName,
			Aliases:  []string{"c"},
			Usage:    "Command to change status",
			Required: true,
		},
		dbFlag,
	},
}

func setAction(c *cli.Context) error {
	db, err := NewStore(c)
	if err != nil {
		return err
	}
	defer db.Close()

	return storage.SetCommandStatus(db,
		c.String(commandNameFlagName),
		c.String(statusFlagName))
}
