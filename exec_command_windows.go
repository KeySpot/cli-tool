// +build windows

package main

import (
	"os"
	"os/exec"
	"strings"
)

func exec_command(commandString string) error {
	commandArray := strings.Split(commandString, " ")
	cmd := exec.Command("powershell", commandArray...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return err
}