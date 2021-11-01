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
	}, []string{"chain_name"})

	chainMetric := metricsprometheus.NewGaugeFrom(prometheus.GaugeOpts{
		Subsystem: "relayer",
		Name:      "chain",
		Help:      "chain status",
	}, []string{"chain_name", "option"})

	model := &Model{
		Sys:      sysMetric,
		Chain:    chainMetric,
		chaiName: chaiName,
	}
	model.initMetric()

	return model
}

func (m *Model) initMetric() {
	sysLabels := []string{"chain_name", m.chaiName}
	m.Sys.With(sysLabels...).Set(1)

	connChainLabels := []string{"chain_name", m.chaiName, "option", "connection"}
	getClientStatusLabels := []string{"chain_name", m.chaiName, "option", "client_get_client_status"}
	updateClientLabels := []string{"chain_name", m.chaiName, "option", "client_update_client_status"}
	recvPacketLabels := []string{"chain_name", m.chaiName, "option", "packet_recv_packet"}
	getPacketLabels := []string{"chain_name", m.chaiName, "option", "packet_get_packet"}
	getCommitmentLabels := []string{"chain_name", m.chaiName, "option", "packet_get_commitment"}
	getProofLabels := []string{"chain_name", m.chaiName, "option", "packet_get_proof"}
	getReceiptLabels := []string{"chain_name", m.chaiName, "option", "packet_get_receipt"}

	m.Chain.With(connChainLabels...).Set(1)
	m.Chain.With(getClientStatusLabels...).Set(1)
	m.Chain.With(updateClientLabels...).Set(1)
	m.Chain.With(recvPacketLabels...).Set(1)
	m.Chain.With(getPacketLabels...).Set(1)
	m.Chain.With(getCommitmentLabels...).Set(1)
	m.Chain.With(getProofLabels...).Set(1)
	m.Chain.With(getReceiptLabels...).Set(1)
}
