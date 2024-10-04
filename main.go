package main

import (
	"github.com/urfave/cli/v2"
	"github.com/yogamandayu/ohmytp/internal/cmd"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yogamandayu/ohmytp/config"
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

	cliApp := cli.NewApp()
	commands := cmd.NewCommand(conf).Commands()
	cliApp.Commands = commands
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
