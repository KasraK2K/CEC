package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	PORT             string `json:"port"`
	MODE             string `json:"mode"`
	PREFORK          bool   `json:"prefork"`
	BACKEND_VERSION  string `json:"backend_version"`
	FRONTEND_VERSION string `json:"frontend_version"`
	APP_VERSION      string `json:"app_version"`
	DB_HOST          string `json:"db_host"`
	DB_PORT          string `json:"db_port"`
	DB_PASSWORD      string `json:"dn_password"`
	DB_USER          string `json:"db_user"`
	DB_NAME          string `json:"db_name"`
	DB_TIMEZONE      string `json:"db_timezone"`
	DB_SSL_MODE      string `json:"db_ssl_mode"`
	MONGODB_URI      string `json:"mongodb_uri"`
}

var AppConfig config

func SetConfig() {
	err := godotenv.Load("pkg/config/.env")
	if err != nil {
		log.Fatal(err)
	}

	prefork, err := strconv.ParseBool(os.Getenv("PREFORK"))
	if err != nil {
		log.Fatal(err)
	}

	// Application
	AppConfig.PORT = os.Getenv("PORT")
	AppConfig.MODE = os.Getenv("MODE")
	AppConfig.PREFORK = prefork
	AppConfig.BACKEND_VERSION = os.Getenv("BACKEND_VERSION")
	AppConfig.FRONTEND_VERSION = os.Getenv("FRONTEND_VERSION")
	AppConfig.APP_VERSION = os.Getenv("APP_VERSION")
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
