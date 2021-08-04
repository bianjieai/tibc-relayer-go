package mertics

import (
	metricsprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type Model struct {
	Sys   *metricsprometheus.Gauge
	Chain *metricsprometheus.Gauge

	chaiName string
}

func NewMetric(chaiName string) *Model {
	sysMetric := metricsprometheus.NewGaugeFrom(prometheus.GaugeOpts{
		Subsystem: "relayer",
		Name:      "system",
		Help:      "system status",
	}, []string{"locality"})

	chainMetric := metricsprometheus.NewGaugeFrom(prometheus.GaugeOpts{
		Subsystem: "relayer",
		Name:      "oss",
		Help:      "oss service status",
	}, []string{"locality"})

	model := &Model{
		Sys:      sysMetric,
		Chain:    chainMetric,
		chaiName: chaiName,
	}
	model.initMetric()

	return model
}

func (m *Model) initMetric() {
	labels := []string{"locality", m.chaiName}
	m.Sys.With(labels...).Set(1)
	m.Chain.With(labels...).Set(1)
}