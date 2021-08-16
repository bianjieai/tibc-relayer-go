package channels

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	typeserr "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
	log "github.com/sirupsen/logrus"
)

var _ IChannel = new(Channel)

type IChannel interface {
	UpdateClient() error
	PendingDatagrams() error
	IsNotRelay() bool
	Context() *domain.Context
}

type Channel struct {
	source repostitory.IChain
	dest   repostitory.IChain

	chainName string
	height    uint64
	signer    string
	context   *domain.Context

	logger log.Logger
}

func NewChannel(source repostitory.IChain, dest repostitory.IChain, height uint64) IChannel {

	return &Channel{
		source:    source,
		dest:      dest,
		chainName: source.ChainName(),
		height:    height,
		context:   domain.NewContext(height, source.ChainName()),
	}
}

func (channel *Channel) UpdateClient() error {

	// 1. get light client state from dest chain
	clientState, err := channel.dest.GetLightClientState(channel.source.ChainName())
	if err != nil {
		channel.logger.Info("failed to get light client state")
		return typeserr.ErrGetLightClientState
	}

	// 2. get source chain updated latest height from dest chain
	heightObj := clientState.GetLatestHeight()
	height := heightObj.GetRevisionHeight()

	logger := channel.logger.WithFields(log.Fields{
		"height": height,
	})

	// 3. get nextHeight block header from source chain
	nextHeight := height + 1
	header, err := channel.source.GetBlockHeader(nextHeight)
	if err != nil {
		logger.Error("failed to get block header")
		return typeserr.ErrGetBlockHeader
	}

	// 4. update client to dest chain
	if err := channel.dest.UpdateClient(header); err != nil {
		logger.Error("failed to update client")
		return typeserr.ErrUpdateClient
	}

	return nil
}

func (channel *Channel) PendingDatagrams() error {
	// todo
	return nil
}

func (channel *Channel) IsNotRelay() bool {
	curHeight := channel.context.Height()
	latestHeight, err := channel.source.GetLatestHeight()
	if err != nil {
		return false
	}

	if curHeight < latestHeight {
		return true
	}

	return false
}

func (channel *Channel) Context() *domain.Context {
	return channel.context
}
