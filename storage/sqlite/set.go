package sqlite

import (
	log "github.com/sirupsen/logrus"

	"github.com/raystlin/bashful/storage"
)

const (
	queryUpsertCommand = `
INSERT INTO command(name, status)
VALUES (?,?)
ON CONFLICT(name) DO UPDATE SET
	status=EXCLUDED.status
`
	queryInsertStatus = `
INSERT INTO status(name, command, script) 
VALUES(?,?,?)
`
)

func (s *SQLiteStore) SetCommand(cmd *storage.Command, fullUpdate bool) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(queryUpsertCommand, cmd.Name, cmd.Status)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"stmt": queryUpsertCommand,
		}).Error("Could not update command")

		return err
	}

	if !fullUpdate {
		tx.Commit()
		return nil
	}

	stmt, err := tx.Prepare(queryInsertStatus)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"stmt": queryInsertStatus,
		}).Error("Could not prepare statement")

		return err
	}

	_, err = tx.Exec(queryDeleteStatus, cmd.Name)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"stmt": queryDeleteStatus,
		}).Error("Could not execute status deletion")

		return err
	}

	for _, v := range cmd.KnownStatus {
		_, err := stmt.Exec(v.Name, cmd.Name, v.Script)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"stmt": queryInsertStatus,
			}).Error("Could not execute status insert")

			return err
		}
	}

	tx.Commit()
	return nil
}
