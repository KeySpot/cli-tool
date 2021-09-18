package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type ImportCommand struct {
}

func importVars(path string, recordName string) error {
	sub, token, err := getSubAndToken()

	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	jsonString := "{"

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		jsonString += (`"` + split[0] + `"` + ":" + `"` + split[1] + `"` + ",")
	}

	jsonString = jsonString[:len(jsonString)-1] + "}"

	_, err = apiPut("/user-records/"+sub+"/"+recordName, jsonString, token)

	if err != nil {
		return err
	}

	return nil
}

func (x *ImportCommand) Execute(args []string) error {
	if len(args) == 0 {
		return flags.ErrExpectedArgument
	}

	err := importVars(args[0], args[1])

	if err != nil {
		return err
	}

	return nil
}

func (x *ImportCommand) Usage() string {
	return "<FILE-PATH> <RECORD-NAME>"
}

var importCommand ImportCommand

func init() {
	parser.AddCommand(
		"import",
		"Import a .env file into KeySpot. Requires your CLI tool to be configured.",
		"Import a .env file into KeySpot. The FILE-PATH specified is the .env file you want to import while RECORD-NAME is the name of the record in KeySpot to create. If RECORD-NAME already exists in your account then the record will be overwritten. Requires your CLI tool to be configured.",
		&importCommand)
}
