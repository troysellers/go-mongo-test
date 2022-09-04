package mongoUtil

import (
	"context"
	"fmt"
	"log"

	"github.com/troysellers/mongotest/imdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMovieByTitle(c *mongo.Client, title string) bson.M {
	coll := c.Database("sample_mflix").Collection("movies")

	cursor, err := coll.Find(context.TODO(), bson.M{"title": title})
	if err != nil {
		fmt.Printf("%v", err)
	}

	var results bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(results)
	for _, result := range results {
		fmt.Print(result)
	}
	return results
}

func MovieExists(c *mongo.Client, title string) (primitive.ObjectID, error) {
	coll := c.Database("sample_mflix").Collection("movies")
	cursor, err := coll.Find(context.TODO(), bson.M{"title": title})
	if err != nil {
		return primitive.NilObjectID, err
	}
	defer cursor.Close(context.TODO())
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return primitive.NilObjectID, err
	}

	if len(results) == 0 {
		return primitive.NilObjectID, nil
	}

	return results[0]["_id"].(primitive.ObjectID), nil
}

func InsertMovie(c *mongo.Client, m *imdb.Movie) error {
	coll := c.Database("sample_mflix").Collection("movies")
	doc := bson.D{
		{"title", m.Title},
		{"crew", m.Crew},
		{"description", m.Description},
		{"full_title", m.FullTitle},
		{"imdb_id", m.Id},
		{"image", m.Image},
		{"imdb_rating", m.ImdbRating},
		{"imdb_rating_count", m.ImdbRatingCount},
		{"rank", m.Rank},
		{"year", m.Year},
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}

func UpdateMovie(c *mongo.Client, m *imdb.Movie, id primitive.ObjectID) error {

	coll := c.Database("sample_mflix").Collection("movies")
	filter := bson.D{
		{"_id", id},
	}
	update := bson.D{
		{"$set",
			bson.D{
				{"crew", m.Crew},
				{"description", m.Description},
				{"full_title", m.FullTitle},
				{"imdb_id", m.Id},
				{"image", m.Image},
				{"imdb_rating", m.ImdbRating},
				{"imdb_rating_count", m.ImdbRatingCount},
				{"rank", m.Rank},
				{"year", m.Year},
			},
		}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	log.Println(result)
	return nil
}
