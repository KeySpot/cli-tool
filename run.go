package main

import (
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	keyspot "github.com/keyspot/gopackage"
)

type RunCommand struct {
	Key    string `short:"k" long:"key" description:"Specify a record by access key to run the command with the values from the record as environment variables."`
	Record string `short:"r" long:"record" description:"Specify a record associated with your KeySpot account by name to run the command with the values from the record as environment variables. Requires your keyspot cli tool to be configured to your account, run $ keyspot configure --help for info on how to configure."`
}

func injectRecord(record *map[string]string) {
	for key, value := range *record {
		os.Setenv(key, value)
	}
}

func recordCallback(recordName string) error {
	configFilePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	token, err := getToken(configFilePath)

	if err != nil {
		return err
	}

	payload, err := parseJwtPayload(token)

	if err != nil {
		return err
	}

	jsonString, err := apiGet("/user-records/"+payload.Sub+"/"+recordName, token)

	if err != nil {
		return err
	}

	record := recordFromJsonString(jsonString)

	injectRecord(record)

	return nil
}

func (x *RunCommand) Execute(args []string) error {
	if x.Key != "" {
		keyspot.SetEnvironment(x.Key)
	}

	if x.Record != "" {
		recordCallback(x.Record)
	}

	if len(args) == 0 {
		return flags.ErrExpectedArgument
	}

	err := exec_command(strings.Join(args, " "))

	return err
}

func (x *RunCommand) Usage() string {
	return "<COMMAND> [run-OPTIONS]"
}

var runCommand RunCommand

func init() {
	parser.AddCommand(
		"run",
		"Run shell command with record as environment variables",
		"The run command will run a command your terminal has access to with environment variables injected from a KeySpot Record. The method for specifying which record is through the various options. If no options are given, the command will be run without any injection.",
		&runCommand)
}
