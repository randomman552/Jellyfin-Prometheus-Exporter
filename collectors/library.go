package collectors

import (
	"jellyfin-exporter/api"

	"github.com/prometheus/client_golang/prometheus"
)

type LibraryCollector struct {
	Client api.JellyfinClient
}

func NewLibraryCollector(client *api.JellyfinClient) *LibraryCollector {
	return &LibraryCollector{
		Client: *client,
	}
}

func (c *LibraryCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

// Collect metrics from sessions returned from the Jellyfin API
func (c *LibraryCollector) Collect(metrics chan<- prometheus.Metric) {

}
