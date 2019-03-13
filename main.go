package main

import (
	"fmt"
	"go_mongo/lib"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name string
	Role string
}

func init() {
	if err := lib.ENV.Init(); err != nil {
		log.Fatal(err)
	}

	if err := lib.Mongo.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	u1 := User{
		Name: "test1",
		Role: "role1",
	}

	u2 := User{
		Name: "test2",
		Role: "role2",
	}

	collection := lib.Mongo.Client.Database("goy").Collection("users")

	insertResult, err := collection.InsertOne(context.TODO(), u1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a user1 : ", insertResult.InsertedID)

	insertResult, err = collection.InsertOne(context.TODO(), u2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a user2 : ", insertResult.InsertedID)

	opt := options.Find()

	cur, err := collection.Find(context.TODO(), bson.D{{}}, opt)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var u User

		if err := cur.Decode(&u); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Find user %v \n", u)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	deleteResult, err := collection.DeleteMany(context.TODO(),  bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v user\n", deleteResult.DeletedCount)
}
