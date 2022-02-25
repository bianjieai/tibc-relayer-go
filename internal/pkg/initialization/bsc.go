package initialization

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	repobsc "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/bsc"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
	log "github.com/sirupsen/logrus"
)

func bscChain(cfg *configs.ChainCfg, logger *log.Logger) repostitory.IChain {
	loggerEntry := logger.WithFields(log.Fields{
		"chain_name": cfg.Bsc.ChainName,
	})

	loggerEntry.Info(" init eth chain start")

	contractCfgGroup := repobsc.NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = cfg.Bsc.Contracts.Packet.Addr
	contractCfgGroup.Packet.Topic = cfg.Bsc.Contracts.Packet.Topic
	contractCfgGroup.Packet.OptPrivKey = cfg.Bsc.Contracts.Packet.OptPrivKey

	contractCfgGroup.AckPacket.Addr = cfg.Bsc.Contracts.AckPacket.Addr
	contractCfgGroup.AckPacket.Topic = cfg.Bsc.Contracts.AckPacket.Topic
	contractCfgGroup.CleanPacket.Addr = cfg.Bsc.Contracts.CleanPacket.Addr
	contractCfgGroup.CleanPacket.Topic = cfg.Bsc.Contracts.CleanPacket.Topic

	contractCfgGroup.Client.Addr = cfg.Bsc.Contracts.Client.Addr
	contractCfgGroup.Client.Topic = cfg.Bsc.Contracts.Client.Topic
	contractCfgGroup.Client.OptPrivKey = cfg.Bsc.Contracts.Client.OptPrivKey

	contractBindOptsCfg := repobsc.NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = cfg.Bsc.ChainID
	contractBindOptsCfg.ClientPrivKey = cfg.Bsc.Contracts.Client.OptPrivKey
	contractBindOptsCfg.PacketPrivKey = cfg.Bsc.Contracts.Packet.OptPrivKey
	contractBindOptsCfg.GasLimit = cfg.Bsc.GasLimit
	contractBindOptsCfg.MaxGasPrice = cfg.Bsc.MaxGasPrice

	bscChainCfg := repobsc.NewChainConfig()
	bscChainCfg.ContractCfgGroup = contractCfgGroup
	bscChainCfg.ContractBindOptsCfg = contractBindOptsCfg

	bscChainCfg.ChainType = constant.BSC
	bscChainCfg.ChainName = cfg.Bsc.ChainName
	bscChainCfg.ChainID = cfg.Bsc.ChainID
	bscChainCfg.ChainURI = cfg.Bsc.URI
	bscChainCfg.Slot = cfg.Bsc.CommentSlot
	bscChainCfg.UpdateClientFrequency = cfg.Bsc.UpdateClientFrequency
	bscChainCfg.TipCoefficient = cfg.Bsc.TipCoefficient

	ethRepo, err := repobsc.NewBsc(bscChainCfg)
	if err != nil {
		loggerEntry.WithFields(log.Fields{
			"err_msg": err,
		}).Fatal("failed to init chain")
	}

	return ethRepo
}
