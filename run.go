package main

import (
	"strings"

	"github.com/jessevdk/go-flags"
	keyspot "github.com/keyspot/gopackage"
)

type RunCommand struct {
	Key string `short:"k" long:"key" description:"Specify a record by access key to run the command with the values from the record as environment variables."`
}

func (x *RunCommand) Execute(args []string) error {
	if x.Key != "" {
		keyspot.SetEnvironment(x.Key)
	}

	if (len(args) == 0) {
		return flags.ErrExpectedArgument
	}

	err := exec_command(strings.Join(args, " "))

	return err
}

var runCommand RunCommand

func init() {
	parser.AddCommand(
		"run",
		"Run shell command",
		"The run command will run a command in your current path.",
		&runCommand)
}
