package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/interfaces/rest"
	"github.com/yogamandayu/ohmytp/internal/db"
	"github.com/yogamandayu/ohmytp/internal/redis"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conf := config.NewConfig()
	conf.With(
		config.WithDBConfig(),
		config.WithRESTConfig(),
		config.WithRedisConfig(),
	)

	dbConn, err := db.NewConnection(*conf)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	redisConn, err := redis.NewConnection(*conf)
	if err != nil {
		log.Fatal(err)
	}
	defer redisConn.Close()

	app := app.NewApp().WithOptions(app.WithDB(dbConn), app.WithRedis(redisConn))

	r := rest.NewREST(app)
	opts := []rest.Option{
		rest.WithConfig(conf),
	}
	if err := r.With(opts...).Run(); err != nil {
		log.Fatal(err)
	}
}
