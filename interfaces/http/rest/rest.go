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

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/interfaces/http/rest/route"
)

type REST struct {
	config *config.Config
	db     *pgxpool.Pool
	redis  *redis.Client

	Port    string
	Handler http.Handler
}

type Option func(r *REST)

func NewREST() *REST {
	return &REST{
		Port: ":8080",
	}
}

func (r *REST) With(opts ...Option) *REST {
	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *REST) Init() *REST {
	if r.config != nil && r.config.REST != nil {
		config := r.config.REST
		if config.Port != "" {
			r.Port = fmt.Sprintf(":%s", config.Port)
		}
	}

	router := route.NewRouter()
	r.Handler = router.Handler(r.db)

	return r
}

func (r *REST) Run() error {
	server := http.Server{
		Addr:         r.Port,
		Handler:      r.Handler,
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
