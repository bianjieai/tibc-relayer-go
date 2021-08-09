package relayer

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	merticsmodel "github.com/bianjieai/tibc-relayer-go/internal/pkg/mertics"
	internelerrors "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
)

var _ IRelayer = new(Metric)

type Metric struct {
	next IRelayer

	metricsModel *merticsmodel.Model
}

func NewMetricMW(svc IRelayer, metricsModel *merticsmodel.Model) IRelayer {

	return &Metric{
		next:         svc,
		metricsModel: metricsModel,
	}
}

func (m *Metric) UpdateClient() error {
	return m.next.UpdateClient()
}

func (m *Metric) PendingDatagrams() error {
	err := m.next.PendingDatagrams()
	defer func(err error) {
		labels := []string{"chain_name", m.next.Context().ChainName()}
		sysErr, ok := err.(internelerrors.IError)
		if !ok {
			m.metricsModel.Sys.With(labels...).Set(-1)
			return
		}
		switch sysErr {
		case internelerrors.ErrChainConn:
			m.metricsModel.Chain.With(labels...).Set(-1)

		default:
			m.metricsModel.Sys.With(labels...).Set(1)
			m.metricsModel.Chain.With(labels...).Set(1)
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
