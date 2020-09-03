package sqlite

import (
	"github.com/raystlin/bashful/storage"
)

const (
	queryDeleteStatus  = "DELETE FROM status WHERE command = ?"
	queryDeleteCommand = "DELETE FROM command WHERE name = ?"
)

func (s *SQLiteStore) DeleteCommand(cmd *storage.Command) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(queryDeleteStatus)
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryDeleteCommand)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
