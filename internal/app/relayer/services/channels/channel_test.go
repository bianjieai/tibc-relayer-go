package channels

import (
	"testing"

	"github.com/irisnet/core-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	corestore "github.com/irisnet/core-sdk-go/types/store"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
)

func TestChannel_UpdateClient(t *testing.T) {
	chainNameA := "testCreateClientA"
	updateClientFrequencyA := 10
	cfgA := repostitory.NewTerndermintConfig()
	cfgA.ChainID = "testB"
	cfgA.GrpcAddr = "127.0.0.1:9090"
	cfgA.RPCAddr = "tcp://127.0.0.1:26657"
	cfgA.BaseTx = coretypes.BaseTx{
		From:               "nodeA",
		Gas:                1000000,
		Memo:               "TEST",
		Fee:                coretypes.NewDecCoins(coretypes.NewDecCoin("stake", coretypes.NewInt(100))),
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	cfgA.Options = []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(corestore.NewMemory(nil))),
		coretypes.TimeoutOption(10),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(0),
		coretypes.CachedOption(true),
	}
	sourceChian, err := repostitory.NewTendermintClient(
		constant.Tendermint,
		chainNameA,
		uint64(updateClientFrequencyA),
		cfgA,
	)
	if err != nil {
		t.Fatal(err)
	}

	chainNameB := "testCreateClientB"
	updateClientFrequencyB := 10
	cfgB := repostitory.NewTerndermintConfig()
	cfgB.ChainID = "testB"
	cfgB.GrpcAddr = "127.0.0.1:9091"
	cfgB.RPCAddr = "tcp://127.0.0.1:36657"
	cfgB.BaseTx = coretypes.BaseTx{
		From:               "nodeB",
		Gas:                1000000,
		Fee:                coretypes.NewDecCoins(coretypes.NewDecCoin("stake", coretypes.NewInt(100))),
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           "12345678",
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	cfgB.Options = []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(corestore.NewMemory(nil))),
		coretypes.TimeoutOption(10),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(0),
		coretypes.CachedOption(true),
	}
	destChian, err := repostitory.NewTendermintClient(
		constant.Tendermint,
		chainNameB,
		uint64(updateClientFrequencyB),
		cfgB,
	)

	if err != nil {
		return
	}

	channel := NewChannel(sourceChian, destChian, 4)
	err = channel.UpdateClient()
	if err != nil {
		t.Fatal(err)
	}
}
