package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/interfaces/http/rest/route"
)

type REST struct {
	config *config.Config
	app    *app.App

	Port    string
	Handler http.Handler
}

func NewREST(app *app.App) *REST {
	return &REST{
		app:  app,
		Port: ":8080",
	}
}

func (r *REST) With(opts ...Option) *REST {
	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *REST) Run() error {
	router := route.NewRouter(r.app)
	server := http.Server{
		Addr:         fmt.Sprintf(":%s", r.Port),
		Handler:      router.Handler(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
		log.Println("Shutdown complete")
	}()

	log.Println("HTTP server is starting ...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen %s error: %v", r.Port, err)
	}
	log.Println("Shutting down HTTP server gracefully...")

	return nil
}
