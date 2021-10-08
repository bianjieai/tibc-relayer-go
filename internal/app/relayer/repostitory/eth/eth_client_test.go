package eth

import (
	"fmt"
	"testing"

	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"

	gethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
)

func TestNewEth(t *testing.T) {
	//ropsten := "https://rinkeby.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810"
	ropsten := "https://ropsten.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810"
	optPrivKey := "c59f553aa4d23dad1db5b42aa8d72ce98223e29e4e6f873d95b1ced451edad39"
	var chainID uint64 = 3

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.Packet.Addr = "0xc1d8C2Daa09989D98CC8Dc157a525555190bEA9D"
	contractCfgGroup.Packet.Topic = "PacketSent((uint64,string,string,string,string,bytes))"
	contractCfgGroup.Packet.OptPrivKey = optPrivKey

	contractCfgGroup.AckPacket.Addr = "0xc1d8C2Daa09989D98CC8Dc157a525555190bEA9D"
	contractCfgGroup.AckPacket.Topic = "AckWritten((uint64,string,string,string,string,bytes),bytes)"
	contractCfgGroup.AckPacket.OptPrivKey = optPrivKey

	contractCfgGroup.CleanPacket.Addr = "0xc1d8C2Daa09989D98CC8Dc157a525555190bEA9D"
	contractCfgGroup.CleanPacket.Topic = "CleanPacketSent((uint64,string,string,string))"
	contractCfgGroup.CleanPacket.OptPrivKey = optPrivKey

	contractCfgGroup.Client.Addr = "0xB470BD3Ae2a3f1fd663cAD19113233A19Ab37842"
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
	ethClient.GetProof(
		"irishub-testnet",
		"eth-testnet",
		1,
		11128921,
		repotypes.AckProof,
	)

	//packets, err := ethClient.GetPackets(11128997)
	////packets, err := ethClient.GetPackets(11128966)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(packets)

}

func Test_Hex(t *testing.T) {

	str := "0000000000000000000000000000000000000000000000000000000000000003"
	dataBytes := gethcmn.HexToHash(str)
	args := abi.Arguments{
		abi.Argument{Type: Uint64},
	}

	headerBytes, err := args.Unpack(dataBytes.Bytes())
	if err != nil {
		return
	}
	fmt.Println("headerBytes: ", headerBytes)
}
