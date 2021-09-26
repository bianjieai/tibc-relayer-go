package eth

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
)

func TestNewEth(t *testing.T) {
	//ropsten := "https://rinkeby.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810"
	ropsten := "https://ropsten.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810"
	contractCfgGroup := NewContracts()
	contractCfgGroup.Packet.Addr = "0x4ecD91c7e4481c8a00254356600962F880c62152"
	contractCfgGroup.Packet.Topic = "PacketSent((uint64,string,string,string,string,bytes) packet)"
	contractCfgGroup.Packet.OptPrivKey = "c59f553aa4d23dad1db5b42aa8d72ce98223e29e4e6f873d95b1ced451edad39"

	contractCfgGroup.AckPacket.Addr = "0x4ecD91c7e4481c8a00254356600962F880c62152"
	contractCfgGroup.AckPacket.Topic = "AckWritten((uint64,string,string,string,string,bytes) packet, bytes ack)"
	contractCfgGroup.AckPacket.OptPrivKey = "c59f553aa4d23dad1db5b42aa8d72ce98223e29e4e6f873d95b1ced451edad39"

	contractCfgGroup.CleanPacket.Addr = "0x4ecD91c7e4481c8a00254356600962F880c62152"
	contractCfgGroup.CleanPacket.Topic = "CleanPacketSent((uint64,string,string,string) packet)"
	contractCfgGroup.CleanPacket.OptPrivKey = "c59f553aa4d23dad1db5b42aa8d72ce98223e29e4e6f873d95b1ced451edad39"

	contractCfgGroup.Client.Addr = "0x776763E02f04445fC3346E99c4dA8588AcA2FD8C"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = "c59f553aa4d23dad1db5b42aa8d72ce98223e29e4e6f873d95b1ced451edad39"

	ethClient, err := NewEth(
		constant.ETH,
		"ETH",
		10,
		ropsten,
		3,
		contractCfgGroup)
	if err != nil {
		t.Fatal(err)
	}
	latestHeight, err := ethClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(latestHeight)

	packets, err := ethClient.GetPackets(11107269)
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
