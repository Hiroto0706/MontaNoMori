package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	SQLDriver string
	DbName    string
	Json      string
	Bucket    string
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
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		Json:      cfg.Section("gcs").Key("json").String(),
		Bucket:    cfg.Section("gcs").Key("bucket").String(),
	}
}
