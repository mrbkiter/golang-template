package repo

import (
	ctx1 "context"
	"database/sql"
	"log"
	"os"

	"template.github.com/server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

type context = ctx1.Context

type Repository struct{}

var driver *sql.DB

var cf *config.DatabaseConfig

var repo *Repository

//Init initialize repository, then return Repo service
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
	repo = &Repository{}
}

//Repo return repository object
func Repo() *Repository {
	if repo == nil {
		Init()
	}
	return repo

}

//OpenConnection open a connection
func (*Repository) openConnection(ctx *context) *gorm.DB {
	db, _ := gorm.Open(cf.Driver, driver)
	return db
}

//Init initialize database driver
func initPostgres() {
	//fmt.Println(cf)
	DB, error := sql.Open(cf.Driver, cf.JdbcUrl)
	if error != nil {
		panic(error)
	}
	DB.SetConnMaxLifetime(60000)
	DB.SetMaxIdleConns(30)
	DB.SetMaxOpenConns(100)
	driver = DB
}

//Init initialize database driver
func initSqlLite() {
	DB, error := sql.Open(cf.Driver, cf.JdbcUrl)
	if error != nil {
		panic(error)
	}
	DB.SetConnMaxLifetime(60000)
	DB.SetMaxIdleConns(30)
	DB.SetMaxOpenConns(100)
	driver = DB
}
