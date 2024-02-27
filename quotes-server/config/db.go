package config

import (
	"go.etcd.io/bbolt"
)

func CreateDB() *bbolt.DB {
	db, err := bbolt.Open("quotes.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	return db
}
