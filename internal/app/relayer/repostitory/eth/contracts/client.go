// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FractionData is an auto generated low-level Go binding around an user-defined struct.
type FractionData struct {
	Numerator   uint64
	Denominator uint64
}

// MerklePrefixData is an auto generated low-level Go binding around an user-defined struct.
type MerklePrefixData struct {
	KeyPrefix []byte
}

// TimestampData is an auto generated low-level Go binding around an user-defined struct.
type TimestampData struct {
	Secs  int64
	Nanos int64
}

// ClientMetaData contains all meta data concerning the Client contract.
var ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clientManagerAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"headerBz\",\"type\":\"bytes\"}],\"name\":\"checkHeaderAndUpdateState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clientState\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"numerator\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"denominator\",\"type\":\"uint64\"}],\"internalType\":\"structFraction.Data\",\"name\":\"trust_level\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"trusting_period\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbonding_period\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"max_clock_drift\",\"type\":\"int64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"latest_height\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"merkle_prefix\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"time_delay\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"consensusStates\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"secs\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"nanos\",\"type\":\"int64\"}],\"internalType\":\"structTimestamp.Data\",\"name\":\"timestamp\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"next_validators_hash\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestHeight\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"clientStateBz\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBz\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"clientStateBz\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBz\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"verifyPacketAcknowledgement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"verifyPacketCleanCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"commitmentBytes\",\"type\":\"bytes\"}],\"name\":\"verifyPacketCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use ClientMetaData.ABI instead.
var ClientABI = ClientMetaData.ABI

// Client is an auto generated Go binding around an Ethereum contract.
type Client struct {
	ClientCaller     // Read-only binding to the contract
	ClientTransactor // Write-only binding to the contract
	ClientFilterer   // Log filterer for contract events
}

// ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClientSession struct {
	Contract     *Client           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClientCallerSession struct {
	Contract *ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClientTransactorSession struct {
	Contract     *ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClientRaw struct {
	Contract *Client // Generic contract binding to access the raw methods on
}

// ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClientCallerRaw struct {
	Contract *ClientCaller // Generic read-only contract binding to access the raw methods on
}

// ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClientTransactorRaw struct {
	Contract *ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClient creates a new instance of Client, bound to a specific deployed contract.
func NewClient(address common.Address, backend bind.ContractBackend) (*Client, error) {
	contract, err := bindClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Client{ClientCaller: ClientCaller{contract: contract}, ClientTransactor: ClientTransactor{contract: contract}, ClientFilterer: ClientFilterer{contract: contract}}, nil
}

// NewClientCaller creates a new read-only instance of Client, bound to a specific deployed contract.
func NewClientCaller(address common.Address, caller bind.ContractCaller) (*ClientCaller, error) {
	contract, err := bindClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClientCaller{contract: contract}, nil
}

// NewClientTransactor creates a new write-only instance of Client, bound to a specific deployed contract.
func NewClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ClientTransactor, error) {
	contract, err := bindClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClientTransactor{contract: contract}, nil
}

// NewClientFilterer creates a new log filterer instance of Client, bound to a specific deployed contract.
func NewClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ClientFilterer, error) {
	contract, err := bindClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClientFilterer{contract: contract}, nil
}

