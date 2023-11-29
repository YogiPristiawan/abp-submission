package main

import "os"

type AppConfig struct {
	AppName       string
	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDB       string
}

func ParseConfig() *AppConfig {
	var cfg = AppConfig{
		AppName:       "Todo",
		MySQLHost:     "127.0.0.1",
		MySQLPort:     "3306",
		MySQLUser:     "root",
		MySQLPassword: "password",
		MySQLDB:       "todo",
	}

	if val, exists := os.LookupEnv("APP_NAME"); exists {
		cfg.AppName = val
	}

	if val, exists := os.LookupEnv("MYSQL_HOST"); exists {
		cfg.MySQLHost = val
	}

	if val, exists := os.LookupEnv("MYSQL_PORT"); exists {
		cfg.MySQLPort = val
	}

	if val, exists := os.LookupEnv("MYSQL_USER"); exists {
		cfg.MySQLUser = val
	}

	if val, exists := os.LookupEnv("MYSQL_PASSWORD"); exists {
		cfg.MySQLPassword = val
	}

	if val, exists := os.LookupEnv("MYSQL_DBNAME"); exists {
		cfg.MySQLDB = val
	}

	return &cfg
}
