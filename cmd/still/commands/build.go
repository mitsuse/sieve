package commands

import (
	"errors"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/still"
)

const (
	nonPositiveIterations = "\"iterations\" should be a positive number."
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
				Usage: "The number of iterations for examples",
			},
		},
	}

	return command
}

func actionBuild(context *cli.Context) {
	p, err := validateBuildParameters(context)
	if err != nil {
		printError(err)
		return
	}

	examples, err := readExamples(p.ExamplesPath)
	if err != nil {
		printError(err)
		return
	}

	s := still.Learn(examples)

	if err := writeModel(s, p.ModelPath); err != nil {
		printError(err)
		return
	}
}

type buildParameters struct {
	ExamplesPath string
	ModelPath    string
	Iterations   int
}

func validateBuildParameters(context *cli.Context) (*buildParameters, error) {
	iterations := context.Int("iterations")
	if iterations <= 0 {
		return nil, errors.New(nonPositiveIterations)
	}

	p := &buildParameters{
		ExamplesPath: context.String("examples"),
		ModelPath:    context.String("model"),
		Iterations:   iterations,
	}

	return p, nil
}
