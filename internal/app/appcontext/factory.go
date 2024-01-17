package appcontext

import (
	"errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"middle-developer-test/internal/app/config"
)

const (
	DBDialectPostgres = "postgres"
)

type AppContext struct {
	config *viper.Viper
}

func NewAppContext(config *viper.Viper) *AppContext {
	return &AppContext{
		config: config,
	}
}

func (a *AppContext) GetDBInstance(dbType string) (*gorm.DB, error) {
	var gormDbMap *gorm.DB
	var err error
	switch dbType {
	case DBDialectPostgres:
		dbOption := a.getPostgreOption()
		gormDbMap, err = config.NewPostgreDatabase(dbOption)
	default:
		err = errors.New("Error get db instance, unknown db type")
	}

	return gormDbMap, err
}

func (a *AppContext) getPostgreOption() config.DBPostgreOption {
	return config.DBPostgreOption{
		Host:            a.config.GetString("POSTGRE_HOST"),
		Port:            a.config.GetInt("POSTGRE_PORT"),
		Username:        a.config.GetString("POSTGRE_USERNAME"),
		Password:        a.config.GetString("POSTGRE_PASSWORD"),
		DBName:          a.config.GetString("POSTGRE_NAME"),
		MaxPoolSize:     a.config.GetInt("POSTGRE_POOL_SIZE"),
		MaxIdleConns:    a.config.GetInt("POSTGRE_MAX_IDLE_CONNS"),
		ConnMaxLifetime: a.config.GetDuration("POSTGRE_MAX_CONN_LIFETIME"),
	}
}
