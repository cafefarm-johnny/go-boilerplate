package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var connection *mongoDB

type mongoDB struct {
	client mongo.Client
	db     mongo.Database
}

func init() {
	env := newMongoEnv()

	connect(env)
}

func connect(env *mongoEnv) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := new(options.Credential)
	credential.Username = env.username
	credential.Password = env.password

	opt := options.Client().ApplyURI(env.uri).SetAuth(*credential)

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal("Failed Connect MongoDB Cause: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed Ping MongoDB Cause: ", err)
	}

	connection = new(mongoDB)
	connection.client = *client
	connection.db = *client.Database(env.db)

	log.Printf("Connected MongoDB to %s (DB: %s) \n", env.uri, env.db)
}

func DB() *mongo.Database {
	return &connection.db
}
