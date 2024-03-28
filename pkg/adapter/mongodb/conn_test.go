package mongodb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestConn(t *testing.T) {
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn := NewMongoDBConn(mongdbConfig)
	defer conn.Close()

	db := conn.GetDatabase(mongdbConfig.DbName, nil)

	coll := db.Collection("ContractType")

	result := coll.FindOne(ctx, bson.M{"_id": 125})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			t.Fatalf("_id 125 not found")
		}
	}

	doc := make(map[string]interface{})

	result.Decode(&doc)

	t.Log(doc)
}

func TestConnSigleton(t *testing.T) {
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

	conn1 := NewMongoDBConn(mongdbConfig)
	conn2 := NewMongoDBConn(mongdbConfig)

	defer conn1.Close()

	t.Logf("conn1 address %v", *conn1)
	t.Logf("conn2 address %v", *conn2)

	if assert.Equal(t, *conn1, *conn2) {
		t.Fatal("2 connection not the same")
	}
}
