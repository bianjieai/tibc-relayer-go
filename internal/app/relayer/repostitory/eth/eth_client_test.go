package eth

import (
	"encoding/hex"
	"fmt"
	"testing"

	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
)

func TestNewEth(t *testing.T) {
	//ropsten := "https://rinkeby.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810"
	ropsten := "https://ropsten.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810"
	optPrivKey := "c59f553aa4d23dad1db5b42aa8d72ce98223e29e4e6f873d95b1ced451edad39"
	var chainID uint64 = 3

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = "0xF06Ba39bce333442Dfa477e313D6439Ac7dc89c4"
	contractCfgGroup.Packet.Topic = "PacketSent((uint64,string,string,string,string,bytes))"
	contractCfgGroup.Packet.OptPrivKey = optPrivKey

	contractCfgGroup.AckPacket.Addr = "0xF06Ba39bce333442Dfa477e313D6439Ac7dc89c4"
	contractCfgGroup.AckPacket.Topic = "AckWritten((uint64,string,string,string,string,bytes),bytes)"
	contractCfgGroup.AckPacket.OptPrivKey = optPrivKey

	contractCfgGroup.CleanPacket.Addr = "0xF06Ba39bce333442Dfa477e313D6439Ac7dc89c4"
	contractCfgGroup.CleanPacket.Topic = "CleanPacketSent((uint64,string,string,string))"
	contractCfgGroup.CleanPacket.OptPrivKey = optPrivKey

	contractCfgGroup.Client.Addr = "0x776763E02f04445fC3346E99c4dA8588AcA2FD8C"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = optPrivKey

	contractBindOptsCfg := NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = chainID
	contractBindOptsCfg.ClientPrivKey = optPrivKey
	contractBindOptsCfg.PacketPrivKey = optPrivKey
	contractBindOptsCfg.GasLimit = 2000000
	contractBindOptsCfg.GasPrice = 1500000000

	chainCfg := NewChainConfig()
	chainCfg.ContractCfgGroup = contractCfgGroup
	chainCfg.ContractBindOptsCfg = contractBindOptsCfg
	chainCfg.ChainType = constant.ETH
	chainCfg.ChainName = "ETH"
	chainCfg.ChainURI = ropsten
	chainCfg.ChainID = chainID
	chainCfg.UpdateClientFrequency = 10

	ethClient, err := NewEth(chainCfg)
	if err != nil {
		t.Fatal(err)
	}
	latestHeight, err := ethClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(latestHeight)

	packets, err := ethClient.GetProof(
		"iris-testnet",
		"eth-testnet",
		27,
		11122903, repotypes.AckProof)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(packets)

}

func Test_Hex(t *testing.T) {
	args := abi.Arguments{
		abi.Argument{Type: Uint64, Name: "revision_number"},
	}

	headerBytes, _ := args.Pack(
		0,
	)
	fmt.Println("headerBytes: ", hex.EncodeToString(headerBytes))
}
