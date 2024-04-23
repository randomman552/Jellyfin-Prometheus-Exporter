# Jellyfin-Prometheus-Exporter
[![Build Status](https://drone.ggrainger.uk/api/badges/randomman552/Jellyfin-Prometheus-Exporter/status.svg)](https://drone.ggrainger.uk/randomman552/Jellyfin-Prometheus-Exporter)

A prometheus metrics exporter for Jellyfin.

Jellyfin does provide a metrics endpoint of it's own, but it doesn't provide any useful metrics aside from memory usage.

## Metrics
This exporter uses the Jellyfin REST API to generate metrics:
- Number of active sessions
  - Device client (Jellyfin Web, etc)
- Number of active streams
  - Media type (Video, Audio, etc)
  - Type (Movie, Episode, etc)
  - Stream type (transcoded, direct play)
  - Codec (hvec, h264)
- Number of items in each library
  - Container
  - Type (Movie, Epioside, etc)
- Number of user accounts
  - Is Admin
  - Authentication provider

## Configuration
The exporter is configured using environment variables
| Variable           | Default | Description                                                                   |
| :----------------- | :-----: | :---------------------------------------------------------------------------- |
| `JELLYFIN_URL`     |         | The url to reach the Jellyfin deployment, e.g. `https://jellyfin.example.com` |
| `JELLYFIN_API_KEY` |         | The API token to use when interacting with the Jellyfin API                   |
| `PORT`             |  2112   | The port to host on                                                           |

## Deployment
A docker image for this exporter is provided [randomman552/jellyfin-prometheus-exporter](https://hub.docker.com/repository/docker/randomman552/jellyfin-prometheus-exporter).

An example docker compose is provided below
```yml
version: "3"
services:
  jellyfin-exporter:
    image: randomman552/jellyfin-prometheus-exporter
    restart: unless-stopped
    environment:
      JELLYFIN_URL: https://jellyfin.example.com
      JELLYFIN_API_KEY: V3rySecretK3y
      PORT: 2112
    ports:
      - 2112:2112
```

You will probably want to set this up with a reverse proxy such as Traefik.