package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/interfaces/http/rest"
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
	r := rest.NewREST()
	opts := []rest.Option{
		rest.WithConfig(conf),
	}
	if err := r.With(opts...).Init().Run(); err != nil {
		log.Fatal(err)
	}
}
