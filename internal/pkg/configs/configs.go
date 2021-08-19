package configs

type (
	Config struct {
		App   App   `mapstructure:"app"`
		Chain Chain `mapstructure:"chain"`
	}

	Chain struct {
		Source ChainCfg `mapstructure:"source"`
		Dest   ChainCfg `mapstructure:"dest"`
	}

	ChainCfg struct {
		Cache      Cache      `mapstructure:"cache"`
		Tendermint Tendermint `mapstructure:"tendermint"`
	}

	Tendermint struct {
		ChainName string   `mapstructure:"chain_name"`
		ChainID   string   `mapstructure:"chain_id"`
		RPCAddr   string   `mapstructure:"rpc_addr"`
		GrpcAddr  string   `mapstructure:"grpc_addr"`
		Gas       uint64   `mapstructure:"gas"`
		Key       ChainKey `mapstructure:"key"`
		Fee       Fee      `mapstructure:"fee"`

		UpdateClientFrequency uint64 `mapstructure:"update_client_frequency"`
	}

	Fee struct {
		Denom  string `mapstructure:"denom"`
		Amount int64  `mapstructure:"amount"`
	}

	ChainKey struct {
		Name         string `mapstructure:"name"`
		Password     string `mapstructure:"password"`
		PrivKeyArmor string `mapstructure:"priv_key_armor"`
	}

	App struct {
		MetricAddr   string   `mapstructure:"metric_addr"`
		Env          string   `mapstructure:"env"`
		LogLevel     string   `mapstructure:"log_level"`
		ChannelTypes []string `mapstructure:"channel_types"`
	}

	Cache struct {
		Filename    string `mapstructure:"filename"`
		StartHeight uint64 `mapstructure:"start_height"`
	}
)

func NewConfig() *Config {
	return &Config{}
}
