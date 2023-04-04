package Config

import (
	"log"

	"github.com/go-ini/ini"
)

var cfg *ini.File

type MysqlParams struct {
	Host     string
	User     string
	Password string
	Name     string
}

func LoadConfig() MysqlParams {
	var err error

	cfg, err = ini.Load("config/cfg.ini")
	if err != nil {
		log.Fatal(2, "failed to load config: %v", err)
	}

	mps, err := cfg.GetSection("mysql")

	cfgParams := MysqlParams{
		Host:     mps.Key("host").String(),
		User:     mps.Key("user").String(),
		Password: mps.Key("password").String(),
		Name:     mps.Key("name").String(),
	}

	return cfgParams
}
