package db

import (
	"fmt"
	"go-demo/config"
	"go-demo/internal/util/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dsn = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Taipei",
	config.Env.Get("postgres.host"),
	config.Env.Get("postgres.user"),
	config.Env.GetString("postgres.password"),
	config.Env.Get("postgres.database"),
	config.Env.GetInt("postgres.port"),
)

var pgcon = postgres.New(postgres.Config{
	DSN:                  dsn,
	PreferSimpleProtocol: true,
})

func connectDB() *gorm.DB {
	db, err := gorm.Open(pgcon, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error.Println("Connect DB failed: ", err)
		panic(err)
	}

	db.Exec(fmt.Sprintf("SET search_path='%s'", config.Env.Get("postgres.schema")))

	return db.Session(&gorm.Session{PrepareStmt: true})
}

var Postgres *gorm.DB

func init() {
	Postgres = connectDB()
	migration()
}

func migration() {

}
