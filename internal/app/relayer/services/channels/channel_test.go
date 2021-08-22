package channels

import (
	"testing"

	"github.com/irisnet/core-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	corestore "github.com/irisnet/core-sdk-go/types/store"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
)

func BSNChainClient() (*repostitory.Tendermint, error) {
	chainNameA := "testCreateClientA"
	updateClientFrequencyA := 10

	cfgA := repostitory.NewTerndermintConfig()
	cfgA.ChainID = "testA"
	cfgA.GrpcAddr = "localhost:9090"
	cfgA.RPCAddr = "tcp://localhost:26657"
	cfgA.Name = "chainANode0"
	cfgA.Password = "12345678"
	cfgA.BaseTx = coretypes.BaseTx{
		From:               cfgA.Name,
		Gas:                1000000,
		Memo:               "TEST",
		Fee:                coretypes.NewDecCoins(coretypes.NewDecCoin("stake", coretypes.NewInt(100))),
		Mode:               types.Commit,
		Password:           cfgA.Password,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	cfgA.PrivKeyArmor = `-----BEGIN TENDERMINT PRIVATE KEY-----
salt: D80938E846B69BC1E77BDF9E90476FB9
type: secp256k1
kdf: bcrypt

tdQjgq3BVr2c2J68eKttmNkA60m5x4FmQ1r1frioy0xBg7u+aIvu2X7n8z72jIkC
pfyksRgCIOZWgCoGctqCJAZXmEUBjgKuKgepppI=
=1GJn
-----END TENDERMINT PRIVATE KEY-----`
	cfgA.Options = []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(nil)),
		coretypes.TimeoutOption(10),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(0),
		coretypes.CachedOption(true),
	}

	return repostitory.NewTendermintClient(
		constant.Tendermint,
		chainNameA,
		uint64(updateClientFrequencyA),
		cfgA,
	)

}

func WenchangChainClient() (*repostitory.Tendermint, error) {
	chainNameB := "testCreateClientB"
	updateClientFrequencyB := 10
	cfgB := repostitory.NewTerndermintConfig()
	cfgB.ChainID = "testB"
	cfgB.GrpcAddr = "localhost:9091"
	cfgB.RPCAddr = "tcp://localhost:36657"
	cfgB.Name = "chainBNode0"
	cfgB.Password = "12345678"
	cfgB.BaseTx = coretypes.BaseTx{
		From:               cfgB.Name,
		Gas:                1000000,
		Memo:               "TEST",
		Fee:                coretypes.NewDecCoins(coretypes.NewDecCoin("stake", coretypes.NewInt(100))),
		Mode:               types.Commit,
		Password:           cfgB.Password,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	cfgB.PrivKeyArmor = `-----BEGIN TENDERMINT PRIVATE KEY-----
type: secp256k1
kdf: bcrypt
salt: AA44D1BBDA9BD3D023E493B0A1808EBE

5N103OwPijbmcjoSZwv6UMA9NUe/3cmemmE2Z+rhfd2NgwF5mp02e6Yt3JuqTBs9
CXxN7kDfbJk8AIL+6/qZgTxhQSRSoVpLVvaZ1BY=
=2Z08
-----END TENDERMINT PRIVATE KEY-----`

	cfgB.Options = []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(nil)),
		coretypes.TimeoutOption(10),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(0),
		coretypes.CachedOption(true),
	}
	return repostitory.NewTendermintClient(
		constant.Tendermint,
		chainNameB,
		uint64(updateClientFrequencyB),
		cfgB,
	)
}

func IrisHubChainClient() (*repostitory.Tendermint, error) {
	chainNameC := "testCreateClientC"
	updateClientFrequencyC := 10
	cfgC := repostitory.NewTerndermintConfig()
	cfgC.ChainID = "testC"
	cfgC.GrpcAddr = "localhost:9092"
	cfgC.RPCAddr = "tcp://localhost:46657"
	cfgC.Name = "chainCNode0"
	cfgC.Password = "12345678"
	cfgC.BaseTx = coretypes.BaseTx{
		From:               cfgC.Name,
		Gas:                1000000,
		Memo:               "TEST",
		Fee:                coretypes.NewDecCoins(coretypes.NewDecCoin("stake", coretypes.NewInt(100))),
		Mode:               types.Commit,
		Password:           cfgC.Password,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	cfgC.PrivKeyArmor = `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 36CB94F06728DECFFF5D6441C0DAE659
type: secp256k1

QzpMGWNySlO/IlntQzC7jz5dv6nqDOk0r/DW/oXuYranc6LvHoYad+F4otu/nCZR
z+4cCoBjltkT7ZupVRh4oFVe/bdWCRflOg2zeRA=
=uwZ4
-----END TENDERMINT PRIVATE KEY-----`

	cfgC.Options = []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(nil)),
		coretypes.TimeoutOption(10),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(0),
		coretypes.CachedOption(true),
	}
	return repostitory.NewTendermintClient(
		constant.Tendermint,
		chainNameC,
		uint64(updateClientFrequencyC),
		cfgC,
	)
}

func TestChannel_UpdateClient(t *testing.T) {
	// wenchang -> bsn hub
	wenchangChian, err := WenchangChainClient()
	if err != nil {
		t.Fatal(err)
	}

	bsnHubChain, err := BSNChainClient()

	if err != nil {
		return
	}

	channel := NewChannel(wenchangChian, bsnHubChain, 4)
	err = channel.UpdateClient()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannel_Relay(t *testing.T) {
	// wenchang -> bsn hub
	wenchangChian, err := WenchangChainClient()
	if err != nil {
		t.Fatal(err)
	}

	bsnHubChain, err := BSNChainClient()
	channel := NewChannel(wenchangChian, bsnHubChain, 430)
	err = channel.Relay()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannel_RelaySourceToHub(t *testing.T) {
	// Iris Hub -> BSN Hub
	irisHubChian, err := IrisHubChainClient() // b
	if err != nil {
		t.Fatal(err)
	}

	bsnHubChain, err := BSNChainClient() // a
	if err != nil {
		t.Fatal(err)
	}

	channel := NewChannel(irisHubChian, bsnHubChain, 4768)
	err = channel.Relay()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannel_RelayHubToSource(t *testing.T) {
	// BSN Hub -> wenchang
	sourceChian, err := WenchangChainClient() // b
	if err != nil {
		t.Fatal(err)
	}
	hubChain, err := BSNChainClient() // a
	if err != nil {
		t.Fatal(err)
	}

	channel := NewChannel(hubChain, sourceChian, 3653)
	err = channel.Relay()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannel_AckRelaySourceToHub(t *testing.T) {
	// Wenchang -> BSN Hub
	wenchangChian, err := WenchangChainClient() // b
	if err != nil {
		t.Fatal(err)
	}

	bsnHubChain, err := BSNChainClient() // a
	channel := NewChannel(wenchangChian, bsnHubChain, 4882)
	err = channel.Relay()
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannel_AckRelayHubToDest(t *testing.T) {
	// BSN Hub -> Iris Hub
	irisHubChian, err := IrisHubChainClient() // b
	if err != nil {
		t.Fatal(err)
	}

	bsnHubChain, err := BSNChainClient() // a
	channel := NewChannel(bsnHubChain, irisHubChian, 4501)
	err = channel.Relay()
	if err != nil {
		t.Fatal(err)
	}
}
