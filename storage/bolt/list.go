package bolt

import (
	"encoding/json"

	bolt "go.etcd.io/bbolt"

	"github.com/raystlin/bashful/storage"
)

func (s *BoltStore) ListCommands() ([]*storage.Command, error) {
	commands := make([]*storage.Command, 0)
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(commandBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			cmd := new(storage.Command)

			err := json.Unmarshal(v, cmd)
			if err != nil {
				return err
			}

			commands = append(commands, cmd)
		}

		return nil
	})

	return commands, err
}
