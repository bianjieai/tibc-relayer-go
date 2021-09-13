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

// ClientMetaData contains all meta data concerning the Client contract.
var ClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"clients\",\"outputs\":[{\"internalType\":\"contractIClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"clientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"createClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"}],\"name\":\"getClient\",\"outputs\":[{\"internalType\":\"contractIClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"registerRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"updateClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"chainName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"}],\"name\":\"upgradeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Clients is a free data retrieval call binding the contract method 0x20ba1e9f.
//
// Solidity: function clients(string ) view returns(address)
func (_Client *ClientCaller) Clients(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "clients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Clients is a free data retrieval call binding the contract method 0x20ba1e9f.
//
// Solidity: function clients(string ) view returns(address)
func (_Client *ClientSession) Clients(arg0 string) (common.Address, error) {
	return _Client.Contract.Clients(&_Client.CallOpts, arg0)
}

// Clients is a free data retrieval call binding the contract method 0x20ba1e9f.
//
// Solidity: function clients(string ) view returns(address)
func (_Client *ClientCallerSession) Clients(arg0 string) (common.Address, error) {
	return _Client.Contract.Clients(&_Client.CallOpts, arg0)
}

// GetChainName is a free data retrieval call binding the contract method 0xd722b0bc.
//
// Solidity: function getChainName() view returns(string)
func (_Client *ClientCaller) GetChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "getChainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetChainName is a free data retrieval call binding the contract method 0xd722b0bc.
//
// Solidity: function getChainName() view returns(string)
func (_Client *ClientSession) GetChainName() (string, error) {
	return _Client.Contract.GetChainName(&_Client.CallOpts)
}

// GetChainName is a free data retrieval call binding the contract method 0xd722b0bc.
//
// Solidity: function getChainName() view returns(string)
func (_Client *ClientCallerSession) GetChainName() (string, error) {
	return _Client.Contract.GetChainName(&_Client.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string chainName) view returns((uint64,uint64))
func (_Client *ClientCaller) GetLatestHeight(opts *bind.CallOpts, chainName string) (HeightData, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "getLatestHeight", chainName)

	if err != nil {
		return *new(HeightData), err
	}

	out0 := *abi.ConvertType(out[0], new(HeightData)).(*HeightData)

	return out0, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string chainName) view returns((uint64,uint64))
func (_Client *ClientSession) GetLatestHeight(chainName string) (HeightData, error) {
	return _Client.Contract.GetLatestHeight(&_Client.CallOpts, chainName)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x329681d0.
//
// Solidity: function getLatestHeight(string chainName) view returns((uint64,uint64))
func (_Client *ClientCallerSession) GetLatestHeight(chainName string) (HeightData, error) {
	return _Client.Contract.GetLatestHeight(&_Client.CallOpts, chainName)
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

// Relayers is a free data retrieval call binding the contract method 0xee1ceb62.
//
// Solidity: function relayers(string , address ) view returns(bool)
func (_Client *ClientCaller) Relayers(opts *bind.CallOpts, arg0 string, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Client.contract.Call(opts, &out, "relayers", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Relayers is a free data retrieval call binding the contract method 0xee1ceb62.
//
// Solidity: function relayers(string , address ) view returns(bool)
func (_Client *ClientSession) Relayers(arg0 string, arg1 common.Address) (bool, error) {
	return _Client.Contract.Relayers(&_Client.CallOpts, arg0, arg1)
}

// Relayers is a free data retrieval call binding the contract method 0xee1ceb62.
//
// Solidity: function relayers(string , address ) view returns(bool)
func (_Client *ClientCallerSession) Relayers(arg0 string, arg1 common.Address) (bool, error) {
	return _Client.Contract.Relayers(&_Client.CallOpts, arg0, arg1)
}

// CreateClient is a paid mutator transaction binding the contract method 0x76262a47.
//
// Solidity: function createClient(string chainName, address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactor) CreateClient(opts *bind.TransactOpts, chainName string, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "createClient", chainName, clientAddress, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x76262a47.
//
// Solidity: function createClient(string chainName, address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientSession) CreateClient(chainName string, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.CreateClient(&_Client.TransactOpts, chainName, clientAddress, clientState, consensusState)
}

// CreateClient is a paid mutator transaction binding the contract method 0x76262a47.
//
// Solidity: function createClient(string chainName, address clientAddress, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactorSession) CreateClient(chainName string, clientAddress common.Address, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.CreateClient(&_Client.TransactOpts, chainName, clientAddress, clientState, consensusState)
}

// GetClient is a paid mutator transaction binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string chainName) returns(address)
func (_Client *ClientTransactor) GetClient(opts *bind.TransactOpts, chainName string) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "getClient", chainName)
}

// GetClient is a paid mutator transaction binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string chainName) returns(address)
func (_Client *ClientSession) GetClient(chainName string) (*types.Transaction, error) {
	return _Client.Contract.GetClient(&_Client.TransactOpts, chainName)
}

// GetClient is a paid mutator transaction binding the contract method 0x7eb78932.
//
// Solidity: function getClient(string chainName) returns(address)
func (_Client *ClientTransactorSession) GetClient(chainName string) (*types.Transaction, error) {
	return _Client.Contract.GetClient(&_Client.TransactOpts, chainName)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x5330a758.
//
// Solidity: function registerRelayer(string chainName, address relayer) returns()
func (_Client *ClientTransactor) RegisterRelayer(opts *bind.TransactOpts, chainName string, relayer common.Address) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "registerRelayer", chainName, relayer)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x5330a758.
//
// Solidity: function registerRelayer(string chainName, address relayer) returns()
func (_Client *ClientSession) RegisterRelayer(chainName string, relayer common.Address) (*types.Transaction, error) {
	return _Client.Contract.RegisterRelayer(&_Client.TransactOpts, chainName, relayer)
}

// RegisterRelayer is a paid mutator transaction binding the contract method 0x5330a758.
//
// Solidity: function registerRelayer(string chainName, address relayer) returns()
func (_Client *ClientTransactorSession) RegisterRelayer(chainName string, relayer common.Address) (*types.Transaction, error) {
	return _Client.Contract.RegisterRelayer(&_Client.TransactOpts, chainName, relayer)
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

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string chainName, bytes header) returns()
func (_Client *ClientTransactor) UpdateClient(opts *bind.TransactOpts, chainName string, header []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "updateClient", chainName, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string chainName, bytes header) returns()
func (_Client *ClientSession) UpdateClient(chainName string, header []byte) (*types.Transaction, error) {
	return _Client.Contract.UpdateClient(&_Client.TransactOpts, chainName, header)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x6fbf8079.
//
// Solidity: function updateClient(string chainName, bytes header) returns()
func (_Client *ClientTransactorSession) UpdateClient(chainName string, header []byte) (*types.Transaction, error) {
	return _Client.Contract.UpdateClient(&_Client.TransactOpts, chainName, header)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0x935aee64.
//
// Solidity: function upgradeClient(string chainName, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactor) UpgradeClient(opts *bind.TransactOpts, chainName string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.contract.Transact(opts, "upgradeClient", chainName, clientState, consensusState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0x935aee64.
//
// Solidity: function upgradeClient(string chainName, bytes clientState, bytes consensusState) returns()
func (_Client *ClientSession) UpgradeClient(chainName string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.UpgradeClient(&_Client.TransactOpts, chainName, clientState, consensusState)
}

// UpgradeClient is a paid mutator transaction binding the contract method 0x935aee64.
//
// Solidity: function upgradeClient(string chainName, bytes clientState, bytes consensusState) returns()
func (_Client *ClientTransactorSession) UpgradeClient(chainName string, clientState []byte, consensusState []byte) (*types.Transaction, error) {
	return _Client.Contract.UpgradeClient(&_Client.TransactOpts, chainName, clientState, consensusState)
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
