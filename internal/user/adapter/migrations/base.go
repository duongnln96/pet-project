package migrations

import (
	"github.com/duongnln96/blog-realworld/pkg/config"

	psqlAdapter "github.com/duongnln96/blog-realworld/pkg/adapter/postgres"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(configs *config.Configs) error {

	adapter := psqlAdapter.NewPostgresDBAdapter(configs.PostgresConfigMap.Get("postgres"))
	driver, _ := postgres.WithInstance(adapter.GetDB(), &postgres.Config{})

	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	return m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
