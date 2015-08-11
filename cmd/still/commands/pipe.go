package commands

import (
	"github.com/codegangsta/cli"
)

func NewPipeCommand() cli.Command {
	command := cli.Command{
		Name:      "pipe",
		ShortName: "p",
		Usage:     "Filter out needless text passed through pipe",
		Action:    actionPipe,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "model,m",
				Value: "model.still",
				Usage: "The input path of a model file to be used for filtering",
			},
		},
	}

	return command
}

func actionPipe(context *cli.Context) {
	// TODO:
}
