package options

import "middle-developer-test/internal/app/appcontext"

type HandlerOption struct {
	AppOptions
	Services *appcontext.Service
}
