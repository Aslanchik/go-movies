package mongodb

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

const DB_NAME = "movie-land"
const MONGO_URI = "mongodb://localhost:27017/ffde015a-5d43-41e2-905b-473fcae8dd29?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"

func ConfigureAndConnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		return err
	}

	err = client.Connect(ctx)
	db := client.Database(DB_NAME)

	if err != nil {
		return err
	}

	Instance = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
