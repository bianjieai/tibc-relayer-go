package initialization

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	repoeth "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/eth"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
	log "github.com/sirupsen/logrus"
)

func ethChain(cfg *configs.ChainCfg, logger *log.Logger) repostitory.IChain {
	loggerEntry := logger.WithFields(log.Fields{
		"chain_name": cfg.Tendermint.ChainName,
	})

	loggerEntry.Info(" init chain start")

	contractCfgGroup := repoeth.NewContracts()
	contractCfgGroup.Packet.Addr = cfg.Eth.Contracts.Packet.Addr
	contractCfgGroup.Packet.Topic = cfg.Eth.Contracts.Packet.Topic
	contractCfgGroup.Packet.OptPrivKey = cfg.Eth.Contracts.Packet.OptPrivKey

	contractCfgGroup.AckPacket.Addr = cfg.Eth.Contracts.AckPacket.Addr
	contractCfgGroup.AckPacket.Topic = cfg.Eth.Contracts.AckPacket.Topic
	contractCfgGroup.CleanPacket.Addr = cfg.Eth.Contracts.CleanPacket.Addr
	contractCfgGroup.CleanPacket.Topic = cfg.Eth.Contracts.CleanPacket.Topic

	contractCfgGroup.Client.Addr = cfg.Eth.Contracts.Client.Addr
	contractCfgGroup.Client.Topic = cfg.Eth.Contracts.Client.Topic
	contractCfgGroup.Client.OptPrivKey = cfg.Eth.Contracts.Client.OptPrivKey

	ethRepo, err := repoeth.NewEth(constant.ETH,
		cfg.Eth.ChainName,
		cfg.Eth.UpdateClientFrequency,
		cfg.Eth.URI,
		cfg.Eth.ChainID,
		contractCfgGroup)
	if err != nil {
		loggerEntry.Fatal(err)
	}

	return ethRepo
}
