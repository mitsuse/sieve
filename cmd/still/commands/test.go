package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func NewTestCommand() cli.Command {
	command := cli.Command{
		Name:      "test",
		ShortName: "t",
		Usage:     "Test a model by filtering out text with the given data",
		Action:    actionTest,

		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "model,m",
				Value: "model.still",
				Usage: "The path of a model to be tested",
			},

			cli.StringFlag{
				Name:  "examples,e",
				Value: "examples.json",
				Usage: "The path of a file containing test examples",
			},
		},
	}

	return command
}

func actionTest(context *cli.Context) {
	p := &testParameters{
		ExamplesPath: context.String("examples"),
		ModelPath:    context.String("model"),
	}

	examples, err := readExamples(p.ExamplesPath)
	if err != nil {
		printError(err)
		return
	}

	s, err := readModel(context.String("model"))
	if err != nil {
		printError(err)
		return
	}

	truePositive := 0
	falsePositive := 0
	falseNegative := 0

	for _, example := range examples {
		filtered := s.Filter(example.Text)
		if example.Class == 0 {
			if filtered {
				truePositive += 1
			} else {
				falseNegative += 1
			}
		} else {
			if filtered {
				falsePositive += 1
			}
		}
	}

	precision := float64(truePositive) / float64(truePositive+falsePositive)
	recall := float64(truePositive) / float64(truePositive+falseNegative)
	f := (2 * precision * recall) / (precision + recall)

	fmt.Printf("precision: %.2f\n", precision)
	fmt.Printf("recall: %.2f\n", recall)
	fmt.Printf("f: %.2f\n", f)
}

type testParameters struct {
	ExamplesPath string
	ModelPath    string
}
