package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/alikazai/golearn/api"
	"github.com/alikazai/golearn/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri    = "mongodb://localhost:27017"
	dbName   = "hotel-reservation"
	userColl = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbName).Collection(userColl)
	user := types.User{
		FirstName: "James",
		LastName:  "Day",
	}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	fmt.Println(client)

	listenAddr := flag.String("listenAddr", ":7500", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	appv1 := app.Group("/api/v1")

	appv1.Get("/user", api.HandleGetUsers)
	appv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
