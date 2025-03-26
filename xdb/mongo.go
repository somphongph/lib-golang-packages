package xdb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Connection string
	Options    string
	DbName     string
	IsDebug    bool
}

func NewMongoClient(ctx context.Context, sect *Mongo) (*mongo.Client, error) {
	connection := fmt.Sprintf("%s/", sect.Connection)

	if sect.Options != "" {
		connection = fmt.Sprintf("%s/?%s", sect.Connection, sect.Options)
	}
	opts := options.Client().
		ApplyURI(connection).
		SetRetryWrites(false)

	if sect.IsDebug {
		cmdMonitor := &event.CommandMonitor{
			Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
				fmt.Printf("MongoDB CMD: %s\n", evt.Command.String())
			},
		}
		opts = opts.SetMonitor(cmdMonitor)
	}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	fmt.Println("MongoDB initialized")

	return client, nil
}
