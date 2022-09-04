package mongoUtil

import (
	"context"
	"fmt"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	os.Setenv("MONGO_URL", "mongodb+srv://troy-mongo-user:thisMongoTest@troys-mongo-test.jy8ejnj.mongodb.net/?retryWrites=true&w=majority")
}

func TestFindByTitle(t *testing.T) {
	uri := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	t.Log("starting")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Println(err)
		}
	}()
	coll := client.Database("sample_mflix").Collection("movies")
	cur, err := coll.Find(context.TODO(), bson.M{"title": "The Shawshank Redemption"})
	if err != nil {
		t.Logf("%v", err)
	}
	var results []bson.M
	if err := cur.All(context.Background(), &results); err != nil {
		t.Log(err)
	}
	t.Logf("%v", results)
}
