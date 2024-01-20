package options

import (
	"middle-developer-test/internal/app/appcontext"
)

type ServiceOption struct {
	AppOptions
	*appcontext.Repository
}
