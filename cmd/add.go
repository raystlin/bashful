package cmd

import (
	"encoding/json"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/raystlin/bashful/storage"
)

var AddCmd = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Adds a command and its statuses to the database. If you need a template call init subcommand",
	Action:  addAction,
	Flags: []cli.Flag{
		commandFlag,
		dbFlag,
	},
}

func addAction(c *cli.Context) error {
	db, err := NewStore(c)
	if err != nil {
		return err
	}
	defer db.Close()

	file, err := os.Open(c.String(commandFlagName))
	if err != nil {
		return err
	}
	defer file.Close()

	cmd := new(storage.Command)

	err = json.NewDecoder(file).Decode(cmd)
	if err != nil {
		return err
	}

	return db.SetCommand(cmd, true)
}
