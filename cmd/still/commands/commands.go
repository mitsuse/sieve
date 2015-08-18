package commands

import (
	"fmt"
	"os"
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
