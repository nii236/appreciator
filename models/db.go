package models

import (
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

// DB Contains an instance of the DB
type DB struct {
	*ledis.DB
}

// DBer contains all the interfaces from the models
type DBer interface {
	Guester
}

// NewDB creates a new instance of a database
func NewDB() *DB {
	cfg := lediscfg.NewConfigDefault()
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(0)
	return &DB{db}
}
