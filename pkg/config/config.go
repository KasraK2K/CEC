package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	PORT                    string `json:"port"`
	MODE                    string `json:"mode"`
	PREFORK                 bool   `json:"prefork"`
	BACKEND_VERSION         string `json:"backend_version"`
	FRONTEND_VERSION        string `json:"frontend_version"`
	APP_VERSION             string `json:"app_version"`
	STDOUT_LOGS             string `json:"stdout_logs"`
	FILE_LOGS               string `json:"file_logs"`
	DB_HOST                 string `json:"db_host"`
	DB_PORT                 string `json:"db_port"`
	DB_PASSWORD             string `json:"dn_password"`
	DB_USER                 string `json:"db_user"`
	DB_NAME                 string `json:"db_name"`
	DB_TIMEZONE             string `json:"db_timezone"`
	DB_SSL_MODE             string `json:"db_ssl_mode"`
	MONGODB_URI             string `json:"mongodb_uri"`
	JWT_SIGNING_KEY         string `json:"jwt_signing_key"`
	MAILGUN_PRIVATE_API_KEY string `json:"mailgun_private_api_key"`
	MAILGUN_DOMAIN          string `json:"mailgun_domain"`
	MAILGUN_API_BASE        string `json:"mailgun_api_base"`
	MAILGUN_SENDER          string `json:"mailgun_sender"`
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
	AppConfig.STDOUT_LOGS = os.Getenv("STDOUT_LOGS")
	AppConfig.FILE_LOGS = os.Getenv("FILE_LOGS")
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
	// JWT
	AppConfig.JWT_SIGNING_KEY = os.Getenv("JWT_SIGNING_KEY")
	// MailGun
	AppConfig.MAILGUN_PRIVATE_API_KEY = os.Getenv("MAILGUN_PRIVATE_API_KEY")
	AppConfig.MAILGUN_DOMAIN = os.Getenv("MAILGUN_DOMAIN")
	AppConfig.MAILGUN_API_BASE = os.Getenv("MAILGUN_API_BASE")
	AppConfig.MAILGUN_SENDER = os.Getenv("MAILGUN_SENDER")
}
