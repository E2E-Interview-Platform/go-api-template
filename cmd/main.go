package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/api"
	"github.com/Suhaan-Bhandary/go-api-template/internal/db"
	"github.com/Suhaan-Bhandary/go-api-template/internal/job"
	customcontext "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
	"github.com/go-co-op/gocron/v2"
)

var (
	// Migration Flags
	migrationType           = flag.String("migration", "", "run the migration action based on type, if empty skip migration action. Allowed Values: ['', run, up, down, force, create]")
	migrationForceVersion   = flag.Int("migration-force-version", 0, "paired with --migration=force to provide force version")
	migrationCreateFileName = flag.String("migration-filename", "", "paired with --migration=create to provide filename generating migration files")
)

func main() {
	// Context for main function
	ctx := context.Background()
	ctx = customcontext.SetRequestID(ctx, "main-function")

	// Parsing terminal flags
	flag.Parse()

	// Loading environment variables
	err := environment.LoadEnvironment()
	if err != nil {
		ctxlogger.Error(ctx, err.Error())
		return
	}

	// Running migration if migrationType flag is set
	if *migrationType != "" {
		handleMigration(*migrationType)
		return
	}

	// Connecting to DB
	db, err := repository.InitializeDatabase(ctx)
	if err != nil {
		err = fmt.Errorf("error initializing database: %s", err)
		ctxlogger.Error(ctx, err.Error())
		return
	}
	defer db.Close()

	// Initializing Cron Job
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		ctxlogger.Error(ctx, "scheduler creation failed with error: %s", err.Error())
		return
	}

	job.InitializeJobs(scheduler)
	defer scheduler.Shutdown()

	// Setting chi router and serving it
	apiRouter := api.NewRouter()

	serverAddr := fmt.Sprintf(":%d", environment.PORT)
	ctxlogger.Info(ctx, "Starting server at %s", serverAddr)

	err = http.ListenAndServe(serverAddr, apiRouter)
	if err != nil {
		ctxlogger.Error(ctx, err.Error())
	}
}

func handleMigration(migrationType string) {
	// Context for migration
	ctx := context.Background()
	ctx = customcontext.SetRequestID(ctx, "migration")

	migrations := db.InitializeDBMigrations(ctx)
	switch migrationType {
	case db.RUN:
		migrations.RunMigrations(ctx)
	case db.UP:
		migrations.MigrationsUp(ctx)
	case db.DOWN:
		migrations.MigrationsDown(ctx)
	case db.FORCE:
		forceVersion := *migrationForceVersion
		if forceVersion <= 0 {
			ctxlogger.Error(ctx, "invalid version found to force migration")
			return
		}
		migrations.ForceVersion(ctx, *migrationForceVersion)
	case db.CREATE:
		migrations.CreateMigrationFile(ctx, *migrationCreateFileName)
	default:
		ctxlogger.Warn(ctx, "migration type not found")
	}
}
