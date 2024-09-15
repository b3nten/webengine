package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
	queries "webengine/database/gen"
)

type Database interface {
	KvGet(key string) (string, error)
	KvSet(value string) error
	Raw(query string) (sql.Result, error)
	Queries() *queries.Queries
}

type DevDatabase struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewDevDatabase() (*DevDatabase, error) {
	db, err := sql.Open("sqlite", "dev.db")

	q := queries.New(db)

	err = Migrate(db)

	if err != nil {
		return nil, err
	}

	return &DevDatabase{db: db, queries: q}, nil
}

func (d *DevDatabase) KvGet(key string) (string, error) {
	var result string
	err := d.db.QueryRow("SELECT value FROM kv WHERE key = ?", key).Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (d *DevDatabase) KvSet(value string) error {
	return d.db.QueryRow("INSERT OR REPLACE INTO kv VALUES (?)", value).Err()
}

func (d *DevDatabase) Raw(query string) (sql.Result, error) {
	return d.db.Exec(query)
}

func (d *DevDatabase) Queries() *queries.Queries {
	return d.queries
}

type ProdDatabase struct {
	db      *sql.DB
	queries *queries.Queries
}

func NewProdDatabase() (*ProdDatabase, error) {
	db, err := sql.Open("sqlite", "prod.db")

	err = Migrate(db)

	queries := queries.New(db)

	if err != nil {
		return nil, err
	}

	return &ProdDatabase{db: db, queries: queries}, nil
}

func (d *ProdDatabase) KvGet(key string) (string, error) {
	var result string
	err := d.db.QueryRow("SELECT value FROM kv WHERE key = ?", key).Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (d *ProdDatabase) KvSet(value string) error {
	return d.db.QueryRow("INSERT OR REPLACE INTO kv VALUES (?)", value).Err()
}

func (d *ProdDatabase) Raw(query string) (sql.Result, error) {
	return d.db.Exec(query)
}

func (d *ProdDatabase) Queries() *queries.Queries {
	return d.queries
}
