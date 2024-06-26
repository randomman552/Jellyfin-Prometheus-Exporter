package collectors

import (
	"jellyfin-exporter/api"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type UsersCollector struct {
	Client api.JellyfinClient

	UsersGauge *prometheus.GaugeVec
}

func NewUsersCollector(client *api.JellyfinClient) *UsersCollector {
	return &UsersCollector{
		Client: *client,

		UsersGauge: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "jellyfin_user_accounts",
			Help: "The number of Jellyfin users",
		}, []string{
			"authProvider",
			"isActive",
			"isAdmin",
		}),
	}
}

func (c *UsersCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

// Collect metrics from sessions returned from the Jellyfin API
func (c *UsersCollector) Collect(metrics chan<- prometheus.Metric) {
	users := c.Client.GetUsers()

	c.UsersGauge.Reset()

	if users == nil {
		return
	}

	// First group by auth provider
	groupedByAuthProvider := GroupByProperty(*users, func(user api.JellyfinUser) string {
		return user.Policy.AuthenticationProvider
	})

	for authProvider, users := range groupedByAuthProvider {
		// Then group by disabled status
		groupedByDisabled := GroupByProperty(users, func(user api.JellyfinUser) bool {
			return user.Policy.IsDisabled
		})

		for isDisabled, users := range groupedByDisabled {
			// Then group by admin status
			groupedByAdmin := GroupByProperty(users, func(user api.JellyfinUser) bool {
				return user.Policy.IsAdministrator
			})

			for isAdmin, users := range groupedByAdmin {
				c.UsersGauge.WithLabelValues(
					authProvider,
					strconv.FormatBool(!isDisabled),
					strconv.FormatBool(isAdmin),
				).Set(float64(len(users)))
			}
		}
	}
}
