package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", "project.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	create table if not exists products(
		id int primary key,
		name string, 
		description string, 
		price numeric(8,2)
);
	create table if not exists users (
		id int primary key,
		username varchar(50),
		password varchar(50)
	);
`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database and tables are present")
}

