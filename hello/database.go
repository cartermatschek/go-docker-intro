package main

//import the necessary packages
import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to represent the MongoDB document
type Person struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

// Create a connection to the MongoDB server
func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Create
func CreatePerson(person Person) error {
	client, err := connect()
	if err != nil {
		return err
	}
	collection := client.Database("myDB").Collection("people")
	_, err = collection.InsertOne(context.Background(), person)
	if err != nil {
		return err
	}
	return nil
}

// Read
func ReadPerson(id string) (Person, error) {
	client, err := connect()
	if err != nil {
		return Person{}, err
	}
	collection := client.Database("myDB").Collection("people")
	filter := bson.M{"_id": id}
	var person Person
	err = collection.FindOne(context.Background(), filter).Decode(&person)
	if err != nil {
		return Person{}, err
	}
	return person, nil
}

// Update
func UpdatePerson(id string, person Person) error {
	client, err := connect()
	if err != nil {
		return err
	}
	collection := client.Database("myDB").Collection("people")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": person}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func DeletePerson(id string) error {
	client, err := connect()
	if err != nil {
		return err
	}
	collection := client.Database("myDB").Collection("people")
	filter := bson.M{"_id": id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
