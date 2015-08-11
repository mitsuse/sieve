package commands

import (
	"github.com/codegangsta/cli"
)

func NewBuildCommand() cli.Command {
	command := cli.Command{
		Name:      "build",
		ShortName: "b",
		Usage:     "Builds a new model for text-filtering",
		Action:    actionBuild,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "model,m",
				Value: "model.still",
				Usage: "The output path of a built model file",
			},

			cli.StringFlag{
				Name:  "examples,e",
				Value: "examples.json",
				Usage: "The input path of a file containing learning examples",
			},

			cli.IntFlag{
				Name:  "iterations,i",
				Value: 10,
				Usage: "The number of iterations for learning.",
			},
		},
	}

	return command
}

func actionBuild(context *cli.Context) {
	// TODO:
}
