package mertics

import (
	metricsprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type Model struct {
	Sys   *metricsprometheus.Gauge
	Chain *metricsprometheus.Gauge

	sourceChainName string
	destChainName   string
}

func NewMetric(sourceChainName, destChainName string) *Model {
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
		Sys:             sysMetric,
		Chain:           chainMetric,
		sourceChainName: sourceChainName,
		destChainName:   destChainName,
	}
	model.initMetric()

	return model
}

func (m *Model) initMetric() {
	sourceSysLabels := []string{"chain_name", m.sourceChainName}
	m.Sys.With(sourceSysLabels...).Set(1)
	destSysLabels := []string{"chain_name", m.destChainName}
	m.Sys.With(destSysLabels...).Set(1)

	sourceConnChainLabels := []string{"chain_name", m.sourceChainName, "option", "connection"}
	sourceGetClientStatusLabels := []string{"chain_name", m.sourceChainName, "option", "client_get_client_status"}
	sourceUpdateClientLabels := []string{"chain_name", m.sourceChainName, "option", "client_update_client_status"}
	sourceRecvPacketLabels := []string{"chain_name", m.sourceChainName, "option", "packet_recv_packet"}
	sourceGetPacketLabels := []string{"chain_name", m.sourceChainName, "option", "packet_get_packet"}
	sourceGetCommitmentLabels := []string{"chain_name", m.sourceChainName, "option", "packet_get_commitment"}
	sourceGetProofLabels := []string{"chain_name", m.sourceChainName, "option", "packet_get_proof"}
	sourceGetReceiptLabels := []string{"chain_name", m.sourceChainName, "option", "packet_get_receipt"}

	destConnChainLabels := []string{"chain_name", m.destChainName, "option", "connection"}
	destGetClientStatusLabels := []string{"chain_name", m.destChainName, "option", "client_get_client_status"}
	destUpdateClientLabels := []string{"chain_name", m.destChainName, "option", "client_update_client_status"}
	destRecvPacketLabels := []string{"chain_name", m.destChainName, "option", "packet_recv_packet"}
	destGetPacketLabels := []string{"chain_name", m.destChainName, "option", "packet_get_packet"}
	destGetCommitmentLabels := []string{"chain_name", m.destChainName, "option", "packet_get_commitment"}
	destGetProofLabels := []string{"chain_name", m.destChainName, "option", "packet_get_proof"}
	destGetReceiptLabels := []string{"chain_name", m.destChainName, "option", "packet_get_receipt"}

	m.Chain.With(sourceConnChainLabels...).Set(1)
	m.Chain.With(sourceGetClientStatusLabels...).Set(1)
	m.Chain.With(sourceUpdateClientLabels...).Set(1)
	m.Chain.With(sourceRecvPacketLabels...).Set(1)
	m.Chain.With(sourceGetPacketLabels...).Set(1)
	m.Chain.With(sourceGetCommitmentLabels...).Set(1)
	m.Chain.With(sourceGetProofLabels...).Set(1)
	m.Chain.With(sourceGetReceiptLabels...).Set(1)

	m.Chain.With(destConnChainLabels...).Set(1)
	m.Chain.With(destGetClientStatusLabels...).Set(1)
	m.Chain.With(destUpdateClientLabels...).Set(1)
	m.Chain.With(destRecvPacketLabels...).Set(1)
	m.Chain.With(destGetPacketLabels...).Set(1)
	m.Chain.With(destGetCommitmentLabels...).Set(1)
	m.Chain.With(destGetProofLabels...).Set(1)
	m.Chain.With(destGetReceiptLabels...).Set(1)
}
