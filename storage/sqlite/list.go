package sqlite

import (
	"github.com/raystlin/bashful/storage"
)

const (
	listQuery = "SELECT * FROM command"
)

func (s *SQLiteStore) ListCommands() ([]*storage.Command, error) {
	commands := []*storage.Command{}
	return commands, s.db.Select(&commands, listQuery)
}
