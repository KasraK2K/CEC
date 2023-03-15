package mongo

import (
	"CEC/pgk/config"
	"CEC/pgk/helper"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct{}

var Conn Connection

type SingleResponse struct {
	ID primitive.ObjectID `json:"id" bson:"id"`
}

func (c *Connection) Connect() *mongo.Client {
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

func (c *Connection) Disconnect(client *mongo.Client) {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Panic(err)
		}
	}()
}

func (c *Connection) Ping(client *mongo.Client) {
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func (c *Connection) InsertOne(database, collection string, s interface{}) SingleResponse {
	client := Conn.Connect()
	defer Conn.Disconnect(client)

	coll := client.Database(database).Collection(collection)

	result, err := coll.InsertOne(context.TODO(), s)
	if err != nil {
		log.Panic(err)
	}

	return SingleResponse{ID: result.InsertedID.(primitive.ObjectID)}
}

func (c *Connection) Find(database, collection string, filter primitive.D, opts ...*options.FindOptions) []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := Conn.Connect()
	defer Conn.Disconnect(client)

	coll := client.Database(database).Collection(collection)

	cursor, err := coll.Find(context.TODO(), filter, opts...)
	if err != nil {
		log.Panic(err)
	}
	defer cursor.Close(ctx)

	var results []primitive.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Panic(err)
	}

	return results
}

func (c *Connection) BsonToJson(results []primitive.M) string {
	var convertedResult []map[string]interface{}
	for _, item := range results {
		convertedItem := make(map[string]interface{})
		for key, val := range item {
			switch v := val.(type) {
			case primitive.ObjectID:
				convertedItem[key] = v.Hex()
			case bson.M:
				subMap := make(map[string]interface{})
				for subKey, subVal := range v {
					subMap[subKey] = subVal
				}
				convertedItem[key] = subMap
			case bson.A:
				subArray := make([]interface{}, len(v))
				for i, subVal := range v {
					subArray[i] = subVal
				}
				convertedItem[key] = subArray
			default:
				convertedItem[key] = v
			}
		}
		convertedResult = append(convertedResult, convertedItem)
	}

	jsonBytes, err := helper.Marshal(convertedResult)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)

	return jsonString
}

func (c *Connection) CursorToM(ctx context.Context, cursor *mongo.Cursor) []primitive.M {
	var results []bson.M
	for cursor.Next(ctx) {
		var document bson.M
		err := cursor.Decode(&document)
		if err != nil {
			log.Panic(err)
		}
		results = append(results, document)
	}

	return results
}
