package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mitsuse/still"
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
	stillFile, err := os.Open(context.String("model"))
	if err != nil {
		// TODO: Show error message.
		return
	}

	s, err := still.Deserialize(stillFile)
	if err != nil {
		// TODO: Show error message.
		return
	}
	stillFile.Close()

	if err := filterWithIo(s, os.Stdin, os.Stdout); err != nil {
		// TODO: Show error message.
		return
	}
}

func filterWithIo(s *still.Still, reader io.Reader, writer io.Writer) (err error) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		if !s.Filter(text) {
			continue
		}

		_, err = fmt.Fprintln(writer, text)
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}
