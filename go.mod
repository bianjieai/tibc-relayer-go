module github.com/bianjieai/tibc-relayer-go

go 1.15

require (
	github.com/go-kit/kit v0.11.0
	github.com/irisnet/core-sdk-go v0.0.0-20210729072452-06544f6270f3
	github.com/pelletier/go-toml v1.9.3
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
