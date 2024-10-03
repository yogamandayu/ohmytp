package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/interfaces/rest"
	"github.com/yogamandayu/ohmytp/internal/db"
	"github.com/yogamandayu/ohmytp/internal/redis"
	"github.com/yogamandayu/ohmytp/internal/slog"
	"log"
)

type Command struct {
	conf config.Config
}

func NewCommand(conf config.Config) *Command {
	return &Command{
		conf: conf,
	}
}

func (cmd *Command) Commands() cli.Commands {
	return []*cli.Command{
		{
			Name:    "http:rest",
			Aliases: []string{"r"},
			Usage:   "Run REST API",
			Action: func(cCtx *cli.Context) error {
				dbConn, err := db.NewConnection(cmd.conf)
				if err != nil {
					log.Fatal(err)
				}
				defer dbConn.Close()

				redisConn, err := redis.NewConnection(cmd.conf)
				if err != nil {
					log.Fatal(err)
				}
				defer redisConn.Close()

				slogger := slog.NewSlog()

				a := app.NewApp().WithOptions(app.WithDB(dbConn), app.WithRedis(redisConn), app.WithSlog(slogger))

				r := rest.NewREST(a)
				opts := []rest.Option{
					rest.WithConfig(cmd.conf),
				}
				if err := r.With(opts...).Run(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "db:migrate",
			Aliases: []string{"m"},
			Usage:   "Run database migration with tern",
			Action: func(cCtx *cli.Context) error {
				return nil
			},
		},
	}
}
