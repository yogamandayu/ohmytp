package tests

import (
	"fmt"
	"log"

	app2 "github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/config"
	"github.com/yogamandayu/ohmytp/pkg/db"
	"github.com/yogamandayu/ohmytp/pkg/redis"
	"github.com/yogamandayu/ohmytp/pkg/slog"

	"github.com/joho/godotenv"
	"github.com/yogamandayu/ohmytp/util"
)

type TestSuite struct {
	App *app2.App
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
	dbConn, err := db.NewConnection(conf.DB.Config)
	if err != nil {
		log.Fatal(err)
	}

	redisConn, err := redis.NewConnection(conf.Redis.Config)
	if err != nil {
		log.Fatal(err)
	}

	slogger := slog.NewSlog()

	t.App = app2.NewApp().WithOptions(app2.WithDB(dbConn), app2.WithRedis(redisConn), app2.WithSlog(slogger), app2.WithDBRepository(dbConn))
}

func (t *TestSuite) Clean() {
	t.App.DB.Close()
	t.App.Redis.Close()
}
