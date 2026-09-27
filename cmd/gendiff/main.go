package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"code"
)

// formatFlag customizes help rendering so the placeholder is shown as "string".
type formatFlag struct {
	*cli.StringFlag
}

func (f *formatFlag) String() string {
	return strings.ReplaceAll(cli.FlagStringer(f), "output string", "output format")
}

func (f *formatFlag) GetUsage() string {
	return "output `string`"
}

func main() {
	app := &cli.App{
		Name:            "gendiff",
		Usage:           "Compares two configuration files and shows a difference.",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&formatFlag{
				StringFlag: &cli.StringFlag{
					Name:    "format",
					Aliases: []string{"f"},
					Value:   code.FormatStylish,
					Usage:   "output format",
				},
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return cli.Exit("Expected two file paths", 1)
			}

			result, err := code.GenDiff(c.Args().Get(0), c.Args().Get(1), c.String("format"))
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
