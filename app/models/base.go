package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo/config"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
	tableNameUser      = "users"
	tableNameImage     = "images"
	tableNameTag       = "tags"
	tableNameImageTags = "image_tags"
)

func InitDB() {
	name := config.Config.User
	password := config.Config.Password
	dbname := config.Config.DbName

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", name, password, dbname)

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// defer Db.Close()

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		uuid VARCHAR(255),
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW())`, tableNameUser)

	_, err = Db.Exec(cmdU)

	if err != nil {
		log.Fatal(err)
	}

	cmdI := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		description VARCHAR(255),
		url_path VARCHAR(255),
		updated_at TIMESTAMP DEFAULT NOW(),
		created_at TIMESTAMP DEFAULT NOW())`, tableNameImage)

	_, err = Db.Exec(cmdI)

	if err != nil {
		log.Fatal(err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		updated_at TIMESTAMP DEFAULT NOW(),
		created_at TIMESTAMP DEFAULT NOW())`, tableNameTag)

	_, err = Db.Exec(cmdT)

	if err != nil {
		log.Fatal(err)
	}

	cmdIT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		image_id INTEGER REFERENCES images(id),
		tag_id INTEGER REFERENCES tags(id),
		UNIQUE(image_id, tag_id),
		updated_at TIMESTAMP DEFAULT NOW(),
		created_at TIMESTAMP DEFAULT NOW())`, tableNameImageTags)

	_, err = Db.Exec(cmdIT)

	if err != nil {
		log.Fatal(err)
	}
}

func ConnectPostgres(user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
