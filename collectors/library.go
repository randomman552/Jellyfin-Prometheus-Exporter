package collectors

import (
	"jellyfin-exporter/api"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type LibraryCollector struct {
	Client api.JellyfinClient

	LibrariesGauge *prometheus.GaugeVec
}

func NewLibraryCollector(client *api.JellyfinClient) *LibraryCollector {
	return &LibraryCollector{
		Client: *client,

		LibrariesGauge: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "jellyfin_library_count",
			Help: "Number of items in each Jellyfin library",
		}, []string{
			"library",
			"libraryType",
			"itemType",
			"container",
		}),
	}
}

func (c *LibraryCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

// Collect metrics from sessions returned from the Jellyfin API
func (c *LibraryCollector) Collect(metrics chan<- prometheus.Metric) {
	virtualFolders := c.Client.GetVirtualFolders()

	if virtualFolders == nil {
		return
	}

	itemsPerFolder := make(map[string]api.JellyfinItemsResponse)

	// Get items
	for _, folder := range *virtualFolders {
		itemResponse := c.Client.GetItems(folder.ItemId)
		itemsPerFolder[folder.ItemId] = *itemResponse
	}

	// Update metrics
	c.LibrariesGauge.Reset()
	for _, folder := range *virtualFolders {
		itemResponse := itemsPerFolder[folder.ItemId]

		// Group items by their type
		// This is because a library can contain series, which can contain seasons, which can contain episodes...
		groupedByType := GroupByProperty(itemResponse.Items, func(item api.JellyfinItem) string {
			return item.Type
		})

		for itemType, items := range groupedByType {
			// Group by container for stats on media containers
			groupedByContainer := GroupByProperty(items, func(item api.JellyfinItem) string {
				return item.Container
			})

			// Finally, report to prometheus
			for container, items := range groupedByContainer {
				c.LibrariesGauge.WithLabelValues(folder.Name, folder.CollectionType, itemType, container).Set(float64(len(items)))
			}
		}
	}
}
