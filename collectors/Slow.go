package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type SlowCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewSlowCollector(db *sql.DB) *SlowCollector {
	desc := prometheus.NewDesc("mysql_slow_query", "mysql slow query", nil, nil)
	return &SlowCollector{newBaseCollector(db), desc}
}

func (c *SlowCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *SlowCollector) Collect(metrics chan<- prometheus.Metric) {
	count := c.status("slow_queries")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}