// bindClient binds a generic wrapper to an already deployed contract.
func bindClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Client *ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Client.Contract.ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Client *ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.Contract.ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Client *ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Client.Contract.ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Client *ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Client *ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Client *ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Client.Contract.contract.Transact(opts, method, params...)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(string chain_id, (uint64,uint64) trust_level, int64 trusting_period, int64 unbonding_period, int64 max_clock_drift, (uint64,uint64) latest_height, (bytes) merkle_prefix, uint64 time_delay)
func (_Client *ClientCaller) ClientState(opts *bind.CallOpts) (struct {
	ChainId         string
	TrustLevel      FractionData
	TrustingPeriod  int64
	UnbondingPeriod int64
	MaxClockDrift   int64
	LatestHeight    HeightData
	MerklePrefix    MerklePrefixData
	TimeDelay       uint64
}, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "clientState")

	outstruct := new(struct {
		ChainId         string
		TrustLevel      FractionData
		TrustingPeriod  int64
		UnbondingPeriod int64
		MaxClockDrift   int64
		LatestHeight    HeightData
		MerklePrefix    MerklePrefixData
		TimeDelay       uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TrustLevel = *abi.ConvertType(out[1], new(FractionData)).(*FractionData)
	outstruct.TrustingPeriod = *abi.ConvertType(out[2], new(int64)).(*int64)
	outstruct.UnbondingPeriod = *abi.ConvertType(out[3], new(int64)).(*int64)
	outstruct.MaxClockDrift = *abi.ConvertType(out[4], new(int64)).(*int64)
	outstruct.LatestHeight = *abi.ConvertType(out[5], new(HeightData)).(*HeightData)
	outstruct.MerklePrefix = *abi.ConvertType(out[6], new(MerklePrefixData)).(*MerklePrefixData)
	outstruct.TimeDelay = *abi.ConvertType(out[7], new(uint64)).(*uint64)

	return *outstruct, err

}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(string chain_id, (uint64,uint64) trust_level, int64 trusting_period, int64 unbonding_period, int64 max_clock_drift, (uint64,uint64) latest_height, (bytes) merkle_prefix, uint64 time_delay)
func (_Client *ClientSession) ClientState() (struct {
	ChainId         string
	TrustLevel      FractionData
	TrustingPeriod  int64
	UnbondingPeriod int64
	MaxClockDrift   int64
	LatestHeight    HeightData
	MerklePrefix    MerklePrefixData
	TimeDelay       uint64
}, error) {
	return _Client.Contract.ClientState(&_Client.CallOpts)
}

// ClientState is a free data retrieval call binding the contract method 0xbd3ce6b0.
//
// Solidity: function clientState() view returns(string chain_id, (uint64,uint64) trust_level, int64 trusting_period, int64 unbonding_period, int64 max_clock_drift, (uint64,uint64) latest_height, (bytes) merkle_prefix, uint64 time_delay)
func (_Client *ClientCallerSession) ClientState() (struct {
	ChainId         string
	TrustLevel      FractionData
	TrustingPeriod  int64
	UnbondingPeriod int64
	MaxClockDrift   int64
	LatestHeight    HeightData
	MerklePrefix    MerklePrefixData
	TimeDelay       uint64
}, error) {
	return _Client.Contract.ClientState(&_Client.CallOpts)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns((int64,int64) timestamp, bytes root, bytes next_validators_hash)
func (_Client *ClientCaller) ConsensusStates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp          TimestampData
	Root               []byte
	NextValidatorsHash []byte
}, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "consensusStates", arg0)

	outstruct := new(struct {
		Timestamp          TimestampData
		Root               []byte
		NextValidatorsHash []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(TimestampData)).(*TimestampData)
	outstruct.Root = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.NextValidatorsHash = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns((int64,int64) timestamp, bytes root, bytes next_validators_hash)
func (_Client *ClientSession) ConsensusStates(arg0 *big.Int) (struct {
	Timestamp          TimestampData
	Root               []byte
	NextValidatorsHash []byte
}, error) {
	return _Client.Contract.ConsensusStates(&_Client.CallOpts, arg0)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns((int64,int64) timestamp, bytes root, bytes next_validators_hash)
func (_Client *ClientCallerSession) ConsensusStates(arg0 *big.Int) (struct {
	Timestamp          TimestampData
	Root               []byte
	NextValidatorsHash []byte
}, error) {
	return _Client.Contract.ConsensusStates(&_Client.CallOpts, arg0)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns((uint64,uint64))
func (_Client *ClientCaller) GetLatestHeight(opts *bind.CallOpts) (HeightData, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "getLatestHeight")

	if err != nil {
		return *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)

	return out0, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns((uint64,uint64))
func (_Client *ClientSession) GetLatestHeight() (HeightData, error) {
	return _Client.Contract.GetLatestHeight(&_Client.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() view returns((uint64,uint64))
func (_Client *ClientCallerSession) GetLatestHeight() (HeightData, error) {
	return _Client.Contract.GetLatestHeight(&_Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Client *ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Client *ClientSession) Owner() (common.Address, error) {
	return _Client.Contract.Owner(&_Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Client *ClientCallerSession) Owner() (common.Address, error) {
	return _Client.Contract.Owner(&_Client.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(int8)
func (_Client *ClientCaller) Status(opts *bind.CallOpts) (int8, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(int8)
func (_Client *ClientSession) Status() (int8, error) {
	return _Client.Contract.Status(&_Client.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(int8)
func (_Client *ClientCallerSession) Status() (int8, error) {
	return _Client.Contract.Status(&_Client.CallOpts)
}

// CheckHeaderAndUpdateState is a paid mutator transaction binding the contract method 0xb47a619a.
//
// Solidity: function checkHeaderAndUpdateState(bytes headerBz) returns()
func (_Client *ClientTransactor) CheckHeaderAndUpdateState(opts *bind.TransactOpts, headerBz []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "checkHeaderAndUpdateState", headerBz)
}

// CheckHeaderAndUpdateState is a paid mutator transaction binding the contract method 0xb47a619a.
//
// Solidity: function checkHeaderAndUpdateState(bytes headerBz) returns()
func (_Client *ClientSession) CheckHeaderAndUpdateState(headerBz []byte) (*types.Transaction, error) {
	return _Client.Contract.CheckHeaderAndUpdateState(&_Client.TransactOpts, headerBz)
}

// CheckHeaderAndUpdateState is a paid mutator transaction binding the contract method 0xb47a619a.
//
// Solidity: function checkHeaderAndUpdateState(bytes headerBz) returns()
func (_Client *ClientTransactorSession) CheckHeaderAndUpdateState(headerBz []byte) (*types.Transaction, error) {
	return _Client.Contract.CheckHeaderAndUpdateState(&_Client.TransactOpts, headerBz)
}

// Initialize is a paid mutator transaction binding the contract method 0x1af19f77.
//
// Solidity: function initialize(bytes clientStateBz, bytes consensusStateBz) returns()
func (_Client *ClientTransactor) Initialize(opts *bind.TransactOpts, clientStateBz []byte, consensusStateBz []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "initialize", clientStateBz, consensusStateBz)
}

// Initialize is a paid mutator transaction binding the contract method 0x1af19f77.
//
// Solidity: function initialize(bytes clientStateBz, bytes consensusStateBz) returns()
func (_Client *ClientSession) Initialize(clientStateBz []byte, consensusStateBz []byte) (*types.Transaction, error) {
	return _Client.Contract.Initialize(&_Client.TransactOpts, clientStateBz, consensusStateBz)
}

// Initialize is a paid mutator transaction binding the contract method 0x1af19f77.
//
// Solidity: function initialize(bytes clientStateBz, bytes consensusStateBz) returns()
func (_Client *ClientTransactorSession) Initialize(clientStateBz []byte, consensusStateBz []byte) (*types.Transaction, error) {
	return _Client.Contract.Initialize(&_Client.TransactOpts, clientStateBz, consensusStateBz)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Client *ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Client *ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _Client.Contract.RenounceOwnership(&_Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Client *ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Client.Contract.RenounceOwnership(&_Client.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Client *ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Client *ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Client.Contract.TransferOwnership(&_Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Client *ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Client.Contract.TransferOwnership(&_Client.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes clientStateBz, bytes consensusStateBz) returns()
func (_Client *ClientTransactor) Upgrade(opts *bind.TransactOpts, clientStateBz []byte, consensusStateBz []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "upgrade", clientStateBz, consensusStateBz)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes clientStateBz, bytes consensusStateBz) returns()
func (_Client *ClientSession) Upgrade(clientStateBz []byte, consensusStateBz []byte) (*types.Transaction, error) {
	return _Client.Contract.Upgrade(&_Client.TransactOpts, clientStateBz, consensusStateBz)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes clientStateBz, bytes consensusStateBz) returns()
func (_Client *ClientTransactorSession) Upgrade(clientStateBz []byte, consensusStateBz []byte) (*types.Transaction, error) {
	return _Client.Contract.Upgrade(&_Client.TransactOpts, clientStateBz, consensusStateBz)
}

// VerifyPacketAcknowledgement is a paid mutator transaction binding the contract method 0xc7ada807.
//
// Solidity: function verifyPacketAcknowledgement((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence, bytes acknowledgement) returns()
func (_Client *ClientTransactor) VerifyPacketAcknowledgement(opts *bind.TransactOpts, height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "verifyPacketAcknowledgement", height, proof, sourceChain, destChain, sequence, acknowledgement)
}

// VerifyPacketAcknowledgement is a paid mutator transaction binding the contract method 0xc7ada807.
//
// Solidity: function verifyPacketAcknowledgement((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence, bytes acknowledgement) returns()
func (_Client *ClientSession) VerifyPacketAcknowledgement(height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Client.Contract.VerifyPacketAcknowledgement(&_Client.TransactOpts, height, proof, sourceChain, destChain, sequence, acknowledgement)
}

// VerifyPacketAcknowledgement is a paid mutator transaction binding the contract method 0xc7ada807.
//
// Solidity: function verifyPacketAcknowledgement((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence, bytes acknowledgement) returns()
func (_Client *ClientTransactorSession) VerifyPacketAcknowledgement(height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64, acknowledgement []byte) (*types.Transaction, error) {
	return _Client.Contract.VerifyPacketAcknowledgement(&_Client.TransactOpts, height, proof, sourceChain, destChain, sequence, acknowledgement)
}

// VerifyPacketCleanCommitment is a paid mutator transaction binding the contract method 0x06023111.
//
// Solidity: function verifyPacketCleanCommitment((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence) returns()
func (_Client *ClientTransactor) VerifyPacketCleanCommitment(opts *bind.TransactOpts, height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "verifyPacketCleanCommitment", height, proof, sourceChain, destChain, sequence)
}

// VerifyPacketCleanCommitment is a paid mutator transaction binding the contract method 0x06023111.
//
// Solidity: function verifyPacketCleanCommitment((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence) returns()
func (_Client *ClientSession) VerifyPacketCleanCommitment(height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64) (*types.Transaction, error) {
	return _Client.Contract.VerifyPacketCleanCommitment(&_Client.TransactOpts, height, proof, sourceChain, destChain, sequence)
}

// VerifyPacketCleanCommitment is a paid mutator transaction binding the contract method 0x06023111.
//
// Solidity: function verifyPacketCleanCommitment((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence) returns()
func (_Client *ClientTransactorSession) VerifyPacketCleanCommitment(height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64) (*types.Transaction, error) {
	return _Client.Contract.VerifyPacketCleanCommitment(&_Client.TransactOpts, height, proof, sourceChain, destChain, sequence)
}

// VerifyPacketCommitment is a paid mutator transaction binding the contract method 0x52904aac.
//
// Solidity: function verifyPacketCommitment((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence, bytes commitmentBytes) returns()
func (_Client *ClientTransactor) VerifyPacketCommitment(opts *bind.TransactOpts, height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64, commitmentBytes []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "verifyPacketCommitment", height, proof, sourceChain, destChain, sequence, commitmentBytes)
}

// VerifyPacketCommitment is a paid mutator transaction binding the contract method 0x52904aac.
//
// Solidity: function verifyPacketCommitment((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence, bytes commitmentBytes) returns()
func (_Client *ClientSession) VerifyPacketCommitment(height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64, commitmentBytes []byte) (*types.Transaction, error) {
	return _Client.Contract.VerifyPacketCommitment(&_Client.TransactOpts, height, proof, sourceChain, destChain, sequence, commitmentBytes)
}

// VerifyPacketCommitment is a paid mutator transaction binding the contract method 0x52904aac.
//
// Solidity: function verifyPacketCommitment((uint64,uint64) height, bytes proof, string sourceChain, string destChain, uint64 sequence, bytes commitmentBytes) returns()
func (_Client *ClientTransactorSession) VerifyPacketCommitment(height HeightData, proof []byte, sourceChain string, destChain string, sequence uint64, commitmentBytes []byte) (*types.Transaction, error) {
	return _Client.Contract.VerifyPacketCommitment(&_Client.TransactOpts, height, proof, sourceChain, destChain, sequence, commitmentBytes)
}

// ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Client contract.
type ClientOwnershipTransferredIterator struct {
	Event *ClientOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClientOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ClientOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClientOwnershipTransferred represents a OwnershipTransferred event raised by the Client contract.
type ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Client *ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClientOwnershipTransferredIterator{contract: _Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Client *ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClientOwnershipTransferred)
				if err := _Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Client *ClientFilterer) ParseOwnershipTransferred(log types.Log) (*ClientOwnershipTransferred, error) {
	event := new(ClientOwnershipTransferred)
	if err := _Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
