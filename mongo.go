package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/troysellers/mongotest/imdb"
	"github.com/troysellers/mongotest/mongoUtil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {

	var imdbListId string
	flag.StringVar(&imdbListId, "list", "", "Enter the IMDB List ID")
	flag.Parse()
	if imdbListId == "" {
		log.Fatal("Enter the list ID of the IMDB List you wish to populate.")
	}
	client, err := getClient()
	if err != nil {
		log.Println("Could initialise mongo client")
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	movies, err := imdb.GetList(imdbListId)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range movies {

		id, err := mongoUtil.MovieExists(client, m.Title)
		if err != nil {
			log.Fatal(err)
		}
		if id == primitive.NilObjectID {
			mongoUtil.InsertMovie(client, m)
		} else {
			mongoUtil.UpdateMovie(client, m, id)
		}
	}

}

func getClient() (*mongo.Client, error) {
	uri := os.Getenv("MONGO_URL")
	log.Printf("URL %s", uri)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	return client, err
}
