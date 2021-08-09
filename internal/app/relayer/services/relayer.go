package services

import "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"

var _ IRelayer = new(Relayer)

type IRelayer interface {
	UpdateClient() error
	PendingDatagrams()
	IsNotRelay() bool
}

type Relayer struct {
	source repostitory.IChain
	dest   repostitory.IChain
}

func (rly *Relayer) UpdateClient() error {
	// todo
	_, err := rly.dest.GetLightClientState(rly.source.ChainName())
	if err != nil {
		return err
	}

	return nil
}

func (rly *Relayer) PendingDatagrams() {
	// todo
}

func (rly *Relayer) IsNotRelay() bool {
	// todo
	return true
}
