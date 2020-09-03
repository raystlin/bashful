package storage

import (
	"context"
	"os/exec"
	"time"

	log "github.com/sirupsen/logrus"
)

type Command struct {
	Name        string             `db:"name" json:"name"`
	Status      string             `db:"status" json:"status"`
	KnownStatus map[string]*Status `json:"known_status"`
}

type Status struct {
	Name    string   `db:"name" json:"name"`
	Script  string   `db:"script" json:"script"`
	Command *Command `json:"-"`
}

type Store interface {
	Context() context.Context
	GetCommand(name string) (*Command, error)
	SetCommand(cmd *Command, fullUpdate bool) error
	ListCommands() ([]*Command, error)
	DeleteCommand(cmd *Command) error
	Close() error
}

/// Sets a command to a new status and triggers its attached script.
/// ctx Execution context
/// s Store with the commands
/// name The name of the command to change
/// newStatus The name of the new status to change.
func SetCommandStatus(s Store, name, newStatus string) error {
	cmd, err := s.GetCommand(name)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"name": name,
		}).Error("Error retrieving command")

		return err
	}

	if cmd.Status == newStatus {
		log.WithFields(log.Fields{
			"name":   name,
			"status": newStatus,
		}).Warn("Required to switch to the same status")

		return err
	}

	status, ok := cmd.KnownStatus[newStatus]
	if !ok {
		log.WithFields(log.Fields{
			"name":   name,
			"status": newStatus,
		}).Error("Unknown new status")

		return err
	}

	timedCtx, cancel := context.WithTimeout(s.Context(), 5*time.Minute)
	defer cancel()

	command := exec.CommandContext(timedCtx, status.Script)
	err = command.Run()
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"name":    name,
			"status":  newStatus,
			"command": status.Command,
		}).Error("Error executing status command")

		return err
	}

	cmd.Status = newStatus
	err = s.SetCommand(cmd, false)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"name":  name,
			"staus": newStatus,
		}).Error("Could not set the new status")

		return err
	}

	return nil
}
