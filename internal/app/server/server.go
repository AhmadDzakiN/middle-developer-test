package server

import (
	"context"
	"fmt"
	"middle-developer-test/internal/app/appcontext"
	"middle-developer-test/internal/app/options"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

type IServer interface {
	StartApp()
}

type server struct {
	opt      options.AppOptions
	services *appcontext.Service
}

func NewServer(opt options.AppOptions, services *appcontext.Service) IServer {
	return &server{
		opt:      opt,
		services: services,
	}
}

func (s *server) StartApp() {
	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Info().Msg("[API] Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Info().Msgf("[API] Fail to shutting down: %v", err)
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.opt.AppConfig.GetString("APP_HOST"), s.opt.AppConfig.GetInt("APP_PORT"))
	hOpt := options.HandlerOption{
		AppOptions: s.opt,
		Services:   s.services,
	}

	srv.Handler = Router(hOpt)

	log.Info().Msgf("[API] HTTP serve at %s", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Info().Msgf("[API] Fail to start listen and server: %v", err)
	}

	<-idleConnectionClosed
	log.Info().Msg("[API] Bye!")
}
