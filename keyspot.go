package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var versionString string = "1.0.27"

func main() {
	parser := argparse.NewParser("keyspot", "Tool for interfacing with KeySpot.app. Primarily used for running programs with KeySpot records injected as environment variables.")
	version := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Prints the version of the keyspot tool being used."})

	runCommand := parser.NewCommand("run", "Runs a command or program with environment variables from a given keyspot record.")

	initializeRun(runCommand)

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if *version {
		fmt.Println(versionString)
	}

	if runCommand.Happened() {
		executeRun()
	}
}
