package ethermint

import (
	"math/big"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/ethermint/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcmn "github.com/ethereum/go-ethereum/common"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	gethethclient "github.com/ethereum/go-ethereum/ethclient"
)

// ==================================================================================================================
// contract bind opts
type bindOpts struct {
	client             *bind.TransactOpts
	packetTransactOpts *bind.TransactOpts
}

func newBindOpts(cfg *ContractBindOptsCfg) (*bindOpts, error) {

	cliPriv, err := gethcrypto.HexToECDSA(cfg.ClientPrivKey)
	if err != nil {
		return nil, err
	}
	clientOpts, err := bind.NewKeyedTransactorWithChainID(cliPriv, new(big.Int).SetUint64(cfg.ChainID))
	if err != nil {
		return nil, err
	}
	clientOpts.GasLimit = cfg.GasLimit

	//================================================================================
	// packet transfer opts
	packPriv, err := gethcrypto.HexToECDSA(cfg.PacketPrivKey)
	if err != nil {
		return nil, err
	}
	packOpts, err := bind.NewKeyedTransactorWithChainID(packPriv, new(big.Int).SetUint64(cfg.ChainID))
	if err != nil {
		return nil, err
	}
	packOpts.GasLimit = cfg.GasLimit

	return &bindOpts{
		client:             clientOpts,
		packetTransactOpts: packOpts,
	}, nil
}

// ==================================================================================================================
// contract client group
type contractGroup struct {
	Packet *contracts.Packet
	Client *contracts.Client
}

func newContractGroup(ethClient *gethethclient.Client, cfgGroup *ContractCfgGroup) (*contractGroup, error) {
	packAddr := gethcmn.HexToAddress(cfgGroup.Packet.Addr)
	packetFilter, err := contracts.NewPacket(packAddr, ethClient)
	if err != nil {
		return nil, err
	}

	clientAddr := gethcmn.HexToAddress(cfgGroup.Client.Addr)
	clientFilter, err := contracts.NewClient(clientAddr, ethClient)
	if err != nil {
		return nil, err
	}

	return &contractGroup{
		Packet: packetFilter,
		Client: clientFilter,
	}, nil
}
