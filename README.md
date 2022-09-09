# tibc-relayer-go

## Introduction

Golang Relayer for Terse IBC

## Support

- Tendermint and Tendermint (done)
- Tendermint and ETH (done)
- Tendermint and BSC (done)

## Configs

### Tendermint and Tendermint example

```toml
[app]

env = "dev"
log_level = "debug"
metric_addr = "0.0.0.0:8084"
channel_types = ["tendermint_and_tendermint"]


# wenchangchain  -> bsn

# bsn

[chain]

[chain.dest]

chain_type = "tendermint"
enabled = true # Whether to enable relay

[chain.dest.cache]
filename = "dest-chain.json"
start_height = 5

[chain.dest.tendermint]
chain_id = "dest-testnet"
chain_name = "dest-testnet-mainnet"
gas = 100000000
grpc_addr = "dest_chain_host:9090" # eg. 127.0.0.1:9090
rpc_addr = "http://dest_chain_host:26657" # eg. http://127.0.0.1:26657
update_client_frequency = 1 # update client frequency, unit hour
request_timeout = 60
clean_packet_enabled = false

[chain.dest.tendermint.fee]
denom = "uiris"
amount = 100

[chain.dest.tendermint.key]
name = "your_dest_chain_relayer_name"
password = "your_dest_chain_relayer_pwd"
priv_key_armor = "your_dest_chain_relayer_priv_key"


# iris-hub

[chain.source]

chain_type = "tendermint"
enabled = true # Whether to enable relay

[chain.source.cache]
filename = "source_chain.json"
start_height = 5

[chain.source.tendermint]
chain_id = "source-testnet"
chain_name = "source-testnet-mainnet"
gas = 100000000
grpc_addr = "source_chain_host:9090" # eg. 127.0.0.1:9090
rpc_addr = "http://source_chain_host:26657" # eg. http://127.0.0.1:26657
update_client_frequency = 1 # update client frequency, unit hour
request_timeout = 60
clean_packet_enabled = false
algo = "sm2" # If there is no special annotation, use secp256k1

[chain.source.tendermint.fee]
denom = "uirita"
amount = 100

[chain.source.tendermint.key]
name = "your_source_chain_relayer_name"
password = "your_source_chain_relayer_pwd"
priv_key_armor = "your_source_chain_relayer_priv_key"

```

### Tendermint and ETH example

```toml
[app]

env = "prod"
log_level = "info"
metric_addr = "0.0.0.0:8083"
channel_types = ["tendermint_and_eth"]

[chain]

[chain.dest]
chain_type = "bsc"
enabled = false

[chain.dest.cache]
filename = "dest_chain.json"
start_height = 17206800

[chain.dest.eth]
chain_id = 97
chain_name = "eth-mainnet" # The name registered when deploying the contract
update_client_frequency = 2 
uri = ""
gas_limit = 2000000
max_gas_price =  150000000000
comment_slot = 104
tip_coefficient = 0.1

[chain.dest.eth.eth_contracts]

[chain.dest.eth.eth_contracts.packet]
addr = "" # packet contract addr
topic = "PacketSent((uint64,string,string,string,string,bytes))"
opt_priv_key = "" # eth relayer priv key

[chain.dest.eth.eth_contracts.ack_packet]
addr = "" # packet contract addr
topic = "AckWritten((uint64,string,string,string,string,bytes),bytes)"
opt_priv_key = "" # eth relayer priv key

[chain.dest.eth.eth_contracts.clean_packet]
addr = "" # packet contract addr地址
topic = "CleanPacketSent((uint64,string,string,string))"
opt_priv_key = "" # eth relayer priv key

[chain.dest.eth.eth_contracts.client]
addr = "" # client manager contract addr
topic = ""
opt_priv_key = ""  # eth relayer priv key



# iris-hub

[chain.source]

chain_type = "tendermint"
enabled = true # Whether to enable relay

[chain.source.cache]
filename = "source_chain.json"
start_height = 5

[chain.source.tendermint]
chain_id = "source-testnet"
chain_name = "source-testnet-mainnet"
gas = 100000000
grpc_addr = "source_chain_host:9090" # eg. 127.0.0.1:9090
rpc_addr = "http://source_chain_host:26657" # eg. http://127.0.0.1:26657
update_client_frequency = 1 # update client frequency, unit hour
request_timeout = 60
clean_packet_enabled = false
algo = "sm2" # chain user algo

[chain.source.tendermint.fee]
denom = "uirita"
amount = 100

[chain.source.tendermint.key]
name = "your_source_chain_relayer_name"
password = "your_source_chain_relayer_pwd"
priv_key_armor = "your_source_chain_relayer_priv_key"

```


### Tendermint and BSC example

