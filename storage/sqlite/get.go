package sqlite

import (
	"github.com/raystlin/bashful/storage"
)

const (
	getQuery       = "SELECT * FROM command WHERE name = ?"
	getStatusQuery = "SELECT name, script FROM status WHERE command = ?"
)

func (s *SQLiteStore) GetCommand(name string) (*storage.Command, error) {
	command := storage.Command{}
	err := s.db.Get(&command, getQuery, name)
	if err != nil {
		return nil, err
	}

	status := []storage.Status{}
	err = s.db.Select(&status, getStatusQuery, name)
	if err != nil {
		return nil, err
	}

	command.KnownStatus = make(map[string]*storage.Status)
	for i := range status {
		command.KnownStatus[status[i].Name] = &status[i]
	}

	return &command, nil
}
