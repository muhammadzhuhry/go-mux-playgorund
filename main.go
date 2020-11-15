package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lasttname,omitempty"`
}

func main() {
	fmt.Println("Starting the application...")
}
