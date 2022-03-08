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
		Bsc        Eth        `mapstructure:"bsc"`
		ChainType  string     `mapstructure:"chain_type"`
		Enabled    bool       `mapstructure:"enabled"`
	}
	// Eth config============================================================
	Eth struct {
		URI                   string       `mapstructure:"uri"`
		ChainID               uint64       `mapstructure:"chain_id"`
		ChainName             string       `mapstructure:"chain_name"`
		Contracts             EthContracts `mapstructure:"eth_contracts"`
		UpdateClientFrequency uint64       `mapstructure:"update_client_frequency"`
		GasLimit              uint64       `mapstructure:"gas_limit"`
		MaxGasPrice           uint64       `mapstructure:"max_gas_price"`
		CommentSlot           int64        `mapstructure:"comment_slot"`
		TipCoefficient        float64      `mapstructure:"tip_coefficient"`
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
	// Tendermint =====================================================================
	Tendermint struct {
		ChainName string   `mapstructure:"chain_name"`
		ChainID   string   `mapstructure:"chain_id"`
		RPCAddr   string   `mapstructure:"rpc_addr"`
		GrpcAddr  string   `mapstructure:"grpc_addr"`
		Gas       uint64   `mapstructure:"gas"`
		Key       ChainKey `mapstructure:"key"`
		Fee       Fee      `mapstructure:"fee"`
		Algo      string   `mapstructure:"algo"`

		RequestTimeout        uint   `mapstructure:"request_timeout"`
		UpdateClientFrequency uint64 `mapstructure:"update_client_frequency"`

		Allows             []Allow `mapstructure:"allows"`
		CleanPacketEnabled bool    `mapstructure:"clean_packet_enabled"`
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

	Allow struct {
		ContractAddr string   `mapstructure:"contract_addr"`
		Senders      []string `mapstructure:"senders"`
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
