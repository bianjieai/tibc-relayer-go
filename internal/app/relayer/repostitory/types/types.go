package types

import "github.com/bianjieai/tibc-sdk-go/packet"

type GetBlockHeaderReq struct {
	LatestHeight   uint64
	TrustedHeight  uint64
	RevisionNumber uint64
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

func NewPackets() *Packets {
	return &Packets{
		BizPackets:   []packet.Packet{},
		AckPackets:   []AckPacket{},
		CleanPackets: []packet.CleanPacket{},
	}
}

type ResultTx struct {
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Hash      string `json:"hash"`
	Height    int64  `json:"height"`
}
