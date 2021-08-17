package configs

type (
	Config struct {
		App   App   `toml:"app";mapstructure:"app"`
		Chain Chain `toml:"chain";mapstructure:"chain"`
	}

	Chain struct {
		Source ChainCfg `toml:"source";mapstructure:"source"`
		Dest   ChainCfg `toml:"dest";mapstructure:"dest"`
	}

	ChainCfg struct {
		Cache      Cache      `toml:"cache";mapstructure:"cache"`
		Tendermint Tendermint `toml:"tendermint";mapstructure:"tendermint"`
	}

	Tendermint struct {
		ChainName string   `toml:"chain_name";mapstructure:"chain_name"`
		ChainID   string   `toml:"chain_id";mapstructure:"chain_id"`
		RPCAddr   string   `toml:"rpc_addr";mapstructure:"rpc_addr"`
		GrpcAddr  string   `toml:"grpc_addr";mapstructure:"grpc_addr"`
		Gas       uint64   `toml:"gas";mapstructure:"gas"`
		Key       ChainKey `toml:"key";mapstructure:"key"`

		UpdateClientFrequency uint64 `toml:"update_client_frequency";mapstructure:"update_client_frequency"`
	}

	ChainKey struct {
		Name     string `toml:"name";mapstructure:"name"`
		Password string `toml:"password";mapstructure:"password"`
		Signer   string `toml:"signer";mapstructure:"signer"`
		Path     string `toml:"path";mapstructure:"path"`
	}

	App struct {
		MetricAddr   string   `toml:"metric_addr";mapstructure:"metric_addr"`
		Env          string   `toml:"env";mapstructure:"env"`
		LogLevel     string   `toml:"log_level";mapstructure:"log_level"`
		ChannelTypes []string `toml:"channel_types";mapstructure:"channel_types"`
	}

	Cache struct {
		Filename    string `toml:"filename";mapstructure:"filename"`
		StartHeight uint64 `toml:"start_height";mapstructure:"start_height"`
	}
)

func NewConfig() *Config {
	return &Config{}
}
