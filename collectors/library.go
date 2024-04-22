package collectors

import (
	"jellyfin-exporter/api"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type LibraryCollector struct {
	Client api.JellyfinClient

	Libraries *prometheus.GaugeVec
}

func NewLibraryCollector(client *api.JellyfinClient) *LibraryCollector {
	return &LibraryCollector{
		Client: *client,

		Libraries: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "jellyfin_library_count",
			Help: "Number of items in each Jellyfin library",
		}, []string{
			"library",
			"libraryType",
			"itemType",
		}),
	}
}

func (c *LibraryCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

// Collect metrics from sessions returned from the Jellyfin API
func (c *LibraryCollector) Collect(metrics chan<- prometheus.Metric) {
	virtualFolders := c.Client.GetVirtualFolders()

	for _, folder := range virtualFolders {
		// Get items
		itemResponse := c.Client.GetItems(folder.ItemId)

		// Group items by their type
		// This is because a library can contain series, which can contain seasons, which can contain episodes...
		groupedItems := GroupByProperty(itemResponse.Items, func(item api.JellyfinItem) string {
			return item.Type
		})

		for key, items := range groupedItems {
			c.Libraries.WithLabelValues(folder.Name, folder.CollectionType, key).Set(float64(len(items)))
		}
	}
}
