package datastore

import (
	"database/sql"
	"errors"

	"github.com/anthonynsimon/parrot/common/datastore/postgres"
	"github.com/anthonynsimon/parrot/common/model"
)

type Store interface {
	model.LocaleStorer
	model.ProjectStorer
	model.ProjectLocaleStorer
	model.UserStorer
	model.ProjectUserStorer
	model.ProjectClientStorer
	Ping() error
	Close() error
}

var (
	ErrNoDB           = errors.New("couldn't get DB")
	ErrNotImplemented = errors.New("database not implemented")
)

type Datastore struct {
	Store
}

func NewDatastore(name string, url string) (*Datastore, error) {
	var ds *Datastore

	switch name {
	case "postgres":
		conn, err := sql.Open("postgres", url)
		if err != nil {
			return nil, err
		}

		p := &postgres.PostgresDB{DB: conn}
		// TODO(anthonynsimon): debug refused connections when db connections > 1
		p.SetMaxIdleConns(1)
		p.SetMaxOpenConns(1)

		ds = &Datastore{p}
	default:
		return nil, ErrNotImplemented
	}

	return ds, nil
}