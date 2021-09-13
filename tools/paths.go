package tools

import (
	"fmt"
)

// KVStore key prefixes for IBC
const (
	KeySequencePrefix              = "sequences"
	KeyNextSeqSendPrefix           = "nextSequenceSend"
	KeyPacketCommitmentPrefix      = "commitments"
	KeyPacketAckPrefix             = "acks"
	KeyPacketReceiptPrefix         = "receipts"
	KeyCleanPacketCommitmentPrefix = "clean"
)

// NextSequenceSendPath defines the next send sequence counter store path
func NextSequenceSendPath(sourceChain, destChain string) string {
	return fmt.Sprintf("%s/%s", KeyNextSeqSendPrefix, packetPath(sourceChain, destChain))
}

// NextSequenceSendKey returns the store key for the send sequence of a particular
// channel binded to a specific port.
func NextSequenceSendKey(sourceChain, destChain string) []byte {
	return []byte(NextSequenceSendPath(sourceChain, destChain))
}

// PacketCommitmentPath defines the commitments to packet data fields store path
func PacketCommitmentPath(sourceChain, destinationChain string, sequence uint64) string {
	return fmt.Sprintf("%s/%d", PacketCommitmentPrefixPath(sourceChain, destinationChain), sequence)
}

// PacketCommitmentKey returns the store key of under which a packet commitment
// is stored
func PacketCommitmentKey(sourceChain, destinationChain string, sequence uint64) []byte {
	return []byte(PacketCommitmentPath(sourceChain, destinationChain, sequence))
}

// PacketCommitmentPrefixPath defines the prefix for commitments to packet data fields store path.
func PacketCommitmentPrefixPath(sourceChain, destinationChain string) string {
	return fmt.Sprintf("%s/%s/%s", KeyPacketCommitmentPrefix, packetPath(sourceChain, destinationChain), KeySequencePrefix)
}

// PacketAcknowledgementPath defines the packet acknowledgement store path
func PacketAcknowledgementPath(sourceChain, destinationChain string, sequence uint64) string {
	return fmt.Sprintf("%s/%d", PacketAcknowledgementPrefixPath(sourceChain, destinationChain), sequence)
}

// PacketAcknowledgementKey returns the store key of under which a packet
// acknowledgement is stored
func PacketAcknowledgementKey(sourceChain, destinationChain string, sequence uint64) []byte {
	return []byte(PacketAcknowledgementPath(sourceChain, destinationChain, sequence))
}

// PacketAcknowledgementPrefixPath defines the prefix for commitments to packet data fields store path.
func PacketAcknowledgementPrefixPath(sourceChain, destinationChain string) string {
	return fmt.Sprintf("%s/%s/%s", KeyPacketAckPrefix, packetPath(sourceChain, destinationChain), KeySequencePrefix)
}

// PacketReceiptPath defines the packet receipt store path
func PacketReceiptPath(sourceChain, destinationChain string, sequence uint64) string {
	return fmt.Sprintf("%s/%d", PacketReceiptPrefixPath(sourceChain, destinationChain), sequence)
}

// PacketReceiptKey returns the store key of under which a packet
// receipt is stored
func PacketReceiptKey(sourceChain, destinationChain string, sequence uint64) []byte {
	return []byte(PacketReceiptPath(sourceChain, destinationChain, sequence))
}

// PacketReceiptKey returns the store key of under which a packet
// receipt is stored
func PacketReceiptPrefixPath(sourceChain, destinationChain string) string {
	return fmt.Sprintf("%s/%s/%s", KeyPacketReceiptPrefix, packetPath(sourceChain, destinationChain), KeySequencePrefix)
}

func packetPath(sourceChain, destinationChain string) string {
	return fmt.Sprintf("%s/%s", sourceChain, destinationChain)
}

// CleanPacketCommitmentKey returns the store key of under which a clean packet commitment
// is stored
func CleanPacketCommitmentKey(sourceChain, destinationChain string) []byte {
	return []byte(CleanPacketCommitmentPath(sourceChain, destinationChain))
}

// CleanPacketCommitmentPrefixPath defines the prefix for commitments to packet data fields store path.
func CleanPacketCommitmentPath(sourceChain, destinationChain string) string {
	return fmt.Sprintf("%s/%s", KeyCleanPacketCommitmentPrefix, packetPath(sourceChain, destinationChain))
}
