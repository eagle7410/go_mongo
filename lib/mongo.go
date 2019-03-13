package lib

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	MongoClient struct {
		Host *string
		Port *string
		Client *mongo.Client
	}
)

var Mongo = MongoClient{}

func (i *MongoClient) Init () error  {
	i.Host = &ENV.HOST
	i.Port = &ENV.PORT

	credential := options.Credential{
		Username:ENV.AUTH_USER,
		Password:ENV.AUTH_PASS,
	}

	connectLink := fmt.Sprintf("mongodb://%v:%v", *i.Host, *i.Port)

	opt := options.Client().ApplyURI(connectLink).SetAuth(credential)

	var err error = nil

	if i.Client, err = mongo.Connect(context.TODO(), opt); err != nil {
		return err;
	}


	return i.Client.Ping(context.TODO(), nil)
}
