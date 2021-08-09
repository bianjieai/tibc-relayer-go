package configs

type (
	Config struct {
		App   App   `mapstructure:"app"`
		Chain Chain `mapstructure:"chain"`
	}

	Chain struct {
		IrisHub IRITA `mapstructure:"iris_hub"`
	}

	IRITA struct {
		GrpcAddr   string `mapstructure:"grpc_addr"`
		RpcAddr    string `mapstructure:"rpc_addr"`
		ModuleName string `mapstructure:"module_name"`
		ChainID    string `mapstructure:"chain_id"`
		ClientID   string `mapstructure:"client_id"`
		Timeout    uint   `mapstructure:"timeout"`
		Signer     string `mapstructure:"signer"`
		Cache      Cache  `mapstructure:"cache"`
	}

	App struct {
		MetricAddr string `mapstructure:"metric_addr"`
		Env        string `mapstructure:"env"`
		LogLevel   string `mapstructure:"log_level"`
	}

	Cache struct {
		Dir         string `mapstructure:"dir"`
		Filename    string `mapstructure:"filename"`
		StartHeight uint64 `mapstructure:"start_height"`
	}
)

func NewConfig() *Config {
	return &Config{}
}
