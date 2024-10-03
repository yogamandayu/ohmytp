package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
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

	model := cmd.NewBubbleTea()
	model.SetListItemAndArgs(map[string]string{
		"DB Migration": "db:migrate",
		"REST API":     "http:rest",
	})
	model.SetCommand(func(args string) error {
		cmdArgs := []string{os.Args[0], args}
		cliApp := cli.NewApp()
		commands := cmd.NewCommand(conf).Commands()
		cliApp.Commands = commands
		return cliApp.Run(cmdArgs)
	})
	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
