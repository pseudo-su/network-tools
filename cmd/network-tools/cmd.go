package main

import (
	"io"
	"log"
	"os"

	"github.com/pseudo-su/network-tools/internal"
	"github.com/urfave/cli/v2"
)

func FindSubnetsCommand() *cli.Command {
	return &cli.Command{
		Name:    "find-subnets",
		Aliases: []string{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input-file",
				Value:   "stdin",
				Aliases: []string{"in"},
			},
			&cli.StringFlag{
				Name:    "output-file",
				Value:   "stdout",
				Aliases: []string{"out"},
			},
		},
		Usage: "Find subnets from csv input",
		Action: func(c *cli.Context) error {
			// Setup input reader
			var in io.Reader
			if c.String("input-file") == "stdin" {
				in = os.Stdin
			} else {
				inFile, err := os.Open(c.String("input-file"))
				if err != nil {
					return err
				}
				defer inFile.Close()
				in = inFile
			}

			// Setup output writer
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

			// Read networks and find subnets
			networks, err := internal.ReadNetworks(in)
			subnets := internal.FindSubnets(networks)

			// Format csv output
			csvOutput := [][]string{
				[]string{"Network", "Subnet"},
			}
			for key, subnets := range subnets {
				for _, subnet := range subnets {
					csvOutput = append(csvOutput, []string{key.String(), subnet.String()})
				}
			}

			// write output
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
			FindSubnetsCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
