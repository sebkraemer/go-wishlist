package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WishlistItem struct {
	Owner string    `json:"owner" bson:"owner"`
	Desc  string    `json:"desc"  bson:"desc"`
	Date  time.Time `json:"date"  bson:"date"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.16.44.133:27017"))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("wishlist").Collection("wishes")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	testItem := WishlistItem{"seb", "test wish", time.Now()}

	res, err := collection.InsertOne(context.Background(), testItem)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v with id %v\n", testItem, res.InsertedID)
}
