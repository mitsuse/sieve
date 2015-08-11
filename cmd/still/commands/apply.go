package commands

import (
	"github.com/codegangsta/cli"
)

func NewApplyCommand() cli.Command {
	command := cli.Command{
		Name:      "apply",
		ShortName: "a",
		Usage:     "Filter out needless text read from the given file",
		Action:    actionApply,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "model,m",
				Value: "model.still",
				Usage: "The input path of a model file to be used for filtering",
			},

			cli.StringFlag{
				Name:  "input,i",
				Value: "input.txt",
				Usage: "The input path of a text file to be filtered out",
			},

			cli.StringFlag{
				Name:  "output,o",
				Value: "output.txt",
				Usage: "The output path of the result file for filtering",
			},
		},
	}

	return command
}

func actionApply(context *cli.Context) {
	// TODO:
}
