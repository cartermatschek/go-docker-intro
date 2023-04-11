package main

import "testing"

func TestDatabase(t *testing.T) {

	client, err := connect()
	collection := client.Database("myDB").Collection("people")

	// Create
	if err := CreatePerson(Person{ID: "1", Name: "John", Age: 20}, collection); err != nil {
		t.Error(err)
	}

	// Read
	person, err := ReadPerson("1")
	if err != nil {
		t.Error(err)
	}
	if person.Name != "John" {
		t.Error("Name should be John")
	}

	// Update
	if err := UpdatePerson("1", Person{ID: "1", Name: "Jane", Age: 20}, collection); err != nil {
		t.Error(err)
	}

	// Delete
	if err := DeletePerson("1", collection); err != nil {
		t.Error(err)
	}
}
