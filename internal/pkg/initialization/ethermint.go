package initialization

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	repoethermint "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/ethermint"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
	coretypes "github.com/irisnet/core-sdk-go/types"
	corestore "github.com/irisnet/core-sdk-go/types/store"
	log "github.com/sirupsen/logrus"
)

func ethermintChain(cfg *configs.ChainCfg, logger *log.Logger) repostitory.IChain {
	logger.WithFields(log.Fields{
		"chain_name": cfg.Ethermint.ChainName,
	}).Info(" init chain start")

	ethermintCfg := repoethermint.NewConfig()

	// Tendermint
	tendermintCfg := repoethermint.NewTendermintConfig()
	tendermintCfg.ChainID = cfg.Ethermint.TendermintChainID
	tendermintCfg.GrpcAddr = cfg.Ethermint.GrpcAddr
	tendermintCfg.RPCAddr = cfg.Ethermint.RPCAddr
	options := []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(corestore.NewMemory(nil))),
		coretypes.TimeoutOption(cfg.Ethermint.RequestTimeout),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(cfg.Ethermint.Gas),
		coretypes.CachedOption(true),
	}
	if cfg.Ethermint.Algo != "" {
		options = append(options, coretypes.AlgoOption(cfg.Ethermint.Algo))
	}
	tendermintCfg.Options = options

	ethermintCfg.Tendermint = tendermintCfg

	// Eth
	contractCfgGroup := repoethermint.NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = cfg.Ethermint.Contracts.Packet.Addr
	contractCfgGroup.Packet.Topic = cfg.Ethermint.Contracts.Packet.Topic
	contractCfgGroup.Packet.OptPrivKey = cfg.Ethermint.Contracts.Packet.OptPrivKey

	contractCfgGroup.AckPacket.Addr = cfg.Ethermint.Contracts.AckPacket.Addr
	contractCfgGroup.AckPacket.Topic = cfg.Ethermint.Contracts.AckPacket.Topic
	contractCfgGroup.CleanPacket.Addr = cfg.Ethermint.Contracts.CleanPacket.Addr
	contractCfgGroup.CleanPacket.Topic = cfg.Ethermint.Contracts.CleanPacket.Topic

	contractCfgGroup.Client.Addr = cfg.Ethermint.Contracts.Client.Addr
	contractCfgGroup.Client.Topic = cfg.Ethermint.Contracts.Client.Topic
	contractCfgGroup.Client.OptPrivKey = cfg.Ethermint.Contracts.Client.OptPrivKey

	contractBindOptsCfg := repoethermint.NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = cfg.Ethermint.EthChainID
	contractBindOptsCfg.ClientPrivKey = cfg.Ethermint.Contracts.Client.OptPrivKey
	contractBindOptsCfg.PacketPrivKey = cfg.Ethermint.Contracts.Packet.OptPrivKey
	contractBindOptsCfg.GasLimit = cfg.Ethermint.GasLimit
	contractBindOptsCfg.MaxGasPrice = cfg.Ethermint.MaxGasPrice

	ethChainCfg := repoethermint.NewEthChainConfig()
	ethChainCfg.ContractCfgGroup = contractCfgGroup
	ethChainCfg.ContractBindOptsCfg = contractBindOptsCfg

	ethChainCfg.ChainURI = cfg.Ethermint.URI
	ethChainCfg.Slot = cfg.Ethermint.CommentSlot
	ethChainCfg.TipCoefficient = cfg.Ethermint.TipCoefficient
	ethermintCfg.Eth = ethChainCfg

	chainRepo, err := repoethermint.NewEthermintClient(
		constant.Ethermint,
		cfg.Ethermint.ChainName,
		cfg.Ethermint.UpdateClientFrequency,
		ethermintCfg,
	)
	if err != nil {
		logger.WithFields(log.Fields{
			"chain_name": cfg.Tendermint.ChainName,
			"err_msg":    err,
		}).Fatal("failed to init chain")
	}

	return chainRepo
}
