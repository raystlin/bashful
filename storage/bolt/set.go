package bolt

import (
	"encoding/json"

	bolt "go.etcd.io/bbolt"
	errors "golang.org/x/xerrors"

	"github.com/raystlin/bashful/storage"
)

func (s *BoltStore) SetCommand(cmd *storage.Command, fullUpdate bool) error {

	return s.db.Update(func(tx *bolt.Tx) error {

		var fullCmd *storage.Command
		b := tx.Bucket([]byte(commandBucket))
		if !fullUpdate {
			v := b.Get([]byte(cmd.Name))
			if v == nil {
				return errors.Errorf("Command not found: %s", cmd.Name)
			}

			fullCmd = &storage.Command{}
			err := json.Unmarshal(v, fullCmd)
			if err != nil {
				return err
			}

			fullCmd.Status = cmd.Status
		} else {
			fullCmd = cmd
		}

		data, err := json.Marshal(fullCmd)
		if err != nil {
			return err
		}

		return b.Put([]byte(cmd.Name), data)
	})
}
