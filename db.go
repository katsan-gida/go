package katsan

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var db *DB
var dbOnce sync.Once

type DB struct {
	db      *sqlx.DB
	queries *struct {
		LoadActiveOrders string `query:"LoadActiveOrders"`
	}
}

func (d *DB) GetDB() *sqlx.DB {
	return d.db
}

type dbConfig struct {
	driver           string
	connectionString string
}

func WithPostgres(driver, connectionString string) *dbConfig {
	dbc := dbConfig{}
	dbc.driver = driver
	dbc.connectionString = connectionString
	return &dbc
}

func NewDB(config *dbConfig) *DB {
	dbOnce.Do(func() {
		db = &DB{
			db: sqlx.MustConnect(config.driver, config.connectionString),
		}
	})
	return db
}
