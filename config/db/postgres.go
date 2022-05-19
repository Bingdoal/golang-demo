package db

import (
	"fmt"
	"go-demo/config"
	"go-demo/internal/util/logger"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func connectDB(dsn string) *gorm.DB {
	var pgcon = pg.New(pg.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(pgcon, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error.Println("Connect DB failed: ", err)
		panic(err)
	}

	db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", config.Env.Get("postgres.schema")))
	db.Exec(fmt.Sprintf("SET search_path='%s'", config.Env.Get("postgres.schema")))

	return db.Session(&gorm.Session{PrepareStmt: true})
}

var DB *gorm.DB

func InitDB() {
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Env.Get("postgres.host"),
		config.Env.Get("postgres.user"),
		config.Env.GetString("postgres.password"),
		config.Env.Get("postgres.database"),
		config.Env.GetInt("postgres.port"),
	)
	DB = connectDB(dsn)
	logger.Info.Println("postgres connection established.")
	if config.Env.GetBool("migration.enabled") {
		logger.Debug.Println("db migration start.")
		migration()
		logger.Debug.Println("db migration end.")
	}
}

func migration() {
	migration := newMigration()
	target := config.Env.GetString("migration.target")
	if target == "latest" {
		migration.Up()
	} else {
		migration.To(config.Env.GetUint("migration.target"))
	}
}
