package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sebkraemer/go-wishlist/pkg/handlers"
)

const portNumber = ":8080"

// wishlistItem internal resource
type wishlistItem struct {
	// what about _id handling?
	// ID          primitive.ObjectID `bson:"_id, omitempty"` // will create 000000 ids if not set
	Owner       string    `bson:"owner"`
	Description string    `bson:"desc"`
	Date        time.Time `bson:"date"`
}

// Wish Resource type for REST API
type Wish struct {
	Description string `json:"desc"  bson:"desc"`
}

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.16.44.133:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var a = 010
	// pa := &a
	// log.Fatal(*pa)

	// collection := client.Database("wishlist").Collection("wishes")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	// testItem := wishlistItem{Owner: "seb", Description: "test wish", Date: time.Now()}
	// res, err := collection.InsertOne(context.Background(), testItem)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Inserted %v with id %v\n", testItem, res.InsertedID)
}
