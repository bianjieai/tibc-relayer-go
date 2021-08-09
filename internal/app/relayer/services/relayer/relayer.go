package relayer

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
)

var _ IRelayer = new(Relayer)

type IRelayer interface {
	UpdateClient() error
	PendingDatagrams() error
	IsNotRelay() bool
	Context() *domain.Context
}

type Relayer struct {
	source repostitory.IChain
	dest   repostitory.IChain

	chainName string
	height    uint64

	context *domain.Context
}

func NewRelayer(source repostitory.IChain, dest repostitory.IChain, height uint64, chainName string) IRelayer {
	return &Relayer{
		source:    source,
		dest:      dest,
		chainName: chainName,
		height:    height,
	}
}

func (rly *Relayer) UpdateClient() error {
	// todo
	_, err := rly.dest.GetLightClientState(rly.source.ChainName())
	if err != nil {
		return err
	}

	return nil
}

func (rly *Relayer) PendingDatagrams() error {
	// todo
	return nil
}

func (rly *Relayer) IsNotRelay() bool {
	// todo
	return true
}

func (rly *Relayer) Context() *domain.Context {
	return rly.context
}
