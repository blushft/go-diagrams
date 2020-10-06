package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "gen",
		Commands: []*cli.Command{
			{
				Name:   "colors",
				Usage:  "generate colors package",
				Action: genColorsCmd,
			},
			{
				Name:   "assets",
				Usage:  "generate asset nodes",
				Action: genAssetsCmd,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
