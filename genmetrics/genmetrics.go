// etcdmetrics
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/text"
	"github.com/prometheus/client_model/go"

	"open-falcon-plugins/pluginlibs"
)

var (
	parser     text.Parser
	filters    string
	url        string
	in         string
	filterMode bool
)

func main() {
	flag.StringVar(&url, "d", "http://etcd.sdp.nd:2379/metrics", "url of api,like and default http://etcd.sdp.nd:2379/metrics")
	flag.StringVar(&filters, "f", "", "filter list of metric don't collect,like 'test_etcd,skydns_count'")
	flag.BoolVar(&filterMode, "m", true, "if set true, allow list. if set false, deney list")
	flag.Parse()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("not Data!")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if string(body) != "" {
		in = string(body)
	}

	i := strings.NewReader(in)

	metricsfamilies, err := parser.TextToMetricFamilies(i)

	if err != nil {
		log.Fatal("parse error")
		return
	}
	m := pluginlibs.NewMetricsEcho()

	if !filterMode {
		m.SetFilterMode(false)
	}
	if filters != "" {
		m.SetFilter(filters)
	}

	var metric, value, counterType string
	for name, family := range metricsfamilies {
		s := pluginlibs.NewMetricSingle()
		metric = name
		if family.GetType() == io_prometheus_client.MetricType_COUNTER {
			value = strconv.FormatFloat(family.GetMetric()[0].GetCounter().GetValue(), 'f', -1, 64)
			//			value = fmt.Sprintf("%.3f", family.GetMetric()[0].GetCounter().GetValue())
			counterType = "COUNTER"
		} else if family.GetType() == io_prometheus_client.MetricType_GAUGE {
			value = strconv.FormatFloat(family.GetMetric()[0].GetGauge().GetValue(), 'f', -1, 64)
			//			value = fmt.Sprintf("%.3f", family.GetMetric()[0].GetGauge().GetValue())
			counterType = "GAUGE"
		} else {
			continue
		}
		s.SetMetric(metric, value, counterType)
		m.AddMetricSingle(*s)
	}
	fmt.Println(m)
}
