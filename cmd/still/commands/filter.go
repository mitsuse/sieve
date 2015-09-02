package commands

import (
	"os"

	"github.com/codegangsta/cli"
)

func NewFilterCommand() cli.Command {
	command := cli.Command{
		Name:      "filter",
		ShortName: "f",
		Usage:     "Filter out needless text passed through pipe",
		Action:    actionFilter,

		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "filtered,f",
				Usage: "Print filtered-out text intead",
			},

			cli.StringFlag{
				Name:  "model,m",
				Value: "model.still",
				Usage: "The input path of a model file to be used for filtering",
			},
		},
	}

	return command
}

func actionFilter(context *cli.Context) {
	s, err := readModel(context.String("model"))
	if err != nil {
		printError(err)
		return
	}

	if err := useWithPipe(s, os.Stdin, os.Stdout, context.Bool("filtered")); err != nil {
		printError(err)
		return
	}
}
