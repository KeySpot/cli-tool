package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type ConfigureCommand struct {
}

func writeConfigFile(path string, token string) error {
	configFilePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	f, err := os.Create(configFilePath)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(token)

	if err != nil {
		return err
	}

	err = os.Chmod(configFilePath, 0600)

	if err != nil {
		return err
	}

	return nil
}

func validAccount(token string) (bool, error) {
	payload, err := parseJwtPayload(token)

	if err != nil {
		return false, err
	}

	_, err = apiGet("/user-records/count/"+payload.Sub, token)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (x *ConfigureCommand) Execute(args []string) error {
	if len(args) == 0 {
		return flags.ErrExpectedArgument
	}

	token := args[0]

	accountValid, err := validAccount(token)

	if err != nil {
		return err
	}

	if !accountValid {
		return errors.New("Invalid account info in json web token")
	}

	configFilePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	err = writeConfigFile(configFilePath, token)

	if err != nil {
		return err
	}

	fmt.Println("Account successfully linked")

	return nil
}

func (x *ConfigureCommand) Usage() string {
	return "<CLI-TOKEN>"
}

var configureCommand ConfigureCommand

func init() {
	parser.AddCommand(
		"configure",
		"Configure the keyspot cli tool to an account",
		fmt.Sprintf("When given a cli token from the KeySpot website %s/account, running the configure command with the token will link the token's account to the keyspot cli tool. This allows a user to specify documents by name instead of just by access key.", websiteUrl),
		&configureCommand)
}
