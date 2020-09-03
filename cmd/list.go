package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List all commands and their current status",
	Action:  listAction,
	Flags: []cli.Flag{
		dbFlag,
	},
}

func listAction(c *cli.Context) error {
	db, err := NewStore(c)
	if err != nil {
		return err
	}
	defer db.Close()

	list, err := db.ListCommands()
	if err != nil {
		return err
	}

	fmt.Println("Registered Commands:")
	for i := range list {
		fmt.Printf("\t%s: %s\n", list[i].Name, list[i].Status)
	}

	return nil
}
