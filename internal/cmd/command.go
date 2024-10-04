package cmd

import (
	"fmt"
	"github.com/jackc/tern/v2/migrate"
	"github.com/urfave/cli/v2"
	"github.com/yogamandayu/ohmytp/app"
	"github.com/yogamandayu/ohmytp/config"
	"github.com/yogamandayu/ohmytp/interfaces/rest"
	"github.com/yogamandayu/ohmytp/internal/db"
	"github.com/yogamandayu/ohmytp/internal/redis"
	"github.com/yogamandayu/ohmytp/internal/slog"
	"github.com/yogamandayu/ohmytp/util"
	"log"
	"os"
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

				dbConn, err := db.NewConnection(cmd.conf)
				if err != nil {
					log.Fatal(err)
				}
				defer dbConn.Close()

				migrationsDir := os.DirFS(fmt.Sprintf("%s/domain/migrations", util.RootDir()))

				pgConn, err := dbConn.Acquire(cCtx.Context)
				if err != nil {
					return err
				}
				defer pgConn.Release()

				migrator, err := migrate.NewMigrator(cCtx.Context, pgConn.Conn(), "schema_version")
				if err != nil {
					log.Fatalf("Unable to create migrator: %v\n", err)
				}

				// Load migrations from the specified directory
				err = migrator.LoadMigrations(migrationsDir)
				if err != nil {
					log.Fatalf("Unable to load migrations: %v\n", err)
				}

				// Apply the migrations (Up)
				err = migrator.Migrate(cCtx.Context)
				if err != nil {
					log.Fatalf("Migration failed: %v\n", err)
				}

				log.Println("Migrations applied successfully!")

				return nil
			},
		},
		{
			Name:    "db:gen-repo",
			Aliases: []string{"g"},
			Usage:   "Run generate repo with sqlc",
			Action: func(cCtx *cli.Context) error {
				dbConn, err := db.NewConnection(cmd.conf)
				if err != nil {
					log.Fatal(err)
				}
				defer dbConn.Close()

				migrationsDir := os.DirFS(fmt.Sprintf("%s/domain/migrations", util.RootDir()))

				pgConn, err := dbConn.Acquire(cCtx.Context)
				if err != nil {
					return err
				}
				defer pgConn.Release()

				migrator, err := migrate.NewMigrator(cCtx.Context, pgConn.Conn(), "schema_version")
				if err != nil {
					log.Fatalf("Unable to create migrator: %v\n", err)
				}

				// Load migrations from the specified directory
				err = migrator.LoadMigrations(migrationsDir)
				if err != nil {
					log.Fatalf("Unable to load migrations: %v\n", err)
				}

				// Apply the migrations (Up)
				err = migrator.Migrate(cCtx.Context)
				if err != nil {
					log.Fatalf("Migration failed: %v\n", err)
				}

				log.Println("Repository generated!")

				return nil
			},
		},
	}
}
