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
		"chain_name": cfg.Eth.ChainName,
	})

	loggerEntry.Info(" init eth chain start")

	contractCfgGroup := repoeth.NewContractCfgGroup()
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

	contractBindOptsCfg := repoeth.NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = cfg.Eth.ChainID
	contractBindOptsCfg.ClientPrivKey = cfg.Eth.Contracts.Client.OptPrivKey
	contractBindOptsCfg.PacketPrivKey = cfg.Eth.Contracts.Packet.OptPrivKey
	contractBindOptsCfg.GasLimit = cfg.Eth.GasLimit
	contractBindOptsCfg.MaxGasPrice = cfg.Eth.MaxGasPrice

	ethChainCfg := repoeth.NewChainConfig()
	ethChainCfg.ContractCfgGroup = contractCfgGroup
	ethChainCfg.ContractBindOptsCfg = contractBindOptsCfg

	ethChainCfg.ChainType = constant.ETH
	ethChainCfg.ChainName = cfg.Eth.ChainName
	ethChainCfg.ChainID = cfg.Eth.ChainID
	ethChainCfg.ChainURI = cfg.Eth.URI
	ethChainCfg.Slot = cfg.Eth.CommentSlot
	ethChainCfg.UpdateClientFrequency = cfg.Eth.UpdateClientFrequency
	ethChainCfg.TipCoefficient = cfg.Eth.TipCoefficient

	ethRepo, err := repoeth.NewEth(ethChainCfg)
	if err != nil {
		loggerEntry.WithFields(log.Fields{
			"err_msg": err,
		}).Fatal("failed to init chain")
	}

	return ethRepo
}
