package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)


type DB struct {
	conn *sql.DB
}

func NewConn() (DB, error){
	db, err := sql.Open("sqlite3", "project.db")
	if err != nil {
		return DB{}, err
	}
	dataConn := DB{
		db,
	}

	return dataConn, nil
}


func init() {
	db, err := sql.Open("sqlite3", "project.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
	create table if not exists products(
		name string, 
		description string, 
		price numeric(8,2)
);
	create table if not exists users (
		username varchar(50) not null,
		password varchar(50) not null,
		email unique not null,
		address unique not null
	);
`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database and tables are present")
}


