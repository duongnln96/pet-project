package scylladb

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type ScyllaDBAdaterI interface {
	GetSession() gocqlx.Session
	Close()
}

type scyllaDBAdapter struct {
	ss gocqlx.Session
}

var (
	onceMutex               = sync.Mutex{}
	onceByConfigName        = make(map[string]*sync.Once)
	onceSessionByConfigName = make(map[string]*scyllaDBAdapter)
)

const MAX_RETRY_CONNECT int = 5

func NewScyllaDBAdapter(config *ScyllaDBConfig) ScyllaDBAdaterI {

	if onceSessionByConfigName[config.Name] != nil {
		return onceSessionByConfigName[config.Name]
	}

	onceMutex.Lock()
	defer onceMutex.Unlock()

	if onceByConfigName[config.Name] == nil {
		onceByConfigName[config.Name] = &sync.Once{}
	}

	onceByConfigName[config.Name].Do(func() {
		log.Printf("[%s][%s] ScyllaDB [Connecting]\n", config.Name, config.Hosts)

		var adapter = scyllaDBAdapter{}

		// Create gocql cluster.
		clusterConfig := gocql.NewCluster(config.Hosts...)

		retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
			Min:        time.Second,
			Max:        10 * time.Second,
			NumRetries: 5,
		}

		clusterConfig.RetryPolicy = retryPolicy

		// clusterConfig.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
		clusterConfig.Timeout = config.Timeout
		clusterConfig.WriteTimeout = config.Timeout
		clusterConfig.ConnectTimeout = 5 * time.Second
		clusterConfig.NumConns = 10

		if config.PoolLimit != 0 {
			clusterConfig.NumConns = config.PoolLimit
		}

		// create keyspace
		err := adapter.createKeySpace(config, clusterConfig)
		if err != nil {
			log.Panicf("[%s][%s] ScyllaDB create keyspace %s error %s\n", config.Name, config.Hosts, config.Keyspace, err.Error())
		}

		sessionWithKeyspace, err := adapter.connect(config.Keyspace, clusterConfig)
		if err != nil {
			log.Panicf("[%s][%s] ScyllaDB create session with keyspace %s error %s\n", config.Name, config.Hosts, config.Keyspace, err.Error())
		}

		adapter.ss = sessionWithKeyspace
		onceSessionByConfigName[config.Name] = &adapter
		log.Printf("[%s][%s] ScyllaDB [Connected]\n", config.Name, config.Hosts)
	})

	return onceSessionByConfigName[config.Name]
}

func (s *scyllaDBAdapter) createKeySpace(config *ScyllaDBConfig, clusterConfig *gocql.ClusterConfig) error {

	clusterConfig.Keyspace = "system"

	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	session, err := gocqlx.WrapSession(clusterConfig.CreateSession())
	if err != nil {
		log.Panicf("[%s][%s] ScyllaDB create keyspace error %s\n", config.Name, config.Hosts, err.Error())
	}
	defer session.Close()

	var stmt = fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = { 'class': '%s', 'replication_factor': '%s' } AND durable_writes = TRUE;`, config.Keyspace, config.ReplicationClass, config.ReplicationFactor)

	return session.ExecStmt(stmt)
}

func (s *scyllaDBAdapter) connect(keyspace string, clusterConfig *gocql.ClusterConfig) (gocqlx.Session, error) {
	clusterConfig.Keyspace = keyspace
	return gocqlx.WrapSession(clusterConfig.CreateSession())
}

func (s *scyllaDBAdapter) GetSession() gocqlx.Session {
	return s.ss
}

func (s *scyllaDBAdapter) Close() {
	s.ss.Close()
}
