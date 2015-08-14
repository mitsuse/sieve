package commands

import (
	"encoding/json"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/still"
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
		},
	}

	return command
}

func actionBuild(context *cli.Context) {
	exampleSeq := make([]*still.Example, 0)

	examplesFile, err := os.Open(context.String("examples"))
	if err != nil {
		// TODO: Show error message.
		return
	}

	if err := json.NewDecoder(examplesFile).Decode(&exampleSeq); err != nil {
		// TODO: Show error message.
		return
	}
	examplesFile.Close()

	s := still.Learn(exampleSeq)

	stillFile, err := os.Create(context.String("model"))
	if err != nil {
		// TODO: Show error message.
		return
	}

	if err := s.Serialize(stillFile); err != nil {
		// TODO: Show error message.
		return
	}
	stillFile.Close()
}
