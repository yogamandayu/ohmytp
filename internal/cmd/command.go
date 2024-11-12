package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	app2 "github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/config"
	rest2 "github.com/yogamandayu/ohmytp/internal/interfaces/rest"
	"github.com/yogamandayu/ohmytp/pkg/db"
	"github.com/yogamandayu/ohmytp/pkg/redis"
	"github.com/yogamandayu/ohmytp/pkg/slog"

	"github.com/jackc/tern/v2/migrate"
	"github.com/urfave/cli/v2"
	"github.com/yogamandayu/ohmytp/util"
)

// Command is a run service command.
type Command struct {
	conf config.Config
}

// NewCommand is a constructor.
func NewCommand(conf config.Config) *Command {
	return &Command{
		conf: conf,
	}
}

// Commands is get list of commands.
func (cmd *Command) Commands() cli.Commands {
	return []*cli.Command{
		{
			Name:    "http:rest",
			Aliases: []string{"r"},
			Usage:   "Run REST API",
			Action: func(cCtx *cli.Context) error {
				dbConn, err := db.NewConnection(cmd.conf.DB.Config)
				if err != nil {
					log.Fatal(err)
				}
				defer dbConn.Close()

				redisConn, err := redis.NewConnection(cmd.conf.Redis.Config)
				if err != nil {
					log.Fatal(err)
				}
				defer redisConn.Close()

				slogger := slog.NewSlog()

				a := app2.NewApp().WithOptions(app2.WithDB(dbConn), app2.WithRedis(redisConn), app2.WithSlog(slogger), app2.WithDBRepository(dbConn))

				r := rest2.NewREST(a)
				opts := []rest2.Option{
					rest2.WithConfig(cmd.conf),
				}
				if err := r.With(opts...).Run(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "db:migrate",
			Aliases: []string{"dbm"},
			Usage:   "Run database migration with tern",
			Action: func(cCtx *cli.Context) error {

				dbConn, err := db.NewConnection(cmd.conf.DB.Config)
				if err != nil {
					log.Fatal(err)
				}
				defer dbConn.Close()

				migrationsDir := os.DirFS(fmt.Sprintf("%s/internal/domain/migrations", util.RootDir()))

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
			Name:    "db:generate",
			Aliases: []string{"dbg"},
			Usage:   "Run database migration with tern",
			Action: func(cCtx *cli.Context) error {
				err := exec.Command("cd", util.RootDir()).Run()
				if err != nil {
					log.Fatal(err)
				}
				err = exec.Command("sqlc", "generate").Run()
				if err != nil {
					log.Fatal(err)
				}

				log.Println("Database generate successfully!")

				return nil
			},
		},
		{
			Name:    "git:pre-commit",
			Aliases: []string{"hooks"},
			Usage:   "Install pre-commit",
			Action: func(cCtx *cli.Context) error {
				err := exec.Command("cp", fmt.Sprintf("%s/.githooks/pre-commit", util.RootDir()), fmt.Sprintf("%s/.git/hooks/pre-commit", util.RootDir())).Run()
				if err != nil {
					log.Fatal(err)
				}
				err = exec.Command("chmod", "+x", fmt.Sprintf("%s/.git/hooks/pre-commit", util.RootDir())).Run()
				if err != nil {
					log.Fatal(err)
				}

				log.Println("Pre-commit installed!")
				return nil
			},
		},
	}
}
