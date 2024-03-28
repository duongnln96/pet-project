package mongodb

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDatabase(t *testing.T) {
	mongdbConfig := MongoDBConfig{
		Name:        "hrm",
		Hosts:       []string{"mongo-online.ghn.dev:27017"},
		RSName:      nil,
		AuthSource:  "admin",
		UserName:    "dba",
		Password:    "sFv6UqUnKZSyez7cjTnXvk9WSdz5bWTs",
		Timeout:     30000,
		IsSSLEnable: false,
		PoolLimit:   nil,
		ReadPref:    "",
		DbName:      "hrm",
	}

	const (
		collectionName string = "ContractType"
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := NewDatabase(mongdbConfig, nil)
	if err != nil {
		t.Fatalf("Can not create new database instance %s", err.Error())
	}
	defer db.Close()

	coll := db.GetCollection(collectionName, nil)

	result := coll.FindOne(ctx, bson.M{"_id": 125})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			t.Fatalf("_id 125 not found")
		}
	}

	doc := make(map[string]interface{})

	err = result.Decode(&doc)
	if err != nil {
		t.Fatalf("Can not decode document %s", err.Error())
	}

	t.Log(doc)
}
