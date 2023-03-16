package main

import (
	"CEC/pgk/config"
	"CEC/pgk/helper"
	"CEC/pgk/storage/mongodb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Environment Variable
	config.SetConfig()

	result := mongodb.Conn.UpdateByID("CEC", "log", "641196ac5986aae6482be366", bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: "new title"},
		}},
	})
	res, _ := helper.Marshal(result)
	fmt.Println(string(res))
}
