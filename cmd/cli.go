package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:     "Search for Events Overlapping",
		Flags:    createFlags(createFileFlag, createTextFlag),
		Commands: createCommands(createImportCMD, createValidateCMD),
		Before:   validateCLI,
	}

	return app.Run(os.Args)
}

func validateCLI(c *cli.Context) error {
	if c == nil {
		return fmt.Errorf("context cannot be nil.")
	}

	if len(c.FlagNames()) > 1 {
		return fmt.Errorf("only one flag is allowed.")
	}

	return nil
}
