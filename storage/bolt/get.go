package bolt

import (
	"encoding/json"

	bolt "go.etcd.io/bbolt"
	errors "golang.org/x/xerrors"

	"github.com/raystlin/bashful/storage"
)

func (s *BoltStore) GetCommand(name string) (*storage.Command, error) {
	command := new(storage.Command)
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(commandBucket))
		v := b.Get([]byte(name))

		if v == nil {
			return errors.Errorf("Command not found: %s", name)
		}

		return json.Unmarshal(v, command)
	})

	return command, err
}
