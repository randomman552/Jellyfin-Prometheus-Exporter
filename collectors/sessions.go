package collectors

import (
	"jellyfin-exporter/api"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type SessionsCollector struct {
	Client api.JellyfinClient

	ActiveSessionsGauge *prometheus.GaugeVec
	ActiveStreamsGauge  *prometheus.GaugeVec
}

func NewSessionsCollector(client *api.JellyfinClient) *SessionsCollector {
	return &SessionsCollector{
		Client: *client,

		ActiveSessionsGauge: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "jellyfin_active_sessions",
			Help: "The number of active Jellyfin sessions",
		}, []string{
			"client",
		}),
		ActiveStreamsGauge: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "jellyfin_active_streams",
			Help: "The number of active streams running from Jellyfin",
		}, []string{
			"codec",
			"name",
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

	if sessions == nil {
		return
	}

	c.ActiveSessionsGauge.Reset()
	c.ActiveStreamsGauge.Reset()

	c.CollectActiveSessionData(*sessions)
	c.CollectStreamData(*sessions)
}

// Collect data about sessions
func (c *SessionsCollector) CollectActiveSessionData(sessions []api.JellyfinSession) {
	grouped := GroupByProperty(sessions, func(s api.JellyfinSession) string {
		return s.Client
	})

	for key, value := range grouped {
		c.ActiveSessionsGauge.WithLabelValues(key).Set(float64(len(value)))
	}
}

// Collect information about streams
func (c *SessionsCollector) CollectStreamData(sessions []api.JellyfinSession) {
	// Only show sessions that are playing
	sessions = Filter(sessions, func(s api.JellyfinSession) bool {
		return s.NowPlayingItem != nil
	})

	for _, session := range sessions {
		item := session.NowPlayingItem
		codec := "unknown"

		// If we are dealing with video, try to determine the codec
		if item.MediaType == "Video" {
			videoStreams := Filter(item.MediaStreams, func(s api.JellyfinMediaStream) bool {
				return s.Type == "Video"
			})

			if len(videoStreams) > 0 {
				codec = videoStreams[0].Codec
			}
		}

		c.ActiveStreamsGauge.WithLabelValues(
			codec,
			item.Name,
			item.Type,
			item.MediaType,
			strconv.FormatBool(session.PlayState.IsPaused),
			session.PlayState.PlayMethod,
		).Set(1)
	}
}
