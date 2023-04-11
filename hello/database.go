package main

//import the necessary packages
import (
	"fmt"
	"context"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

// Define a struct to represent the MongoDB document
type Person struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

// Create a connection to the MongoDB server
func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
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

func Run(db *mongo.collection.Collection) error  {

	// Create
	func CreatePerson(person Person, collection *mongo.Collection) error {
		// client, err := connect()
		// if err != nil {
		// 	return err
		// }
		// collection := client.Database("myDB").Collection("people")
		_, err := collection.InsertOne(context.Background(), person)
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
	func UpdatePerson(id string, person Person, collection *mongo.Collection) error {
		// client, err := connect()
		// if err != nil {
		// 	return err
		// }
		// collection := client.Database("myDB").Collection("people")
		filter := bson.M{"_id": id}
		update := bson.M{"$set": person}
		_, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return err
		}
		return nil
	}

	// Delete
	func DeletePerson(id string, collection *mongo.Collection) error {
		// client, err := connect()
		// if err != nil {
		// 	return err
		// }
		defer client.Disconnect(context.Background())

		// collection := client.Database("myDB").Collection("people")
		filter := bson.M{"_id": id}
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return err
		}
		return nil
	}

	// LoggingService
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log.Info().Msg("hello world")
}

func main() {
	// Load the environment variables
    dbHost := os.Getenv("db.host")
    dbPort := os.Getenv("db.port")
    dbName := os.Getenv("db.database")
    dbCollection := os.Getenv("db.collection")

    // Check if any of the environment variables are missing
    if dbHost == "" || dbPort == "" || dbName == "" || dbCollection == "" {
        fmt.Println("One or more required environment variables are missing")
        return

	type Config struct {
		DbHost       string
		DbPort       string
		DbName       string
		DbCollection string
	}

	func NewConfig() (*Config, error) {
		dbHost := os.Getenv("db.host")
		dbPort := os.Getenv("db.port")
		dbName := os.Getenv("db.database")
		dbCollection := os.Getenv("db.collection")
	
		if dbHost == "" || dbPort == "" || dbName == "" || dbCollection == "" {
			return nil, fmt.Errorf("One or more required environment variables are missing")
		}
	
		return &Config{
			DbHost:       dbHost,
			DbPort:       dbPort,
			DbName:       dbName,
			DbCollection: dbCollection,
		}, nil
	}
	
	// Load config object
	config, err := config.NewConfig.env()
	if err != nil { 
		log.Error().Err(err).Msg(err.Error()) }
	// New database connection
	db, err := NewDBConnection(config)
	
	// Run 
	if err := Run(db); err != nil { panic(err) }
	}
	}