package pg

import (
	"CEC/pgk/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Conn *gorm.DB
	DSN  any
}

type Environment struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Port     int    `json:"port"`
	SSLMode  string `json:"ssl_mode"`
	TimeZone string `json:"time_zone"`
}

func Connect() Connection {
	appConfig := config.AppConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		appConfig.DB_HOST,
		appConfig.DB_USER,
		appConfig.DB_PASSWORD,
		appConfig.DB_NAME,
		appConfig.DB_PORT,
		appConfig.DB_SSL_MODE,
		appConfig.DB_TIMEZONE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return Connection{Conn: db}
}

func Migrate(modelStruct interface{}) {
	var db = Connect().Conn

	//Create Table If Not Exist
	err := db.AutoMigrate(&modelStruct)
	if err != nil {
		return
	}
}
