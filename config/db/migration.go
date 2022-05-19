package db

import (
	"fmt"
	"go-demo/config"
	"go-demo/internal/util/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	client *migrate.Migrate
}

func newMigration() *Migration {
	m := Migration{}
	path := "file://_assets/db/migration"
	var err error
	var dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=%s",
		config.Env.GetString("postgres.user"),
		config.Env.GetString("postgres.password"),
		config.Env.GetString("postgres.host"),
		config.Env.GetInt("postgres.port"),
		config.Env.GetString("postgres.database"),
		config.Env.GetString("postgres.schema"),
	)
	m.client, err = migrate.New(path, dbUrl)
	if err != nil {
		logger.Error.Panic(err)
	}
	return &m
}

func (m *Migration) To(targetVersion uint) {
	if err := m.client.Migrate(targetVersion); err != nil && err != migrate.ErrNoChange {
		logger.Error.Panic(err)
	}
	afterVersion, _, _ := m.client.Version()
	logger.Info.Printf("Migration to version:%d success", afterVersion)
}

func (m *Migration) Up() {
	if err := m.client.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error.Panic(err)
	}
	afterVersion, _, _ := m.client.Version()
	logger.Info.Printf("Migration up version:%d success", afterVersion)
}

func (m *Migration) Down() {
	if err := m.client.Down(); err != nil && err != migrate.ErrNoChange {
		logger.Error.Panic(err)
	}
	version, _, _ := m.client.Version()
	logger.Info.Printf("Migration down version:%d success", version)
}
