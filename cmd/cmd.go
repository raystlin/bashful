package cmd

import (
	"strings"

	"github.com/urfave/cli/v2"
	errors "golang.org/x/xerrors"

	"github.com/raystlin/bashful/storage"
	"github.com/raystlin/bashful/storage/bolt"
	"github.com/raystlin/bashful/storage/sqlite"
)

const (
	dbFlagName    = "database"
	dbFlagDefault = "data.db"

	commandFlagName    = "command"
	commandFlagDefault = "command.json"
)

const (
	StoreSQLite = "sqlite3"
	StoreBolt   = "bolt"

	StoreFlagName    = "store"
	StoreFlagDefault = StoreBolt
)

var knownStores = []string{
	StoreSQLite,
	StoreBolt,
}

func KnownStores() string {
	return strings.Join(knownStores, ", ")
}

var (
	dbFlag = &cli.StringFlag{
		Name:    dbFlagName,
		Aliases: []string{"d"},
		Value:   dbFlagDefault,
		Usage:   "Database with the command definitions",
	}
	commandFlag = &cli.StringFlag{
		Name:  commandFlagName,
		Value: commandFlagDefault,
		Usage: "Command file",
	}
)

func NewStore(c *cli.Context) (storage.Store, error) {
	switch c.String(StoreFlagName) {
	case StoreSQLite:
		return sqlite.New(c.String(dbFlagName))
	case StoreBolt:
		return bolt.New(c.String(dbFlagName))
	default:
		return nil, errors.Errorf("Unknown store type '%s'", c.String(StoreFlagName))
	}
}
