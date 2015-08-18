package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mitsuse/still"
)

const (
	Name        = "still"
	Version     = "0.1.0"
	Description = "A command-line tool to filter out needless text by using statistical classifier."
	Author      = "Tomoya Kose (mitsuse)"
	AuthorEmail = "tomoya@mitsuse.jp"
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", Name, err)
}

func readExamples(path string) ([]*still.Example, error) {
	examplesFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer examplesFile.Close()

	examples := make([]*still.Example, 0)

	if err := json.NewDecoder(examplesFile).Decode(&examples); err != nil {
		return nil, err
	}

	return examples, nil
}

func writeModel(s *still.Still, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := s.Serialize(file); err != nil {
		return err
	}

	return nil
}
