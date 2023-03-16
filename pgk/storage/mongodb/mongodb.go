package mongodb

import (
	"CEC/pgk/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connection struct{}

var Conn connection

type InsertOneResult struct {
	ID interface{} `json:"id" bson:"id"`
}
type InsertManyResult struct {
	IDS []interface{} `json:"ids" bson:"ids"`
}

func (c *connection) Connect() *mongo.Client {
	uri := config.AppConfig.MONGODB_URI

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Panic(err)
	}

	return client
}

func (c *connection) Disconnect(client *mongo.Client) {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Panic(err)
		}
	}()
}

func (c *connection) Ping(client *mongo.Client) {
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

/* -------------------------------------------------------------------------- */
/*                                 Insert One                                 */
/* -------------------------------------------------------------------------- */
// document := struct {
// 	Title   string
// 	Message string
// 	Data    interface{}
// }{
// 	Title:   "Test Log",
// 	Message: "Test our log system by kasra's code",
// 	Data: struct {
// 		Name string `json:"name" bson:"name"`
// 	}{"Kasra"},
// }
// result := mongodb.Conn.InsertOne("CEC", "log", document)
// res, _ := helper.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) InsertOne(database, collection string, document interface{}) InsertOneResult {
	client := Conn.Connect()
	defer Conn.Disconnect(client)

	coll := client.Database(database).Collection(collection)

	result, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		log.Panic(err)
	}

	return InsertOneResult{ID: result.InsertedID}
}

/* -------------------------------------------------------------------------- */
/*                                 Insert Many                                */
/* -------------------------------------------------------------------------- */
// documents := []interface{}{
// 	struct {
// 		Title   string
// 		Message string
// 	}{Title: "Insert Many Title", Message: "My Insert Many Message"},
// 	struct {
// 		Name string
// 		Age  int
// 	}{Name: "Kasra", Age: 37},
// }
// result := mongodb.Conn.InsertMany("CEC", "log", documents)
// res, _ := helper.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) InsertMany(database, collection string, documents []interface{}, opts ...*options.InsertManyOptions) InsertManyResult {
	client := Conn.Connect()
	defer Conn.Disconnect(client)

	coll := client.Database(database).Collection(collection)

	results, err := coll.InsertMany(context.TODO(), documents, opts...)
	if err != nil {
		log.Panic(err)
	}

	return InsertManyResult{IDS: results.InsertedIDs}
}

/* -------------------------------------------------------------------------- */
/*                                  Find One                                  */
/* -------------------------------------------------------------------------- */
// filter := bson.D{{Key: "id", Value: "641196ac5986aae6482be366"}}
// opts := options.FindOne()
// result := mongodb.Conn.FindOne("CEC", "log", filter, opts)
// res, _ := helper.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) FindOne(database, collection string, filter bson.D, opts ...*options.FindOneOptions) bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := Conn.Connect()
	defer Conn.Disconnect(client)

	Conn.sanitizeFilter(&filter)

	coll := client.Database(database).Collection(collection)

	var result bson.M
	coll.FindOne(ctx, filter, opts...).Decode(&result)

	return result
}

/* -------------------------------------------------------------------------- */
/*                                    Find                                    */
/* -------------------------------------------------------------------------- */
// filter := bson.D{{Key: "priority", Value: 0}}
// opts := options.Find()
// opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
// results := mongodb.Conn.Find("CEC", "log", filter, opts)
// res, _ := helper.Marshal(results)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) Find(database, collection string, filter bson.D, opts ...*options.FindOptions) []bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := Conn.Connect()
	defer Conn.Disconnect(client)

	Conn.sanitizeFilter(&filter)

	coll := client.Database(database).Collection(collection)

	cursor, err := coll.Find(context.TODO(), filter, opts...)
	if err != nil {
		log.Panic(err)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Panic(err)
	}

	return results
}

/* ----------------------------- Sanitize Filter ---------------------------- */
// This filter have been created to get bson.D and change id to ObjectId
/* -------------------------------------------------------------------------- */
func (c *connection) sanitizeFilter(filter *bson.D) {
	reservedFilter := *filter
	for index, filterItem := range reservedFilter {
		if filterItem.Key == "id" || filterItem.Key == "_id" {
			// Check id is valid or not
			isValidId := primitive.IsValidObjectID(filterItem.Value.(string))
			if !isValidId {
				message := fmt.Sprintf("id %s is not valid", filterItem.Value)
				log.Println(message)
				return
			}
			// Change id to ObjectId
			docID, err := primitive.ObjectIDFromHex(filterItem.Value.(string))
			if err != nil {
				message := fmt.Sprintf("hex string `%s` is not a valid ObjectID", filterItem.Value)
				log.Println(message)
				return
			}
			reservedFilter[index].Key = "_id"
			reservedFilter[index].Value = docID
		}
	}
}
