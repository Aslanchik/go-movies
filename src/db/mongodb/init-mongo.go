package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Instance MongoInstance

const dbName = "movie-land"
const mongoURI = "mongodb://localhost:27017/ffde015a-5d43-41e2-905b-473fcae8dd29?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"

func ConfigureAndConnect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	Instance = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
