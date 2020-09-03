package bolt

import (
	"context"

	bolt "go.etcd.io/bbolt"
)

const (
	commandBucket = "Command"
)

type BoltStore struct {
	ctx context.Context
	db  *bolt.DB
}

func New(path string) (*BoltStore, error) {
	return NewWithContext(context.Background(), path)
}

func NewWithContext(ctx context.Context, path string) (*BoltStore, error) {
	db, err := bolt.Open(path, 0644, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(commandBucket))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &BoltStore{
		ctx: ctx,
		db:  db,
	}, nil
}

func (s *BoltStore) Context() context.Context {
	return s.ctx
}

func (s *BoltStore) Close() error {
	return s.db.Close()
}
