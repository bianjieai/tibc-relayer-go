module github.com/bianjieai/tibc-relayer-go

go 1.15

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/golang/snappy v0.0.2 // indirect
	github.com/irisnet/core-sdk-go v0.0.0-20210729072452-06544f6270f3
	github.com/jasonlvhit/gocron v0.0.1
	github.com/prometheus/common v0.15.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/objx v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
