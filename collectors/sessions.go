package collectors

import (
	"jellyfin-exporter/api"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type SessionsCollector struct {
	Client api.JellyfinClient

	ActiveSessions *prometheus.GaugeVec
}

func NewSessionsCollector(client *api.JellyfinClient) *SessionsCollector {
	return &SessionsCollector{
		Client: *client,

		ActiveSessions: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "jellyfin_active_sessions",
			Help: "The number of active Jellyfin sessions",
		}, []string{
			"client",
		}),
	}
}

func (c *SessionsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *SessionsCollector) Collect(metrics chan<- prometheus.Metric) {
	// Get data
	sessions := c.Client.GetSessions()

	c.CollectActiveSessionData(*sessions)
}

func (c *SessionsCollector) CollectActiveSessionData(sessions []api.JellyfinSession) {
	grouped := GroupByProperty(sessions, func(s api.JellyfinSession) string {
		return s.Client
	})

	for key, value := range grouped {
		c.ActiveSessions.WithLabelValues(key).Set(float64(len(value)))
	}
}
