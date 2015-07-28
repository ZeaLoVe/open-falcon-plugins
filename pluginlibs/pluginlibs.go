// pluginlibs
package pluginlibs

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

//open-falcon plugin upload fomat
type MetricSingle struct {
	EndPoint    string `json:"endpoint,omitempty"`
	Tags        string `json:"tags,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	Metric      string `json:"metric,omitempty"`
	Value       string `json:"value,omitempty"`
	CounterType string `json:"counterType,omitempty"`
	Step        int64  `json:"step,omitempty"`
}

func (m *MetricSingle) SetTags(tags string) {
	m.Tags = tags
}

func (m *MetricSingle) SetMetric(metric string, value string, counterType string) {
	m.Value = value
	m.Metric = metric
	m.CounterType = counterType
}

type MetricsEcho struct {
	Metrics    []MetricSingle  `json:Metrics",omitempty"`
	infilter   map[string]bool `json:"-"`
	filterMode bool            `json:"-"`
}

//filterMode == true ,allow list
//filterMode == false ,deney list
func (m *MetricsEcho) AddMetricSingle(single MetricSingle) {
	if m.filterMode {
		if m.infilter[single.Metric] {
			m.Metrics = append(m.Metrics, single)
		} else {
			return
		}
	} else {
		if m.infilter[single.Metric] {
			return
		} else {
			m.Metrics = append(m.Metrics, single)
		}
	}
}

func (m *MetricsEcho) SetFilterMode(flag bool) {
	m.filterMode = flag
}

func (m *MetricsEcho) SetFilter(args string) {
	filters := strings.Split(args, ",")
	for _, filter := range filters {
		m.infilter[filter] = true
	}
}

func (m MetricsEcho) String() string {
	str, err := json.Marshal(&m)
	if err != nil {
		return ""
	} else {
		return string(str)[11 : len(str)-1]
	}
}

func NewMetricsEcho() *MetricsEcho {
	return &MetricsEcho{
		infilter:   make(map[string]bool, 100),
		filterMode: true,
	}
}

func NewMetricSingle() *MetricSingle {
	ts := time.Now().Unix()
	host, err := os.Hostname()
	if err != nil {
		return &MetricSingle{}
	} else {
		return &MetricSingle{
			EndPoint:  host,
			Timestamp: ts,
			Step:      60,
		}
	}
}
