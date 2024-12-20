package tests

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/config"
	"github.com/yogamandayu/ohmytp/pkg/db"
	"github.com/yogamandayu/ohmytp/pkg/redis"
	"github.com/yogamandayu/ohmytp/pkg/rollbar"
	"github.com/yogamandayu/ohmytp/pkg/slog"
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
		config.WithRedisAPIConfig(),
		config.WithRedisWorkerNotificationConfig(),
		config.WithRESTConfig(),
		config.WithTelegramBotConfig(),
		config.WithRollbarConfig(),
	)

	dbConn, err := db.NewConnection(conf.DB.Config)
	if err != nil {
		log.Fatal(err)
	}

	redisAPIConn, err := redis.NewConnection(conf.RedisAPI.Config)
	if err != nil {
		log.Fatal(err)
	}

	redisNotificationConn, err := redis.NewConnection(conf.RedisAPI.Config)
	if err != nil {
		log.Fatal(err)
	}

	slogger := slog.NewSlog()

	rollbarClient := rollbar.NewRollbar(conf.Rollbar.Config)
	if util.GetEnvAsBool("ENABLE_ROLLBAR", false) {
		rollbarClient.SetEnabled(true)
	}

	t.App = app.NewApp().WithOptions(
		app.WithConfig(conf),
		app.WithDB(dbConn),
		app.WithRedisAPI(redisAPIConn),
		app.WithRedisWorkerNotification(redisNotificationConn),
		app.WithSlog(slogger),
		app.WithRollbar(rollbarClient),
	)
}

func (t *TestSuite) Clean() {
	t.App.DB.Close()
	t.App.RedisAPI.Close()
	t.App.RedisWorkerNotification.Close()
	t.App.Rollbar.Close()
}
