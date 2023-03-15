package main

import (
	"CEC/pgk/config"
	"CEC/pgk/helper"
	"CEC/pgk/storage/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type object struct {
	Data bson.RawValue `bson:"data"`
}

var d interface{}

func main() {
	// Environment Variable
	config.SetConfig()

	/* -------------------------------------------------------------------------- */
	/*                                 Insert One                                 */
	/* -------------------------------------------------------------------------- */
	// person := new(model.Person)
	// result, jsonResult := person.InsertOne("CEC", "log", model.Person{
	// 	Title:   "Test Log",
	// 	Message: "Test our log system by kasra's code",
	// 	Data: struct {
	// 		Name string `json:"name" bson:"name"`
	// 	}{"Kasra"},
	// }, true)
	// fmt.Println("result:", result)
	// fmt.Println("jsonResult:", jsonResult)

	/* -------------------------------------------------------------------------- */

	/* -------------------------------------------------------------------------- */
	/*                                    Find                                    */
	/* -------------------------------------------------------------------------- */
	// person := new(model.Person)
	// filter := bson.D{{Key: "priority", Value: 0}}
	// opts := options.Find()
	// opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	// results, _ := person.Find("CEC", "log", filter, opts, true)
	// fmt.Println(results)
	// fmt.Println(jsonPayload)
	/* -------------------------------------------------------------------------- */

	filter := bson.D{{Key: "priority", Value: 0}}
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	results := mongo.Conn.Find("CEC", "log", filter, opts)

	res, _ := helper.Marshal(results)
	fmt.Println(string(res))
}
