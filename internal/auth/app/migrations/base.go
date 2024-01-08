package migrations

import (
	"context"
	"log/slog"
	"os"

	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/duongnln96/blog-realworld/pkg/logger"

	scylladbAdapter "github.com/duongnln96/blog-realworld/pkg/adapter/scylladb"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"

	"github.com/duongnln96/blog-realworld/db/cql"

	"github.com/sirupsen/logrus"
)

func Run(config *config.Configs) error {

	// set up logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	// integrate Logrus with the slog logger
	slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	scylladbConfigs := config.ScyllaDBConfigMap.Get("scylladb")

	adapter := scylladbAdapter.NewScyllaDBAdapter(scylladbConfigs)
	session := adapter.GetSession()
	defer session.Close()

	callbackFunc := func(ctx context.Context, session gocqlx.Session, ev migrate.CallbackEvent, name string) error {
		slog.Info("callbackFunc", "ev", ev, "name", name)
		return nil
	}

	reg := migrate.CallbackRegister{}
	reg.Add(migrate.BeforeMigration, "token.cql", callbackFunc)
	reg.Add(migrate.AfterMigration, "token.cql", callbackFunc)

	migrate.Callback = reg.Callback

	// First run prints data
	if err := migrate.FromFS(context.Background(), session, cql.Files); err != nil {
		logrus.Panicln("Migrate ", err.Error())
	}

	// Second run skips the processed files
	if err := migrate.FromFS(context.Background(), session, cql.Files); err != nil {
		logrus.Panicln("Migrate ", err.Error())
	}

	return nil
}
