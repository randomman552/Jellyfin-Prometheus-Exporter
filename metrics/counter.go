package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type CounterCollector struct {
}

func NewCounterCollector() *CounterCollector {
	return &CounterCollector{}
}

func (c *CounterCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *CounterCollector) Collect(metrics chan<- prometheus.Metric) {
	metrics <- prometheus.MustNewConstMetric(opsProcessed.Desc(), prometheus.CounterValue, 0)
	opsProcessed.Add(1)
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "counter",
		Help: "Numbber of metric gathering cycles that have been performed",
	})
)
