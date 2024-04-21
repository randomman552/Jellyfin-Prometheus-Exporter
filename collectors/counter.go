package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type CounterCollector struct {
	Counter prometheus.Counter
}

func NewCounterCollector() *CounterCollector {
	return &CounterCollector{
		Counter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "counter",
			Help: "Numbber of metric gathering cycles that have been performed",
		}),
	}
}

func (c *CounterCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *CounterCollector) Collect(metrics chan<- prometheus.Metric) {
	// Register metrics
	metrics <- prometheus.MustNewConstMetric(c.Counter.Desc(), prometheus.CounterValue, 0)

	// Get data
	// TODO

	// Update metrics
	c.Counter.Add(1)
}
