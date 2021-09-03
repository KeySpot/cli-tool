package main

import (
	"fmt"
)

type VersionCommand struct {

}

func (x *VersionCommand) Execute(args []string) error {
	fmt.Println(versionString)

	return nil
}

var versionCommand VersionCommand

func init() {
	parser.AddCommand(
		"version",
		"Display version number",
		"The version command displays the version number for the command line tool using the <major>.<minor>.<patch> convention.",
		&versionCommand)
}