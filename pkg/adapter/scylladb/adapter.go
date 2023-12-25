package scylladb

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type scyllaDBAdapter struct {
	Client *gocqlx.Session
}

var (
	onceMutex               = sync.Mutex{}
	onceByConfigName        = make(map[string]*sync.Once)
	onceSessionByConfigName = make(map[string]*scyllaDBAdapter)
)

const MAX_RETRY_CONNECT int = 5

func NewScyllaDBAdapter(ctx context.Context, config *ScyllaDBConfig) *scyllaDBAdapter {

	if onceSessionByConfigName[config.Name] != nil {
		return onceSessionByConfigName[config.Name]
	}

	onceMutex.Lock()
	defer onceMutex.Unlock()

	var adapter = scyllaDBAdapter{}

	if onceByConfigName[config.Name] == nil {
		onceByConfigName[config.Name] = &sync.Once{}
	}

	onceByConfigName[config.Name].Do(func() {
		log.Printf("[%s][%s] ScyllaDB [Connecting]\n", config.Name, config.Hosts)

		var retryConnect = 1

		for retryConnect > 0 && retryConnect < MAX_RETRY_CONNECT {
			log.Printf("[%s][%s] ScyllaDB [Retry Connect]\n", config.Name, config.Hosts)
			retryConnect += 1

			// Create gocql cluster.
			clusterConfig := gocql.NewCluster(config.Hosts...)

			clusterConfig.Keyspace = config.Keyspace
			clusterConfig.Timeout = config.Timeout
			clusterConfig.WriteTimeout = config.Timeout
			clusterConfig.ConnectTimeout = 5 * time.Second
			clusterConfig.NumConns = 10

			if config.PoolLimit != 0 {
				clusterConfig.NumConns = config.PoolLimit
			}

			// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
			session, err := gocqlx.WrapSession(clusterConfig.CreateSession())
			if err != nil {
				log.Printf("[%s][%s] ScyllaDB create session error %s\n", config.Name, config.Hosts, err.Error())
				session.Close()
				time.Sleep(10 * time.Millisecond)
				continue
			}

			retryConnect = 1
			adapter.Client = &session
			break
		}

		onceSessionByConfigName[config.Name] = &adapter
		log.Printf("[%s][%s] ScyllaDB [Connected]\n", config.Name, config.Hosts)
	})

	return onceSessionByConfigName[config.Name]
}