```toml
[app]

env = "prod"
log_level = "info"
metric_addr = "0.0.0.0:8083"
channel_types = ["tendermint_and_bsc"]

[chain]

[chain.dest]
chain_type = "bsc"
enabled = false

[chain.dest.cache]
filename = "dest_chain.json"
start_height = 17206800

[chain.dest.bsc]
chain_id = 97
chain_name = "bsc-mainnet" # The name registered when deploying the contract
update_client_frequency = 2 
uri = "https://data-seed-prebsc-1-s1.binance.org:8545"
gas_limit = 2000000
max_gas_price =  150000000000
comment_slot = 104
tip_coefficient = 0.1

[chain.dest.bsc.eth_contracts]

[chain.dest.bsc.eth_contracts.packet]
addr = "" # packet contract addr
topic = "PacketSent((uint64,string,string,string,string,bytes))"
opt_priv_key = "" # eth relayer priv key

[chain.dest.bsc.eth_contracts.ack_packet]
addr = "" # packet contract addr
topic = "AckWritten((uint64,string,string,string,string,bytes),bytes)"
opt_priv_key = "" # eth relayer priv key

[chain.dest.bsc.eth_contracts.clean_packet]
addr = "" # packet contract addr地址
topic = "CleanPacketSent((uint64,string,string,string))"
opt_priv_key = "" # eth relayer priv key

[chain.dest.bsc.eth_contracts.client]
addr = "" # client manager contract addr
topic = ""
opt_priv_key = ""  # eth relayer priv key

# iris-hub

[chain.source]

chain_type = "tendermint"
enabled = true # Whether to enable relay

[chain.source.cache]
filename = "source_chain.json"
start_height = 5

[chain.source.tendermint]
chain_id = "source-testnet"
chain_name = "source-testnet-mainnet"
gas = 100000000
grpc_addr = "source_chain_host:9090" # eg. 127.0.0.1:9090
rpc_addr = "http://source_chain_host:26657" # eg. http://127.0.0.1:26657
update_client_frequency = 1 # update client frequency, unit hour
request_timeout = 60
clean_packet_enabled = false
algo = "sm2" # chain user algo

[chain.source.tendermint.fee]
denom = "uirita"
amount = 100

[chain.source.tendermint.key]
name = "your_source_chain_relayer_name"
password = "your_source_chain_relayer_pwd"
priv_key_armor = "your_source_chain_relayer_priv_key"

```

### Tendermint and Ethermint example

```toml
[app]

env = "dev"
log_level = "info"
metric_addr = "0.0.0.0:8083"
channel_types = ["tendermint_and_ethermint"]

[chain]

[chain.dest]
chain_type = "ethermint"
enabled = true

[chain.dest.cache]
filename = "irita-b-evm-dest.json"
start_height = 100

[chain.dest.ethermint]
# comment
chain_name = "irita-b-evm-testnet"

# eth
eth_chain_id = 1223
update_client_frequency = 2 
uri = "http://localhost:8545"
gas_limit = 2000000
max_gas_price =  150000000000
comment_slot = 104
tip_coefficient = 0.1

# tendermint
tendermint_chain_id = "irita-b-testnet"
gas = 300000
grpc_addr = "127.0.0.1:9090" 
rpc_addr = "http://127.0.0.1:26657" 
request_timeout = 60

[chain.dest.ethermint.eth_contracts]

[chain.dest.ethermint.eth_contracts.packet]
addr = "" # packet 合约
topic = "PacketSent((uint64,string,string,string,string,bytes))"
opt_priv_key = "" # eth relayer 私钥

[chain.dest.ethermint.eth_contracts.ack_packet]
addr = "" # packet 合约
topic = "AckWritten((uint64,string,string,string,string,bytes),bytes)"
opt_priv_key = "" # eth relayer 私钥

[chain.dest.ethermint.eth_contracts.clean_packet]
addr = "" # packet 合约
topic = "CleanPacketSent((uint64,string,string,string))"
opt_priv_key = "" # eth relayer 私钥

[chain.dest.ethermint.eth_contracts.client]
addr = "" # client manager 合约
topic = ""
opt_priv_key = ""  # eth relayer 私钥

[chain.source]

chain_type = "tendermint"
enabled = true # Whether to enable relay

[chain.source.cache]
filename = "irita-a-testnet-source.json"
start_height = 100

[chain.source.tendermint]
chain_id = "irita-a-testnet"
chain_name = "irita-a-testnet"
gas = 300000
grpc_addr = "127.0.0.1:19090" 
rpc_addr = "http://127.0.0.1:36657" 
update_client_frequency = 1 
request_timeout = 60
clean_packet_enabled = false
algo="sm2"

[chain.source.tendermint.fee]
denom = "ugas"
amount = 100

[[chain.source.tendermint.allows]]
contract_addr = ""
senders = [""]

[chain.source.tendermint.key]
name = "relayereth-1"
password = "12345678"
priv_key_armor = ""

```