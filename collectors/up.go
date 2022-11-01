package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type UpCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewUpCollector(db *sql.DB) *UpCollector {
	desc := prometheus.NewDesc("mysql_up", "mysql hhealth", nil, nil)
	return &UpCollector{newBaseCollector(db), desc}
}

func (c *UpCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}
func (c *UpCollector) Collect(metrics chan<- prometheus.Metric) {
	var up float64 = 1
	if err := c.db.Ping(); err != nil {
		up = 0
	}
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, up)
}
