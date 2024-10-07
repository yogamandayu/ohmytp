package tests

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/internal/db"
	"github.com/yogamandayu/ohmytp/internal/redis"
	"github.com/yogamandayu/ohmytp/internal/slog"
	"github.com/yogamandayu/ohmytp/util"
)

type TestSuite struct {
	App *app.App
}

func NewTestSuite() *TestSuite {
	return &TestSuite{}
}

func (t *TestSuite) LoadApp() {
	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Fatal(err)
	}

	conf := config.NewConfig()
	conf.WithOptions(
		config.WithDBConfig(),
		config.WithRESTConfig(),
		config.WithRedisConfig(),
	)
	dbConn, err := db.NewConnection(conf)
	if err != nil {
		log.Fatal(err)
	}

	redisConn, err := redis.NewConnection(conf)
	if err != nil {
		log.Fatal(err)
	}

	slogger := slog.NewSlog()

	t.App = app.NewApp().WithOptions(app.WithDB(dbConn), app.WithRedis(redisConn), app.WithSlog(slogger))
}

func (t *TestSuite) Clean() {
	t.App.DB.Close()
	t.App.Redis.Close()
}
