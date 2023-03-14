package main

import (
	"CEC/pgk/config"
	"CEC/pgk/storage/mongo"
	"net/http"
	"time"
)

type Person struct {
	Title    string      `json:"title" bson:"title"`
	Message  string      `json:"message" bson:"message,omitempty"`
	Priority int         `json:"priority" bson:"priority"`
	Data     interface{} `json:"data" bson:"data,omitempty"`
}

func main() {
	// Environment Variable
	config.SetConfig()

	mongo.Conn.InsertOne("CEC", "log", Person{
		Title:   "Test Log",
		Message: "Test our log system by kasra's code",
		Data: struct {
			Creator string
			Time    time.Time
		}{
			Creator: "Kasra",
			Time:    time.Now(),
		},
	})

	http.ListenAndServe(":3000", nil)
}
