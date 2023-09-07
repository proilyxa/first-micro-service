package database

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/exp/slog"
	"os"
	"path"
)

//go:embed all:sqlite/migrations
var StaticFiles embed.FS

type Storage struct {
	Db *sql.DB
}

func (s *Storage) Close() {
	err := s.Db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Storage) Begin() (*sql.Tx, error) {
	return s.Db.Begin()
}

func NewSqliteConnection(log *slog.Logger) (*Storage, error) {
	const op = "sqlite.New"

	//dbDir, err := osext.ExecutableFolder()
	//if err != nil {
	//	log.Fatal(err)
	//}

	dbDir, err := os.Getwd()
	storage := path.Join(dbDir, "storage.db")
	db, err := sql.Open("sqlite3", storage)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	addr := fmt.Sprintf("sqlite3://file:%s", storage)
	driver, err := iofs.New(StaticFiles, path.Join("sqlite", "migrations"))
	migration, err := migrate.NewWithSourceInstance(
		"iofs",
		driver,
		addr,
	)

	if err := migration.Up(); err != nil {
		log.Info("Database migrations: " + err.Error())
	}

	return &Storage{Db: db}, nil
}
