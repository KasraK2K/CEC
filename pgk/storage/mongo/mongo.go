package mongo

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

/* -------------------------------------------------------------------------- */
/*                                 Insert One                                 */
/* -------------------------------------------------------------------------- */
// result := mongo.Conn.InsertOne("CEC", "log", struct {
// 	Title   string
// 	Message string
// 	Data    interface{}
// }{
// 	Title:   "Test Log",
// 	Message: "Test our log system by kasra's code",
// 	Data: struct {
// 		Name string `json:"name" bson:"name"`
// 	}{"Kasra"},
// })
// res, _ := helper.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
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

/* -------------------------------------------------------------------------- */
/*                                    Find                                    */
/* -------------------------------------------------------------------------- */
// filter := bson.D{{Key: "priority", Value: 0}}
// opts := options.Find()
// opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
// results := mongo.Conn.Find("CEC", "log", filter, opts)
// res, _ := helper.Marshal(results)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *Connection) Find(database, collection string, filter primitive.D, opts ...*options.FindOptions) []bson.M {
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

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Panic(err)
	}

	return results
}
