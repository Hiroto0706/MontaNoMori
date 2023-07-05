package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo/config"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Db  *sql.DB
	err error
)

const (
	tableNameUser      = "users"
	tableNameImage     = "images"
	tableNameTag       = "tags"
	tableNameImageTags = "image_tags"
)

func InitDB() {
	Db, err = ConnectSqlite()
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		uuid TEXT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at datetime)`, tableNameUser)

	Db.Exec(cmdU)

	cmdI := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		title TEXT,
		description TEXT,
		url_path TEXT,
		updated_at datetime,
		created_at datetime)`, tableNameImage)

	Db.Exec(cmdI)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		name TEXT,
		updated_at datetime,
		created_at datetime)`, tableNameTag)

	Db.Exec(cmdT)

	cmdIT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		image_id INTEGER,
		tag_id INTEGER,
		updated_at datetime,
		created_at datetime,
		FOREIGN KEY (image_id) REFERENCES images(id),
		FOREIGN KEY (tag_id) REFERENCES tags(id),
		UNIQUE(image_id, tag_id))`, tableNameImageTags)

	Db.Exec(cmdIT)
}

func ConnectSqlite() (*sql.DB, error) {
	db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
