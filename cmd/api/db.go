package main

import (
	"database/sql"
	"log"

	_ "github.com/godror/godror"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Oracle!")
	return connection, nil
}
