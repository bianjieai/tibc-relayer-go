package repostitory

import "github.com/bianjieai/tibc-sdk-go/packet"

type GetBlockHeaderReq struct {
	LatestHeight  uint64
	TrustedHeight uint64
}

type Packets struct {
	BizPackets   []packet.Packet
	AckPackets   []AckPacket
	CleanPackets []packet.CleanPacket
}

type AckPacket struct {
	Packet          packet.Packet
	Acknowledgement []byte
}

func newPackets() *Packets {
	return &Packets{
		BizPackets:   []packet.Packet{},
		AckPackets:   []AckPacket{},
		CleanPackets: []packet.CleanPacket{},
	}
}
