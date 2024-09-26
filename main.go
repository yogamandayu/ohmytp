package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/interfaces/http/rest"
	"github.com/yogamandayu/ohmytp/repository/persistence/db"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conf := config.NewConfig()
	conf.With(
		config.WithDBConfig(),
		config.WithRESTConfig(),
	)

	dbConn, err := db.NewConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close(context.Background())

	r := rest.NewREST()
	opts := []rest.Option{
		rest.WithConfig(conf),
		rest.WithDB(dbConn),
	}
	if err := r.With(opts...).Init().Run(); err != nil {
		log.Fatal(err)
	}
}
