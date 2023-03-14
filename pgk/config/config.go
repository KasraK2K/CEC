package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port        string `json:"port"`
	DB_HOST     string `json:"db_host"`
	DB_PORT     string `json:"db_port"`
	DB_PASSWORD string `json:"dn_password"`
	DB_USER     string `json:"db_user"`
	DB_NAME     string `json:"db_name"`
	DB_TIMEZONE string `json:"db_timezone"`
	DB_SSL_MODE string `json:"db_ssl_mode"`
	MONGODB_URI string `json:"mongodb_uri"`
}

var AppConfig config

func SetConfig() {
	err := godotenv.Load("pgk/config/.env")
	if err != nil {
		log.Fatal(err)
	}

	// Application
	AppConfig.Port = os.Getenv("PORT")
	// PostgreSQL Database
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_PORT = os.Getenv("DB_PORT")
	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	AppConfig.DB_USER = os.Getenv("DB_USER")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
	AppConfig.DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	// MongoDB Database
	AppConfig.MONGODB_URI = os.Getenv("MONGODB_URI")
}
