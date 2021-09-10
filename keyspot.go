package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

var versionString string = "1.1.3"

var opts struct {
}

var parser = flags.NewParser(&opts, flags.Default)

func main() {
	_, err := parser.Parse()

	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
