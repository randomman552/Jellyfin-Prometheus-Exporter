package collectors

import (
	"jellyfin-exporter/api"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type SessionsCollector struct {
	Client api.JellyfinClient

	ActiveSessions *prometheus.GaugeVec
	Streams        *prometheus.GaugeVec
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
		Streams: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "jellyfin_active_streams",
			Help: "The number of active streams running from Jellyfin",
		}, []string{
			"name",
			"container",
			"type",
			"mediaType",
			"paused",
			"playMethod",
		}),
	}
}

func (c *SessionsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

// Collect metrics from sessions returned from the Jellyfin API
func (c *SessionsCollector) Collect(metrics chan<- prometheus.Metric) {
	// Get data
	sessions := c.Client.GetSessions()

	c.CollectActiveSessionData(*sessions)
	c.CollectStreamData(*sessions)
}

// Collect data about sessions
func (c *SessionsCollector) CollectActiveSessionData(sessions []api.JellyfinSession) {
	grouped := GroupByProperty(sessions, func(s api.JellyfinSession) string {
		return s.Client
	})

	for key, value := range grouped {
		c.ActiveSessions.WithLabelValues(key).Set(float64(len(value)))
	}
}

// Collect information about streams
func (c *SessionsCollector) CollectStreamData(sessions []api.JellyfinSession) {
	sessions = Filter(sessions, func(s api.JellyfinSession) bool {
		return s.NowPlayingItem != nil
	})
	for _, session := range sessions {
		stream := session.NowPlayingItem
		containers := strings.Split(stream.Container, ",")

		c.Streams.WithLabelValues(
			stream.Name,
			containers[0],
			stream.Type,
			stream.MediaType,
			strconv.FormatBool(session.PlayState.IsPaused),
			session.PlayState.PlayMethod,
		).Set(1)
	}
}
