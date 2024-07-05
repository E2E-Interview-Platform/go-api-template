package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/helpers"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"

	// migrate dependency
	_ "github.com/golang-migrate/migrate/v4/database/mysql"

	// migrate dependency
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	// Defines the folder of migration
	migrationsDIR = "./internal/db/migrations"

	// Defines path for migration files
	migrationFilesPath = "file://" + migrationsDIR

	// Migration Type
	RUN    = "run"
	UP     = "up"
	DOWN   = "down"
	FORCE  = "force"
	CREATE = "create"
)

type Migration struct {
	m             *migrate.Migrate
	directoryName string
	filesPath     string
}

func InitializeDBMigrations(ctx context.Context) Migration {
	dbConnection := fmt.Sprintf(
		"mysql://%v:%v@tcp(%v)/%v",
		url.QueryEscape(environment.DB_USER),
		url.QueryEscape(environment.DB_PASSWORD),
		environment.DB_URL,
		environment.DB_NAME,
	)

	migrations := Migration{}
	migrations.directoryName = migrationsDIR
	migrations.filesPath = migrationFilesPath

	var err error
	migrations.m, err = migrate.New(migrations.filesPath, dbConnection)
	if err != nil {
		if err == migrate.ErrNoChange {
			return migrations
		}

		log.Fatal(fmt.Errorf("error during migration, err: %s", err))
	}

	return migrations
}

func (migration Migration) RunMigrations(ctx context.Context) {
	ctxlogger.Info(ctx, "%s Migrations started", migration.directoryName)

	startTime := time.Now()
	defer func() {
		ctxlogger.Info(ctx, "%s Migrations complete, total time taken  %s", migration.directoryName, time.Since(startTime))
	}()

	// dbVersion is the currently active database migration version
	dbVersion, dirty, err := migration.m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatal(err)
	}

	// localVersion is the local migration version
	localVersion, err := migration.MigrationLocalVersion(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if dbVersion > uint(localVersion) {
		log.Fatalf("Your database migration %d is ahead of local migration %d, you might need to rollback few migrations", dbVersion, localVersion)
	}

	if dbVersion < uint(localVersion) && dirty {
		log.Fatalf("Your currently active database migration %d is dirty, you might need to rollback it and then deploy again.", dbVersion)
	}

	err = migration.m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		dbVersion, _, err2 := migration.m.Version()
		if err2 != nil {
			log.Fatal(err2)
		}

		log.Fatalf("Migration failed with error %s, current active database migration is %d", err, dbVersion)
	}
}

func (migration Migration) MigrationsUp(ctx context.Context) {
	err := migration.m.Steps(1)
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		log.Fatal(err)
	}

	ctxlogger.Info(ctx, "Migration up complete")
}

func (migration Migration) MigrationsDown(ctx context.Context) {
	err := migration.m.Steps(-1)
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		log.Fatal(err)
	}

	ctxlogger.Info(ctx, "Migration down complete")
}

func (migration Migration) ForceVersion(ctx context.Context, version int) {
	err := migration.m.Force(version)
	if err != nil {
		log.Fatal(err)
	}

	ctxlogger.Info(ctx, "Migration force version %v complete\n", version)
}

func (migration Migration) CreateMigrationFile(ctx context.Context, filename string) error {
	var err error

	if len(filename) == 0 {
		return customerrors.Error{
			CustomMessage: "filename is not provided",
			InternalError: err,
		}
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", migration.directoryName, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", migration.directoryName, timeStamp, filename)

	defer func() {
		if err != nil {
			os.Remove(upMigrationFilePath)
			os.Remove(downMigrationFilePath)
		}
	}()

	err = createFile(upMigrationFilePath)
	if err != nil {
		return err
	}

	ctxlogger.Info(ctx, "created %s\n", upMigrationFilePath)

	err = createFile(downMigrationFilePath)
	if err != nil {
		return err
	}

	ctxlogger.Info(ctx, "created %s\n", downMigrationFilePath)
	return nil
}

func (migration Migration) GetMigrationVersion(ctx context.Context) (uint, bool, error) {
	version, dirty, err := migration.m.Version()
	if err != nil {
		return 0, false, err
	}

	return version, dirty, nil
}

func (migration Migration) MigrationLocalVersion(ctx context.Context) (int, error) {
	localDIRFileVersions, err := getMigrationVersionsFromDir(migration.directoryName)
	if err != nil {
		err = fmt.Errorf("error during reading files information from local file system, err: %s", err)
		return 0, err
	}

	//if there are no files present in local file system, then rollback not required
	if len(localDIRFileVersions) == 0 {
		ctxlogger.Warn(ctx, "no migration files found in local file system")
		return 0, err
	}

	ctxlogger.Warn(ctx, "latest migration version from local file system: %d", localDIRFileVersions[0])
	return localDIRFileVersions[0], nil
}

func (migration Migration) GoTo(ctx context.Context, version uint) error {
	localDIRFileVersions, err := getMigrationVersionsFromDir(migration.directoryName)
	if err != nil {
		err = fmt.Errorf("error during reading files information from local file system, err: %s", err)
		return err
	}

	//if there are no files present in local file system, then migration not required
	if len(localDIRFileVersions) == 0 {
		ctxlogger.Warn(ctx, "no migration files found in local file system, hence migration not required")
		return nil
	}

	//get the database version from database
	dbversion, dirty, err := migration.m.Version()
	if err != nil {
		if err != migrate.ErrNilVersion {
			return err
		}

		ctxlogger.Info(ctx, "no migration found, initializing DB with latest migration")
		err = migration.m.Migrate(version)
		if err != migrate.ErrNoChange {
			return err
		}

		ctxlogger.Info(ctx, "database successfully initialized with migration: %d", version)
		return nil
	}

	// if the database is in dirty state, we pick the previous successful executed migration
	// and force the database to that version
	if dirty {
		//get the index of version that we get from the database
		index, err := helpers.GetIndexOfElementInSlice(localDIRFileVersions, int(dbversion))
		if err != nil {
			return errors.New("database version corresponding file not found in local file system")
		}

		if len(localDIRFileVersions) <= index+1 {
			return errors.New("previous successfully executed migration not found in local file system")
		}
		//get the version just before that version
		forceMigrateVersion := localDIRFileVersions[index+1]

		err = migration.m.Force(forceMigrateVersion)
		if err != nil {
			return err
		}
	}

	err = migration.m.Migrate(version)
	if err != migrate.ErrNoChange {
		return err
	}

	ctxlogger.Info(ctx, "database successfully migrated to version: %d", version)
	return nil
}
