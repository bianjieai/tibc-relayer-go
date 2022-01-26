module github.com/bianjieai/tibc-relayer-go

go 1.15

require (
	github.com/bianjieai/tibc-sdk-go v0.0.0-20211028070139-cce81c8277a7
	github.com/ethereum/go-ethereum v1.10.8
	github.com/go-kit/kit v0.11.0
	github.com/irisnet/core-sdk-go v0.0.0-20211019075829-8bb6cca8d315
	github.com/irisnet/irismod-sdk-go/nft v0.0.0-20210810032454-3ae775c15f1e
	github.com/jasonlvhit/gocron v0.0.1
	github.com/pelletier/go-toml v1.9.3
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.11
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/irisnet/core-sdk-go => github.com/irisnet/core-sdk-go v0.0.0-20220126064947-34046bc65a06
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
