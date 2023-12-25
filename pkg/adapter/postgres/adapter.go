package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type postgresDBAdapter struct {
	Client *sql.DB
}

var (
	onceMutex               = sync.Mutex{}
	onceByConfigName        = make(map[string]*sync.Once)
	onceSessionByConfigName = make(map[string]*postgresDBAdapter)
)

const MAX_RETRY_CONNECT int = 5

func NewPostgresDBAdapter(ctx context.Context, config *ScyllaDBConfig) *postgresDBAdapter {

	if onceSessionByConfigName[config.Name] != nil {
		return onceSessionByConfigName[config.Name]
	}

	onceMutex.Lock()
	defer onceMutex.Unlock()

	var adapter = postgresDBAdapter{}

	if onceByConfigName[config.Name] == nil {
		onceByConfigName[config.Name] = &sync.Once{}
	}

	onceByConfigName[config.Name].Do(func() {
		log.Printf("[%s][%s] Postgres [Connecting]\n", config.Name, config.Host)

		var retryConnect = 1

		for retryConnect > 0 && retryConnect < MAX_RETRY_CONNECT {
			log.Printf("[%s][%s] Postgres [Retry Connect]\n", config.Name, config.Host)
			retryConnect += 1

			dbConfig := fmt.Sprintf(
				"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				config.Host, config.Port, config.UserName, config.Password, config.DBName,
			)

			dbSession, err := sql.Open("postgres", dbConfig)
			if err != nil {
				log.Fatalf("Cannot establish the connection to database %s. Error %+v", dbConfig, err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if err := dbSession.PingContext(ctx); err != nil {
				dbSession.Close()
				log.Fatalf("Cannot ping to database %+v", err)
			}
			defer cancel()

			retryConnect = 1
			adapter.Client = dbSession
			break
		}

		onceSessionByConfigName[config.Name] = &adapter
		log.Printf("[%s][%s] Postgres [Connected]\n", config.Name, config.Host)
	})

	return onceSessionByConfigName[config.Name]
}
