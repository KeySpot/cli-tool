package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type ExportCommand struct {
	Key    string `short:"k" long:"key" description:"Specify a record by access key to export the values from the record as a .env file."`
	Record string `short:"r" long:"record" description:"Specify a record associated with your KeySpot account by name to export the values from the record as a .env file. Requires your keyspot cli tool to be configured to your account, run $ keyspot configure --help for info on how to configure."`
}

func exportVars(path string, record *map[string]string) error {
	fileString := ""
	for k, v := range *record {
		fileString += fmt.Sprintf("%s=%s\n", k, v)
	}

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(fileString)

	if err != nil {
		return err
	}

	return nil
}

func (x *ExportCommand) Execute(args []string) error {
	var secrets *map[string]string

	if x.Key != "" {
		route := fmt.Sprintf("/%s", x.Key)
		jsonString, err := apiGet(route, "")

		if err != nil {
			return err
		}

		secrets = recordFromJsonString(jsonString)
	}

	if x.Record != "" {
		sub, token, err := getSubAndToken()

		if err != nil {
			return err
		}

		route := fmt.Sprintf("/user-records/%s/%s", sub, x.Record)
		jsonString, err := apiGet(route, token)

		if err != nil {
			return err
		}

		document := documentFromJsonString(jsonString)
		secrets = &document.Record
	}

	if len(args) == 0 {
		return flags.ErrExpectedArgument
	}

	err := exportVars(args[0], secrets)

	if err != nil {
		return err
	}

	return nil
}

func (x *ExportCommand) Usage() string {
	return "<NEW-FILE-PATH> [export command options]"
}

var exportCommand ExportCommand

func init() {
	parser.AddCommand(
		"export",
		"Export a record stored with KeySpot as a .env file.",
		"Export a record stored with KeySpot as a file following the .env file syntax. Use one of the options to specify which KeySpot record to export. <NEW-FILE-PATH> is the name of the new exported file.",
		&exportCommand)
}
