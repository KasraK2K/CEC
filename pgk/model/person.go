package model

type Person struct {
	Title    string      `json:"title" bson:"title"`
	Message  string      `json:"message" bson:"message,omitempty"`
	Priority int         `json:"priority" bson:"priority"`
	Data     interface{} `json:"data" bson:"data"`
}

// func (p *Person) Find(database, collection string, filter primitive.D, opts *options.FindOptions, isNeedMarshal ...bool) ([]Person, []string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cursor := mongo.Conn.Find(database, collection, filter, opts)
// 	defer cursor.Close(ctx)

// 	var results []Person
// 	if err := cursor.All(context.TODO(), &results); err != nil {
// 		message := fmt.Sprintf("error FindOne::cursor.All database: `%s` collection: `%s`", database, collection)
// 		log.Panic(message, err)
// 	}

// 	if len(isNeedMarshal) > 0 && isNeedMarshal[0] {
// 		var jsonPayload []string
// 		for _, result := range results {
// 			res, _ := helper.Marshal(result)
// 			jsonPayload = append(jsonPayload, string(res))
// 		}
// 		return results, jsonPayload
// 	}

// 	return results, []string{}
// }
