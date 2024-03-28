package mongodb

import (
	"io"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnI interface {
	GetDatabase(dbName string, opts *options.DatabaseOptions) *mongo.Database
	GetRawConn() *mongo.Client
	Close() error
}

type DatabaseI interface {
	io.Closer
	GetConn() MongoDBConnI
	GetCollection(collName string, opts *CollectionOptions) *Collection
}

type Collection struct {
	*mongo.Collection
}
