package eth

import (
	"testing"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
)

func TestNewEth(t *testing.T) {
	ropsten := "https://rinkeby.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810"
	contractCfgGroup := NewContracts()
	contractCfgGroup.Packet.Addr = ""
	contractCfgGroup.Packet.Topic = "PacketSent((uint64,string,string,string,string,bytes) packet)"
	contractCfgGroup.Packet.OptPrivKey = "45760456b8181a0c3a313e8d9031b1f9343b1f45baaf5043262c19b63b163d5f"
	contractCfgGroup.AckPacket.Addr = ""
	contractCfgGroup.AckPacket.Topic = "AckWritten((uint64,string,string,string,string,bytes) packet, bytes ack)"
	contractCfgGroup.CleanPacket.Addr = ""
	contractCfgGroup.CleanPacket.Topic = "CleanPacketSent((uint64,string,string,string) packet)"
	contractCfgGroup.Client.Addr = "0xfE9edD8c6948F5Eb30Efb29D4a6F576Bf42d6266"
	contractCfgGroup.Client.Topic = ""
	contractCfgGroup.Client.OptPrivKey = "45760456b8181a0c3a313e8d9031b1f9343b1f45baaf5043262c19b63b163d5f"

	ethClient, err := NewEth(constant.ETH, "ETH", 10, ropsten, contractCfgGroup)
	if err != nil {
		t.Fatal(err)
	}
	latestHeight, err := ethClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(latestHeight)
	clientStatus, err := ethClient.GetLightClientState("irishub-testnet10")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("clientStatus: ", clientStatus.GetLatestHeight())

}
