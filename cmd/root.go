package cmd

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"middle-developer-test/internal/app/appcontext"
	"middle-developer-test/internal/app/config"
	"middle-developer-test/internal/app/domain/employee"
	"middle-developer-test/internal/app/options"
	"middle-developer-test/internal/app/server"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "middle-developer-test",
	Short: "Start the API for Middle Developer Test",
	Run: func(cmd *cobra.Command, args []string) {
		startAPI()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

// API bootstrap
func startAPI() {
	var err error
	cfg := config.AppConfig()
	app := appcontext.NewAppContext(cfg)
	config.NewLogger()

	var dbPostgre *gorm.DB
	if cfg.GetBool("POSTGRE_IS_ENABLE") {
		dbPostgre, err = app.GetDBInstance(appcontext.DBDialectPostgres)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start, error connect to DB Postgre")
			return
		}
	}

	opt := options.AppOptions{
		AppCtx:    app,
		AppConfig: cfg,
		DbPostgre: dbPostgre,
	}

	repo := wiringRepository(options.RepositoryOption{
		AppOptions: opt,
	})

	services := wiringService(options.ServiceOption{
		AppOptions: opt,
		Repository: repo,
	})

	newServer := server.NewServer(opt, services)

	// Run the API server
	newServer.StartApp()
}

func wiringRepository(repoOption options.RepositoryOption) (repo *appcontext.Repository) {
	repo = &appcontext.Repository{
		Employee: employee.NewRepository(repoOption),
	}

	return
}

func wiringService(svcOption options.ServiceOption) (service *appcontext.Service) {
	service = &appcontext.Service{
		Employee: employee.NewService(svcOption),
	}

	return
}
