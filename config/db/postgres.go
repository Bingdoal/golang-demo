package db

import (
	"fmt"
	"go-demo/config"
	"go-demo/internal/util/logger"

	"github.com/spf13/viper"
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

func GetDB() *gorm.DB {
	return DB
}

func NewDB(config *viper.Viper) *gorm.DB {
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Get("postgres.host"),
		config.Get("postgres.user"),
		config.GetString("postgres.password"),
		config.Get("postgres.database"),
		config.GetInt("postgres.port"),
	)
	db := connectDB(dsn)
	logger.Info.Println("postgres connection established.")
	if config.GetBool("migration.enabled") {
		logger.Debug.Println("db migration start.")
		migration()
		logger.Debug.Println("db migration end.")
	}
	return db
}

func InitDB() {
	DB = NewDB(config.Env)
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
