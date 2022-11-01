package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"mysql_exporter_demo/collectors"
	"net/http"
)

func main() {
	var listenAddress string
	flag.StringVar(&listenAddress, "H", "localhost:3306", "远程mysql主机地址+port吗,默认localhost:3306")
	flag.Parse()
	db, err := sql.Open("mysql", fmt.Sprintf("root:123456@tcp(%s)/mysql?charset=utf8mb4&loc=PRC&parseTime=true", listenAddress))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//定义指标
	prometheus.MustRegister(collectors.NewUpCollector(db))
	prometheus.MustRegister(collectors.NewSlowCollector(db))
	prometheus.MustRegister(collectors.NewConnectCollector(db))
	//注册指标
	//1.时间触发 2.业务触发 3.metrics请求触发

	//启动web服务
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9898", nil)
}
