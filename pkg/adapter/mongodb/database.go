package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var _ DatabaseI = (*mongodbDatabase)(nil)

type mongodbDatabase struct {
	conn     MongoDBConnI
	database *mongo.Database
}

func NewDatabase(dbConfig MongoDBConfig, opts *DatabaseOptions) (*mongodbDatabase, error) {
	conn := NewMongoDBConn(dbConfig)

	var dbopts = NewDatabaseOptions()
	if opts != nil {
		dbopts = opts
	}

	mongodb := conn.GetDatabase(dbConfig.DbName, dbopts._rawDatabaseOptions())

	return &mongodbDatabase{
		conn:     conn,
		database: mongodb,
	}, nil
}

func (m *mongodbDatabase) Close() error {
	return m.conn.Close()
}

func (m *mongodbDatabase) GetConn() MongoDBConnI {
	return m.conn
}

func (m *mongodbDatabase) GetRawMongoDatabase() *mongo.Database {
	return m.database
}

func (m *mongodbDatabase) GetCollection(collName string, opts *CollectionOptions) *Collection {

	var collopts = NewCollectionOptions()
	if opts != nil {
		collopts = opts
	}

	collection := m.database.Collection(collName, collopts._rawCollectionOptions())

	return &Collection{
		collection,
	}
}
