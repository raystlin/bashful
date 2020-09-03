package bolt

import (
	bolt "go.etcd.io/bbolt"

	"github.com/raystlin/bashful/storage"
)

func (s *BoltStore) DeleteCommand(cmd *storage.Command) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(commandBucket))
		return b.Delete([]byte(cmd.Name))
	})
}
