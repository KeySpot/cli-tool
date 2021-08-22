package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/akamensky/argparse"
	keyspot "github.com/keyspot/gopackage"
)

type command struct {
	Name string
	Help string
	Run  func([]string)
}

var versionString string = "1.0.16"

func main() {
	parser := argparse.NewParser("keyspot", "Tool for interfacing with KeySpot.app. Primarily used for running programs with KeySpot records injected as environment variables.")
	version := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Prints the version of the keyspot tool being used."})

	runCommand := parser.NewCommand("run", "Runs a command or program with environment variables from a given keyspot record.")

	command := runCommand.String("c", "command", &argparse.Options{Required: true, Help: "Command to be run."})
	accessKey := runCommand.String("k", "key", &argparse.Options{Required: false, Help: "Determines the record to be run by the access key provided."})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if *version {
		fmt.Println(versionString)
	}

	if runCommand.Happened() {

		if *accessKey != "" {
			keyspot.SetEnvironment(*accessKey)
		}

		commandArray := strings.Split(*command, " ")
		cmd := exec.Command(commandArray[0], commandArray[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
