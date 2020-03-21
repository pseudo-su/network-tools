package main

import (
	"log"
	"os"

	"github.com/pseudo-su/network-tools/internal"
	"github.com/urfave/cli/v2"
)

func FindOverlappingNetworksCommand() *cli.Command {
	return &cli.Command{
		Name:    "find-overlapping-networks",
		Aliases: []string{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input-file",
				Required: true,
				Aliases:  []string{"in"},
			},
			&cli.StringFlag{
				Name:     "output-file",
				Required: true,
				Aliases:  []string{"out"},
			},
		},
		Usage: "Find overlapping networks from csv input",
		Action: func(c *cli.Context) error {
			inputFile := c.String("input-file")
			outputFile := c.String("output-file")
			inputCSV, err := internal.ReadCSVFile(inputFile)
			if err != nil {
				return err
			}
			overlappingNetworks := internal.FindOverlappingNetworks(inputCSV)
			err = internal.WriteCSVFile(outputFile, overlappingNetworks)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			FindOverlappingNetworksCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
