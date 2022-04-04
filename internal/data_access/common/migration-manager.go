package data_access

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"go-rest/internal/application/contracts"
	db_context "go-rest/internal/data_access/db-context"
	"go-rest/internal/data_access/migrations"
)

type IMigrationRunner interface {
	MigrateUp() error
	MigrateDown() error
}

type migrationRunner struct {
	migrator *gormigrate.Gormigrate
	logger   contracts.ILogger
}

func NewMigrationRunner(c *db_context.AppDbContext, l contracts.ILogger) IMigrationRunner {
	return &migrationRunner{
		migrator: gormigrate.New(c.GetConnection(), gormigrate.DefaultOptions, migrations.MigrationCollection),
		logger:   l,
	}
}

func (m migrationRunner) MigrateUp() error {
	m.logger.Info("Applying migrations...")
	return m.migrator.Migrate()
}

func (m migrationRunner) MigrateDown() error {
	m.logger.Info("Reverting migrations...")
	return m.migrator.RollbackLast()
}
