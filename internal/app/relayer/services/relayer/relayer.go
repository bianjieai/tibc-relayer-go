package relayer

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	typeserr "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
	log "github.com/sirupsen/logrus"
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
	signer    string
	context   *domain.Context

	logger log.Logger
}

func NewRelayer(source repostitory.IChain, dest repostitory.IChain, height uint64) IRelayer {

	return &Relayer{
		source:    source,
		dest:      dest,
		chainName: source.ChainName(),
		height:    height,
		context:   domain.NewContext(height, source.ChainName()),
	}
}

func (rly *Relayer) UpdateClient() error {

	// 1. get light client state from dest chain
	clientState, err := rly.dest.GetLightClientState(rly.source.ChainName())
	if err != nil {
		rly.logger.Info("failed to get light client state")
		return typeserr.ErrGetLightClientState
	}

	// 2. get source chain updated latest height from dest chain
	heightObj := clientState.GetLatestHeight()
	height := heightObj.GetRevisionHeight()

	logger := rly.logger.WithFields(log.Fields{
		"height": height,
	})

	// 3. get nextHeight block header from source chain
	nextHeight := height + 1
	header, err := rly.source.GetBlockHeader(nextHeight)
	if err != nil {
		logger.Error("failed to get block header")
		return typeserr.ErrGetBlockHeader
	}

	// 4. update client to dest chain
	if err := rly.dest.UpdateClient(header); err != nil {
		logger.Error("failed to update client")
		return typeserr.ErrUpdateClient
	}

	return nil
}

func (rly *Relayer) PendingDatagrams() error {
	// todo
	return nil
}

func (rly *Relayer) IsNotRelay() bool {
	curHeight := rly.context.Height()
	latestHeight, err := rly.source.GetLatestHeight()
	if err != nil {
		return false
	}

	if curHeight < latestHeight {
		return true
	}

	return false
}

func (rly *Relayer) Context() *domain.Context {
	return rly.context
}
