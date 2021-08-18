package repostitory

import tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"

type GetBlockHeaderReq struct {
	LatestHeight  uint64
	TrustedHeight uint64

	TrustedTendermintValidtorSet *tenderminttypes.ValidatorSet
}

type QueryLightClientValidatorResp struct {
	ChainType string

	TendermintValidtorSet *tenderminttypes.ValidatorSet
}
