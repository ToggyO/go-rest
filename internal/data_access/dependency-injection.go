package data_access

import (
	"fmt"
	"go-rest/internal/application/contracts"
	"go-rest/internal/config"
	data_access "go-rest/internal/data_access/common"
	db_context "go-rest/internal/data_access/db-context"
	"go-rest/internal/data_access/repositories"
	"go-rest/internal/shared/models/di"
	settings "go-rest/internal/shared/models/settings/db"
	"go.uber.org/dig"
)

func BindDataAccess() []di.ServiceDescriptor {
	return []di.ServiceDescriptor{
		{Service: bindDbSettings},
		{Service: bindDbContext},
		{Service: data_access.NewMigrationRunner},
		{Service: repositories.NewUsersRepository},
	}
}

func AfterBuild(c *dig.Container) {
	c.Invoke(func(runner data_access.IMigrationRunner, logger contracts.ILogger) {
		if err := runner.MigrateUp(); err != nil {
			logger.Fatal(fmt.Sprintf("Error occured during migration process: %s", err))
		}
		logger.Info("Migrations successfully applied!")
	})
}

func bindDbSettings(configuration *config.Configuration) settings.IDbSettings {
	dbPort := configuration.DbPort
	if configuration.IsDev {
		dbPort = configuration.DbExternalPort
	}

	return &settings.PgDbSettings{
		DbSettings: settings.DbSettings{
			DbHost:     configuration.DbHost,
			DbPort:     dbPort,
			DbUser:     configuration.DbUser,
			DbPassword: configuration.DbPassword,
			DbName:     configuration.DbName,
		},
	}
}

func bindDbContext(s settings.IDbSettings) *db_context.AppDbContext {
	c, _ := db_context.NewDbContext(s.GetConnectionString())
	return c
}
