package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/interfaces/rest/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// REST is a http rest api struct.
type REST struct {
	app *app.App

	Port    string
	Handler http.Handler
}

// NewREST is a constructor.
func NewREST(app *app.App) *REST {
	return &REST{
		app:  app,
		Port: ":8080",
	}
}

// With is to set option.
func (r *REST) With(opts ...Option) *REST {
	for _, opt := range opts {
		opt(r)
	}

	return r
}

// Run is to run http rest api service.
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
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen %s error: %v", r.Port, err)
	}
	log.Println("Shutting down HTTP server gracefully...")

	return nil
}
