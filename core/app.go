package core

import (
	"fmt"
	"log/slog"
	"os"
	"webengine/database"
	queries "webengine/database/gen"
)

type Version string

const (
	VersionDev  Version = "dev"
	VersionProd Version = "prod"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type Application interface {
	Database() database.Database
	Query() *queries.Queries
	Logger() Logger
}

type DevApp struct {
	Db  database.Database
	log *slog.Logger
}

func (a *DevApp) Database() database.Database {
	return a.Db
}

func (a *DevApp) Query() *queries.Queries {
	return a.Db.Queries()
}

func (a *DevApp) Logger() Logger {
	return a.log
}

type ProdApp struct {
	Db  database.Database
	log *slog.Logger
}

func (a *ProdApp) Database() database.Database {
	return a.Db
}

func (a *ProdApp) Query() *queries.Queries {
	return a.Db.Queries()
}

func (a *ProdApp) Logger() Logger {
	return a.log
}

func NewApplication(v Version) (Application, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))
	switch v {
	case VersionDev:
		{
			db, err := database.NewDevDatabase()
			if err != nil {
				return nil, err
			}
			return &DevApp{
				Db:  db,
				log: logger,
			}, nil
		}
	case VersionProd:
		{
			db, err := database.NewProdDatabase()
			if err != nil {
				return nil, err
			}
			return &ProdApp{
				Db:  db,
				log: logger,
			}, nil
		}
	default:
		return nil, fmt.Errorf("unknown version: %s", v)
	}
}
