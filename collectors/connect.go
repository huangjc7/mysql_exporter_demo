package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type ConnectCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewConnectCollector(db *sql.DB) *ConnectCollector {
	desc := prometheus.NewDesc("mysql_connect_num", "mysql_connect_num", nil, nil)
	return &ConnectCollector{newBaseCollector(db), desc}
}

func (c *ConnectCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *ConnectCollector) Collect(metrics chan<- prometheus.Metric) {
	count := c.variables("max_connections")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}
