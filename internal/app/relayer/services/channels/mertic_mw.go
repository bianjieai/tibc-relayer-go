package channels

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	internelerrors "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
	merticsmodel "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/mertics"
)

var _ IChannel = new(Metric)

type Metric struct {
	next IChannel

	metricsModel *merticsmodel.Model
}

func NewMetricMW(svc IChannel, metricsModel *merticsmodel.Model) IChannel {

	return &Metric{
		next:         svc,
		metricsModel: metricsModel,
	}
}

func (m *Metric) UpdateClientFrequency() uint64 {
	return m.next.UpdateClientFrequency()
}

func (m *Metric) UpdateClient() error {
	return m.next.UpdateClient()
}

func (m *Metric) Relay() error {
	err := m.next.Relay()
	defer func(err error) {
		labels := []string{"chain_name", m.Context().ChainName()}

		connChainLabels := []string{"chain_name", m.Context().ChainName(), "option", "connection"}
		getClientStatusLabels := []string{"chain_name", m.Context().ChainName(), "option", "client_get_client_status"}
		updateClientLabels := []string{"chain_name", m.Context().ChainName(), "option", "client_update_client_status"}
		recvPacketLabels := []string{"chain_name", m.Context().ChainName(), "option", "packet_recv_packet"}
		getPacketLabels := []string{"chain_name", m.Context().ChainName(), "option", "packet_get_packet"}
		getCommitmentLabels := []string{"chain_name", m.Context().ChainName(), "option", "packet_get_commitment"}
		getProofLabels := []string{"chain_name", m.Context().ChainName(), "option", "packet_get_proof"}
		getReceiptLabels := []string{"chain_name", m.Context().ChainName(), "option", "packet_get_receipt"}

		sysErr, ok := err.(internelerrors.IError)
		if !ok && sysErr != nil {
			m.metricsModel.Sys.With(labels...).Set(-1)
			return
		}
		switch sysErr {
		case internelerrors.ErrChainConn:
			m.metricsModel.Chain.With(connChainLabels...).Set(-1)
		case internelerrors.ErrGetLightClientState:
			m.metricsModel.Chain.With(getClientStatusLabels...).Set(-1)
		case internelerrors.ErrUpdateClient:
			m.metricsModel.Chain.With(updateClientLabels...).Set(-1)
		case internelerrors.ErrRecvPacket:
			m.metricsModel.Chain.With(recvPacketLabels...).Set(-1)
		case internelerrors.ErrGetPackets:
			m.metricsModel.Chain.With(getPacketLabels...).Set(-1)
		case internelerrors.ErrGetCommitmentPacket:
			m.metricsModel.Chain.With(getCommitmentLabels...).Set(-1)
		case internelerrors.ErrGetProof:
			m.metricsModel.Chain.With(getProofLabels...).Set(-1)
		case internelerrors.ErrGetReceiptPacket:
			m.metricsModel.Chain.With(getReceiptLabels...).Set(-1)

		default:
			m.metricsModel.Sys.With(labels...).Set(1)
			m.metricsModel.Chain.With(connChainLabels...).Set(1)
			m.metricsModel.Chain.With(getClientStatusLabels...).Set(1)
			m.metricsModel.Chain.With(updateClientLabels...).Set(1)
			m.metricsModel.Chain.With(recvPacketLabels...).Set(1)
			m.metricsModel.Chain.With(getPacketLabels...).Set(1)
			m.metricsModel.Chain.With(getCommitmentLabels...).Set(1)
			m.metricsModel.Chain.With(getProofLabels...).Set(1)
			m.metricsModel.Chain.With(getReceiptLabels...).Set(1)
		}
	}(err)
	return err
}

func (m *Metric) IsNotRelay() bool {
	return m.next.IsNotRelay()
}

func (m *Metric) Context() *domain.Context {
	return m.next.Context()
}
