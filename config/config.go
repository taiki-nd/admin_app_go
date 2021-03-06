package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Logfile         string
	SqlDevelop      string
	HostDevelop     string
	PortDevelop     string
	NameDevelop     string
	UserDevelop     string
	PasswordDevelop string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Panicf("failed to load config.ini: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Logfile:         cfg.Section("admin_app_go").Key("log_file").String(),
		SqlDevelop:      cfg.Section("db_development").Key("sql_develop").String(),
		HostDevelop:     cfg.Section("db_development").Key("host_develop").String(),
		PortDevelop:     cfg.Section("db_development").Key("port_develop").String(),
		NameDevelop:     cfg.Section("db_development").Key("name_develop").String(),
		UserDevelop:     cfg.Section("db_development").Key("user_develop").String(),
		PasswordDevelop: cfg.Section("db_development").Key("password_develop").String(),
	}
}
