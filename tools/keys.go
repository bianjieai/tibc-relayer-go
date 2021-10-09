package tools

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ProofKeyConstructor struct {
	sourceChain string
	destChain   string
	sequence    uint64
}

func NewProofKeyConstructor(sourceChain string, destChain string, sequence uint64) ProofKeyConstructor {
	return ProofKeyConstructor{
		sourceChain: sourceChain,
		destChain:   destChain,
		sequence:    sequence,
	}
}

func (k ProofKeyConstructor) GetPacketCommitmentProofKey(slot int64) []byte {
	hash := crypto.Keccak256Hash(
		PacketCommitmentKey(k.sourceChain, k.destChain, k.sequence),
		common.LeftPadBytes(big.NewInt(slot).Bytes(), 32),
	)
	return hash.Bytes()
}

func (k ProofKeyConstructor) GetAckProofKey(slot int64) []byte {
	hash := crypto.Keccak256Hash(
		PacketAcknowledgementKey(k.sourceChain, k.destChain, k.sequence),
		common.LeftPadBytes(big.NewInt(slot).Bytes(), 32),
	)
	return hash.Bytes()
}

func (k ProofKeyConstructor) GetCleanPacketCommitmentProofKey(slot int64) []byte {
	hash := crypto.Keccak256Hash(
		CleanPacketCommitmentKey(k.sourceChain, k.destChain),
		common.LeftPadBytes(big.NewInt(slot).Bytes(), 32),
	)
	return hash.Bytes()
}
