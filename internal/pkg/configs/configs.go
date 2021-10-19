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
		Eth        Eth        `mapstructure:"eth"`
	}
	// eth config============================================================
	Eth struct {
		URI                   string       `mapstructure:"uri"`
		ChainID               uint64       `mapstructure:"chain_id"`
		ChainName             string       `mapstructure:"chain_name"`
		Contracts             EthContracts `mapstructure:"eth_contracts"`
		UpdateClientFrequency uint64       `mapstructure:"update_client_frequency"`
		GasLimit              uint64       `mapstructure:"gas_limit"`
		GasPrice              uint64       `mapstructure:"gas_price"`
		CommentSlot           int64        `mapstructure:"comment_slot"`
	}

	EthContracts struct {
		Packet      EthContractCfg `mapstructure:"packet"`
		AckPacket   EthContractCfg `mapstructure:"ack_packet"`
		CleanPacket EthContractCfg `mapstructure:"clean_packet"`
		Client      EthContractCfg `mapstructure:"client"`
	}

	EthContractCfg struct {
		Addr       string `mapstructure:"addr"`
		Topic      string `mapstructure:"topic"`
		OptPrivKey string `mapstructure:"opt_priv_key"`
	}
	// =====================================================================
	// Tendermit config=====================================================
	Tendermint struct {
		ChainName string   `mapstructure:"chain_name"`
		ChainID   string   `mapstructure:"chain_id"`
		RPCAddr   string   `mapstructure:"rpc_addr"`
		GrpcAddr  string   `mapstructure:"grpc_addr"`
		Gas       uint64   `mapstructure:"gas"`
		Key       ChainKey `mapstructure:"key"`
		Fee       Fee      `mapstructure:"fee"`
		Algo      string   `mapstructure:"algo"`

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
	// =====================================================================

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
