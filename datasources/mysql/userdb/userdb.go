package userdb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	postgresdbUsername = "postgresdb_username"
	postgresdbPassword = "postgresdb_password"
	postgresdbHostname = "postgresdb_hostname"
	postgresdbDbname   = "postgresdb_dbname"
)

var (
	DBClient *sql.DB
	username = os.Getenv(postgresdbUsername)
	password = os.Getenv(postgresdbPassword)
	hostname = os.Getenv(postgresdbHostname)
	dbName   = os.Getenv(postgresdbDbname)
)

func init() {

	datasourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username, password, hostname, dbName)
	var connErr error
	DBClient, connErr = sql.Open("postgres", datasourceName)
	if connErr != nil {
		panic(connErr) // We do not start the application if we have problems connecting to the DB
	}
	if err := DBClient.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connection successful")
}

// Notes : db, err := sql.Open("postgres", "postgres://username:password@localhost/db_name?sslmode=disable")
