package Conn

import (
	"context"
	"time"

	"github.com/mcharriere/mongocheck/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/readpref"

	"fmt"
)

type Conn struct {
	Client  *mongo.Client
	Context context.Context
}

func New(opts *Config.Config) *Conn {
	c := new(Conn)

	client, err := mongo.NewClient(opts.GetOptions())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	c.Client = client
	return c
}

func (cnx *Conn) Check() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cnx.Client.Connect(ctx)

	err := cnx.Client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
