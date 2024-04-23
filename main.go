package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"

	"jellyfin-exporter/api"
	"jellyfin-exporter/collectors"
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
		&cli.StringFlag{
			Name: "port",
			EnvVars: []string{
				"PORT",
			},
			Value: "2112",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	// Create Jellyfin API client
	jellyfinUrl := c.String("jellyfin.url")
	jellyfinToken := c.String("jellyfin.api_key")
	apiClient := api.NewJellyfinClient(jellyfinUrl, jellyfinToken)

	registry := prometheus.NewRegistry()

	// Register collectors
	collectors := []prometheus.Collector{
		collectors.NewSessionsCollector(apiClient),
		collectors.NewLibraryCollector(apiClient),
		collectors.NewUsersCollector(apiClient),
	}

	registry.MustRegister(collectors...)

	// Start gather as a background operation
	go func() {
		for {
			registry.Gather()
			time.Sleep(60 * time.Second)
		}
	}()

	// Start hosting
	port := c.String("port")
	http.Handle("/metrics", promhttp.Handler())

	log.Print("Listening on port " + port)
	http.ListenAndServe(":"+port, nil)
	return nil
}
