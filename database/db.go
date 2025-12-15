package database

import "github.com/lib/pq"

type DB struct {
	*pq.Connector
}

func New(credentials string) (*DB, error) {
	connector, err := pq.NewConnector(credentials)
	if err != nil {
		return nil, err
	}
	return &DB{
		connector,
	}, nil
}
