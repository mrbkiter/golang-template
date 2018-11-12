package repo

import (
	"database/sql"
	"log"
	"os"

	"template.github.com/server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"template.github.com/server/web"
)

type Repo struct{}

type context = web.Context

var driver *sql.DB

var cf *config.DatabaseConfig

func Init() {
	cf = config.GetConfig().DatabaseConfig
	if cf == nil {
		log.Fatalln("Failed to initialize database. Database Config need not nil")
		os.Exit(1)
	}

	switch cf.Driver {
	case "sqlite3":
		initSqlLite()
		break
	case "postgres":
		initPostgres()
		break
	}
}

//OpenConnection open a connection
func (*Repo) OpenConnection(ctx context) *gorm.DB {
	db, _ := gorm.Open(cf.Driver, driver)
	return db
}

//Init initialize database driver
func initPostgres() {
	DB, error := sql.Open(cf.Driver, cf.JdbcUrl)
	if error == nil {
		panic(error)
	}
	DB.SetConnMaxLifetime(60000)
	DB.SetMaxIdleConns(30)
	DB.SetMaxOpenConns(100)
}

//Init initialize database driver
func initSqlLite() {
	DB, error := sql.Open(cf.Driver, cf.JdbcUrl)
	if error == nil {
		panic(error)
	}
	DB.SetConnMaxLifetime(60000)
	DB.SetMaxIdleConns(30)
	DB.SetMaxOpenConns(100)
}
