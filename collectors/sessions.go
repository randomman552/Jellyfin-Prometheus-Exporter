package collectors

import (
	"jellyfin-exporter/api"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type SessionsCollector struct {
	Client api.JellyfinClient

	ActiveSessions prometheus.Gauge
}

func NewSessionsCollector(client *api.JellyfinClient) *SessionsCollector {
	return &SessionsCollector{
		Client: *client,

		ActiveSessions: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "jellyfin_active_sessions",
			Help: "The number of active Jellyfin sessions",
		}),
	}
}

func (c *SessionsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *SessionsCollector) Collect(metrics chan<- prometheus.Metric) {
	// Create metrics
	metrics <- prometheus.MustNewConstMetric(c.ActiveSessions.Desc(), prometheus.GaugeValue, 0)

	// Get data
	sessions := c.Client.GetSessions()

	// Update metrics
	c.ActiveSessions.Set(float64(len(*sessions)))
}
