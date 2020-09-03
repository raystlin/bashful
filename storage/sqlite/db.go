package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

const (
	createTableCommand = `
CREATE TABLE IF NOT EXISTS command (
	name TEXT NOT NULL PRIMARY KEY,
	status TEXT NOT NULL
)`

	createTableStatus = `
CREATE TABLE IF NOT EXISTS status (
	name TEXT NOT NULL,
	command TEXT NOT NULL,
	script TEXT NOT NULL,
	FOREIGN KEY(command) REFERENCES command(name),
	PRIMARY KEY(name, command)
)`
)

var createQueries = []string{
	createTableCommand,
	createTableStatus,
}

type SQLiteStore struct {
	ctx context.Context
	db  *sqlx.DB
}

func New(file string) (*SQLiteStore, error) {
	return NewWithContext(context.Background(), file)
}

func NewWithContext(ctx context.Context, file string) (*SQLiteStore, error) {
	db, err := sqlx.Connect("sqlite3", file)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"db":   "sqlite3",
			"file": file,
		}).Error("Could not open storage")
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"db":   "sqlite3",
			"file": file,
		}).Error("Could not begin tx")
		return nil, err
	}
	defer tx.Rollback()

	for i := range createQueries {
		_, err = tx.Exec(createQueries[i])
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"query": createQueries[i],
			}).Error("Could not execute create query")

			return nil, err
		}
	}
	tx.Commit()

	return &SQLiteStore{
		ctx: ctx,
		db:  db,
	}, nil
}

func (s *SQLiteStore) Context() context.Context {
	return s.ctx
}

func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
