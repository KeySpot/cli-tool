// +build windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/akamensky/argparse"
	keyspot "github.com/keyspot/gopackage"
)

var command *string
var accessKey *string

func initializeRun(runCommand *argparse.Command) {
	command = runCommand.String("c", "command", &argparse.Options{Required: true, Help: "Command to be run."})
	accessKey = runCommand.String("k", "key", &argparse.Options{Required: false, Help: "Determines the record to be run by the access key provided."})
}

func executeRun() {
	if *accessKey != "" {
		keyspot.SetEnvironment(*accessKey)
	}

	commandArray := strings.Split(*command, " ")
	cmd := exec.Command("powershell", commandArray...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
