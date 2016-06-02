package models

import (
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

type DB struct {
	*ledis.DB
}

// NewDB creates a new instance of a database
func NewDB() *DB {
	cfg := lediscfg.NewConfigDefault()
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(0)
	return &DB{db}
}
