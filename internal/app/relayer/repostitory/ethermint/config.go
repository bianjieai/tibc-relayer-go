package ethermint

import (
	coretypes "github.com/irisnet/core-sdk-go/types"
)

type Config struct {
	Tendermint *TendermintConfig
	Eth        *EthChainConfig
}

func NewConfig() *Config {
	return &Config{}
}

///===============
// eth config

type TendermintConfig struct {
	Options []coretypes.Option

	RPCAddr  string
	GrpcAddr string
	ChainID  string
}

func NewTendermintConfig() *TendermintConfig {
	return &TendermintConfig{}
}

///===============
// eth config

func NewEthChainConfig() *EthChainConfig {
	return &EthChainConfig{}
}

type EthChainConfig struct {
	ChainURI string

	Slot           int64
	TipCoefficient float64

	ContractCfgGroup    *ContractCfgGroup
	ContractBindOptsCfg *ContractBindOptsCfg
}

func NewContractCfgGroup() *ContractCfgGroup {
	return &ContractCfgGroup{}
}

type ContractCfgGroup struct {
	Client      ContractCfg
	Packet      ContractCfg
	AckPacket   ContractCfg
	CleanPacket ContractCfg
}

type ContractCfg struct {
	Addr       string
	Topic      string
	OptPrivKey string
}

func NewContractBindOptsCfg() *ContractBindOptsCfg {
	return &ContractBindOptsCfg{}
}

type ContractBindOptsCfg struct {
	ClientPrivKey string
	PacketPrivKey string
	GasLimit      uint64
	MaxGasPrice   uint64
	ChainID       uint64
}
