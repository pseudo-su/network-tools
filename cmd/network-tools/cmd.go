package main

import (
	"io"
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
			inFile, err := os.Open(c.String("input-file"))
			if err != nil {
				return err
			}
			defer inFile.Close()

			networks, err := internal.ReadNetworks(inFile)
			overlappingNetworks := internal.FindOverlappingNetworks(networks)
			csvOutput := [][]string{
				[]string{"Network", "Subnet"},
			}
			for key, subnets := range overlappingNetworks {
				for _, subnet := range subnets {
					csvOutput = append(csvOutput, []string{key.String(), subnet.String()})
				}
			}

			var out io.Writer
			if c.String("output-file") == "stdout" {
				out = os.Stdout
			} else {

				outFile, err := os.OpenFile(c.String("output-file"), os.O_RDWR|os.O_CREATE, 0755)
				if err != nil {
					return err
				}
				defer outFile.Close()
				out = outFile
			}
			err = internal.WriteCSV(out, csvOutput)
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
