package options

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"middle-developer-test/internal/app/appcontext"
)

type AppOptions struct {
	AppCtx    *appcontext.AppContext
	AppConfig *viper.Viper
	DbPostgre *gorm.DB
}
