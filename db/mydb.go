package db

import (
	"database/sql"

	"gitlab.com/parkby/parkby-service/config"
	"gitlab.com/parkby/parkby-service/log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// MyDB structure
type MyDB struct {
	*sql.DB
}

var (
	myDB         MyDB
	maxIdleConns = 3
	maxOpenConns = 10
)

// GetDB - get MyDB instance
func GetDB() MyDB {
	return myDB
}

func init() {
	log.Info.Println("init myDB")

	credential := config.Config.MySQLUser + ":" + config.Config.MySQLPass
	dbName := config.Config.MySQLDBName

	db, err := sql.Open("mysql", credential+"@/"+dbName)
	if err != nil {
		log.Error.PrintlnErr(err)
		panic(err)
	}

	myDB = MyDB{db}

	myDB.configure()
}

func (myDB *MyDB) configure() {
	myDB.SetMaxIdleConns(maxIdleConns)
	myDB.SetMaxOpenConns(maxOpenConns)
}
