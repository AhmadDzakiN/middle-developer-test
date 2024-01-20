package cmd

import (
	"fmt"
	"github.com/rs/zerolog"
	"middle-developer-test/internal/app/appcontext"
	"middle-developer-test/internal/app/config"
	"os"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Up the database",
	Long:  `Execute this command to do migration up`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.AppConfig()
		appCtx := appcontext.NewAppContext(cfg)
		logger := config.NewLogger()
		mSource := getMigrationSource()

		doMigrate(appCtx, logger, mSource, appcontext.DBDialectPostgres, migrate.Up)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Down the database",
	Long:  `Execute this command to do migration down`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.AppConfig()
		appCtx := appcontext.NewAppContext(cfg)
		logger := config.NewLogger()
		mSource := getMigrationSource()

		doMigrate(appCtx, logger, mSource, appcontext.DBDialectPostgres, migrate.Down)
	},
}

var migrateNewCmd = &cobra.Command{
	Use:   "migratenew [migration name]",
	Short: "Create new migration file",
	Long:  `Create new migration file on folder migrations/sql with timestamp as prefix`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := config.NewLogger()
		mDir := "migrations/sql/"

		createMigrationFile(logger, mDir, args[0])
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateNewCmd)
}

func getMigrationSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	return source
}

func doMigrate(appCtx *appcontext.AppContext, log zerolog.Logger, mSource migrate.FileMigrationSource, dbDialect string, direction migrate.MigrationDirection) error {
	db, err := appCtx.GetDBInstance(dbDialect)
	if err != nil {
		log.Fatal().Msgf("Error connection to DB | Error: %+v", err)
		return err
	}

	dbObj, _ := db.DB()

	total, err := migrate.Exec(dbObj, dbDialect, mSource, direction)
	if err != nil {
		log.Fatal().Msgf("Error migrate %s DB | Error: %+v", err)
		return err
	}

	log.Info().Msgf("Success migrate DB | Total: %d", total)
	return nil
}

func createMigrationFile(log zerolog.Logger, mDir string, mName string) error {
	var migrationContent = `-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- [your SQL script here]

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

-- [your SQL script here]
`

	filename := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), mName)
	filepath := fmt.Sprintf("%s%s", mDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		log.Err(err).Msgf("Error create migration file")
		return err
	}
	defer f.Close()

	f.WriteString(migrationContent)
	f.Sync()

	log.Info().Msgf("Success create migration file | filepath: %s", filepath)
	return nil
}
