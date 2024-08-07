package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// define 'Trainer' struct
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//  get a handle for the 'trainers' collection in the 'demo' database
	collection := client.Database("demo").Collection("trainers")

	// define some Trainer instances
	hermon := Trainer{"Hermon", 21, "Addis Ababa"}
	beamlak := Trainer{"Beamlak", 24, "Chicago"}
	fenu := Trainer{"Fenu", 15, "Seoul"}

	//---------------------------------------------------------------

	//--------------------INSERT ONE DOCUMENT-------------------
	insertResult, err := collection.InsertOne(context.TODO(), hermon)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single documeent: ", insertResult.InsertedID)

	// ------------------INSERT MULTIPLE DOCUMENTS--------------
	trainers := []interface{}{beamlak, fenu}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//--------------------UPDATE DOCUMENTS--------------------
	filter := bson.D{{"name", "Ash"}}

	// increment age by 1
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents. \n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//--------------------FIND A SINGLE DOCUMENT--------------------
	var result Trainer
	err = collection.FindOne(context.TODO(), filter).Decode(&result) // notice that the pointer for result is passed here

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %v\n", result)

	//--------------------FIND MULTIPLE DOCUMENTS--------------------
	findOptions := options.Find()
	findOptions.SetLimit(2) // limit the number of documents returned to 2

	// cur -> 'Cursor' provides a stream of documents through which you can iterat and decode
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	var results []*Trainer

	for cur.Next(context.TODO()) {
		// decode each element one by one and put into 'results'
		var curr_document Trainer
		err := cur.Decode(&curr_document)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &curr_document)
	}

	//--------------------DELETE ONE DOCUMENT--------------------
	delete_filter := bson.D{{"name", "Hermon"}}

	deleteResult, err := collection.DeleteOne(context.TODO(), delete_filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//--------------------DELETE MULTIPLE DOCUMENTS--------------------
	delete_filter2 := bson.D{{"name", "Beamlak"}}

	deleteResult2, err := collection.DeleteMany(context.TODO(), delete_filter2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult2.DeletedCount)

	//------------------------------------------------------------------

	// It is best practice to keep a client that is connected to
	// MongoDB around so that the application can make use of
	// connection pooling
	// However, if your application no longer requires a connection,
	// you can close the connection with the following script
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
