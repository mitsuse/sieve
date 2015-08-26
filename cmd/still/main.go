package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/still/cmd/still/commands"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {
	app := cli.NewApp()

	app.Name = commands.Name
	app.Version = commands.Version
	app.Usage = commands.Description
	app.Author = commands.Author
	app.Email = commands.AuthorEmail

	app.Commands = []cli.Command{
		commands.NewBuildCommand(),
		commands.NewTestCommand(),
		commands.NewFilterCommand(),
	}

	return app
}
