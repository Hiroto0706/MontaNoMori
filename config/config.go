package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	User     string
	Password string
	DbName   string
	Json     string
	Bucket   string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		User:     cfg.Section("db").Key("user").String(),
		Password: cfg.Section("db").Key("password").String(),
		DbName:   cfg.Section("db").Key("dbname").String(),
		Json:     cfg.Section("gcs").Key("json").String(),
		Bucket:   cfg.Section("gcs").Key("bucket").String(),
	}
}
