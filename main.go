package main

import (
	"log"
	"os"

	"github.com/yogamandayu/ohmytp/internal/config"

	"github.com/urfave/cli/v2"
	"github.com/yogamandayu/ohmytp/internal/cmd"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conf := config.NewConfig()
	conf.WithOptions(
		config.WithDBConfig(),
		config.WithRedisAPIConfig(),
		config.WithRedisWorkerNotificationConfig(),
		config.WithRESTConfig(),
		config.WithTelegramBotConfig(),
	)

	cliApp := cli.NewApp()
	commands := cmd.NewCommand(conf).Commands()
	cliApp.Commands = commands
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
