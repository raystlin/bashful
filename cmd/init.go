package cmd

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/raystlin/bashful/storage"
)

var InitCmd = &cli.Command{
	Name:   "init",
	Usage:  "Create a sample command file",
	Action: initAction,
	Flags: []cli.Flag{
		commandFlag,
	},
}

func initAction(c *cli.Context) error {

	cmd := storage.Command{
		Name:   "example",
		Status: "off",
		KnownStatus: map[string]*storage.Status{
			"off": &storage.Status{
				Name:   "off",
				Script: "off.sh",
			},
			"on": &storage.Status{
				Name:   "on",
				Script: "on.sh",
			},
		},
	}

	file, err := os.Create(c.String(commandFlagName))
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"path": c.String(commandFlagName),
		}).Error("Could not create command file")
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(&cmd)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"path": c.String(commandFlagName),
		}).Error("Could not save sample config")
	}

	return err
}
