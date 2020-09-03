package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/raystlin/bashful/storage"
)

const (
	nameFlagName = "name"
)

var DeleteCmd = &cli.Command{
	Name:    "delete",
	Aliases: []string{"d"},
	Usage:   "Delete a command from the database",
	Action:  deleteAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     nameFlagName,
			Aliases:  []string{"n"},
			Required: true,
		},
		dbFlag,
	},
}

func deleteAction(c *cli.Context) error {

	db, err := NewStore(c)
	if err != nil {
		return err
	}
	defer db.Close()

	cmd := &storage.Command{
		Name: c.String(nameFlagName),
	}

	return db.DeleteCommand(cmd)
}
