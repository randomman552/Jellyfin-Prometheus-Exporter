package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	godotenv.Load()

	app := cli.NewApp()
	app.Name = "Jellyfin Prometheus Exporter"
	app.Usage = "Prometheus metrics exporter for Jellfin"
	app.Action = run

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "jellyfin.url",
			EnvVars: []string{
				"JELLYFIN_URL",
			},
		},
		&cli.StringFlag{
			Name: "jellyfin.api_key",
			EnvVars: []string{
				"JELLYFIN_API_KEY",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	print("Hello World!")
	return nil
}
