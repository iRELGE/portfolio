package dbmanager

import (
	"database/sql"
	"fmt"

	"rabie.com/portfolio/errorrepo"
)

// Connectdb : connect db
func Connectdb() (*sql.DB, error) {
	//open db (regular sql open call)
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Portfolio")
	err = testping(db)
	errorrepo.DieIf(err)

	//close deferred
	return db, err

}

// Testping : testing db ping
func testping(db *sql.DB) error {
	//check if db is connected
	err := db.Ping()

	fmt.Println("connected")
	return err
}
