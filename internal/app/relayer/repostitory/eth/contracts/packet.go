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

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// PacketTypesCleanPacket is an auto generated low-level Go binding around an user-defined struct.
type PacketTypesCleanPacket struct {
	Sequence    uint64
	SourceChain string
	DestChain   string
	RelayChain  string
}

// PacketTypesPacket is an auto generated low-level Go binding around an user-defined struct.
type PacketTypesPacket struct {
	Sequence    uint64
	Port        string
	SourceChain string
	DestChain   string
	RelayChain  string
	Data        []byte
}

// PacketMetaData contains all meta data concerning the Packet contract.
var PacketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clientManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routing\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ack\",\"type\":\"bytes\"}],\"name\":\"AckWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structPacketTypes.CleanPacket\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"CleanPacketSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofAcked\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"acknowledgePacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"}],\"internalType\":\"structPacketTypes.CleanPacket\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"cleanPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clientManager\",\"outputs\":[{\"internalType\":\"contractIClientManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commitBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"}],\"name\":\"getNextSequenceSend\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"receipts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"}],\"internalType\":\"structPacketTypes.CleanPacket\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"recvCleanPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"recvPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routing\",\"outputs\":[{\"internalType\":\"contractIRouting\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sourceChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destChain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayChain\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structPacketTypes.Packet\",\"name\":\"packet\",\"type\":\"tuple\"}],\"name\":\"sendPacket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sequences\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clientManager\",\"type\":\"address\"}],\"name\":\"setClientManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_routing\",\"type\":\"address\"}],\"name\":\"setRouting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PacketABI is the input ABI used to generate the binding from.
// Deprecated: Use PacketMetaData.ABI instead.
var PacketABI = PacketMetaData.ABI

// Packet is an auto generated Go binding around an Ethereum contract.
type Packet struct {
	PacketCaller     // Read-only binding to the contract
	PacketTransactor // Write-only binding to the contract
	PacketFilterer   // Log filterer for contract events
}

// PacketCaller is an auto generated read-only Go binding around an Ethereum contract.
type PacketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PacketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PacketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PacketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PacketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PacketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PacketSession struct {
	Contract     *Packet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PacketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PacketCallerSession struct {
	Contract *PacketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PacketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PacketTransactorSession struct {
	Contract     *PacketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PacketRaw is an auto generated low-level Go binding around an Ethereum contract.
type PacketRaw struct {
	Contract *Packet // Generic contract binding to access the raw methods on
}

// PacketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PacketCallerRaw struct {
	Contract *PacketCaller // Generic read-only contract binding to access the raw methods on
}

// PacketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PacketTransactorRaw struct {
	Contract *PacketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPacket creates a new instance of Packet, bound to a specific deployed contract.
func NewPacket(address common.Address, backend bind.ContractBackend) (*Packet, error) {
	contract, err := bindPacket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Packet{PacketCaller: PacketCaller{contract: contract}, PacketTransactor: PacketTransactor{contract: contract}, PacketFilterer: PacketFilterer{contract: contract}}, nil
}

// NewPacketCaller creates a new read-only instance of Packet, bound to a specific deployed contract.
func NewPacketCaller(address common.Address, caller bind.ContractCaller) (*PacketCaller, error) {
	contract, err := bindPacket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PacketCaller{contract: contract}, nil
}

// NewPacketTransactor creates a new write-only instance of Packet, bound to a specific deployed contract.
func NewPacketTransactor(address common.Address, transactor bind.ContractTransactor) (*PacketTransactor, error) {
	contract, err := bindPacket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PacketTransactor{contract: contract}, nil
}

// NewPacketFilterer creates a new log filterer instance of Packet, bound to a specific deployed contract.
func NewPacketFilterer(address common.Address, filterer bind.ContractFilterer) (*PacketFilterer, error) {
	contract, err := bindPacket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PacketFilterer{contract: contract}, nil
}

// bindPacket binds a generic wrapper to an already deployed contract.
func bindPacket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PacketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Packet *PacketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Packet.Contract.PacketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Packet *PacketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packet.Contract.PacketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Packet *PacketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Packet.Contract.PacketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Packet *PacketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Packet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Packet *PacketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Packet *PacketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Packet.Contract.contract.Transact(opts, method, params...)
}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Packet *PacketCaller) ClientManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "clientManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Packet *PacketSession) ClientManager() (common.Address, error) {
	return _Packet.Contract.ClientManager(&_Packet.CallOpts)
}

// ClientManager is a free data retrieval call binding the contract method 0x79e8be1d.
//
// Solidity: function clientManager() view returns(address)
func (_Packet *PacketCallerSession) ClientManager() (common.Address, error) {
	return _Packet.Contract.ClientManager(&_Packet.CallOpts)
}

// CommitBytes is a free data retrieval call binding the contract method 0x29606763.
//
// Solidity: function commitBytes() view returns(bytes)
func (_Packet *PacketCaller) CommitBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "commitBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CommitBytes is a free data retrieval call binding the contract method 0x29606763.
//
// Solidity: function commitBytes() view returns(bytes)
func (_Packet *PacketSession) CommitBytes() ([]byte, error) {
	return _Packet.Contract.CommitBytes(&_Packet.CallOpts)
}

// CommitBytes is a free data retrieval call binding the contract method 0x29606763.
//
// Solidity: function commitBytes() view returns(bytes)
func (_Packet *PacketCallerSession) CommitBytes() ([]byte, error) {
	return _Packet.Contract.CommitBytes(&_Packet.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0x7912b8e6.
//
// Solidity: function commitments(bytes ) view returns(bytes32)
func (_Packet *PacketCaller) Commitments(opts *bind.CallOpts, arg0 []byte) ([32]byte, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x7912b8e6.
//
// Solidity: function commitments(bytes ) view returns(bytes32)
func (_Packet *PacketSession) Commitments(arg0 []byte) ([32]byte, error) {
	return _Packet.Contract.Commitments(&_Packet.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x7912b8e6.
//
// Solidity: function commitments(bytes ) view returns(bytes32)
func (_Packet *PacketCallerSession) Commitments(arg0 []byte) ([32]byte, error) {
	return _Packet.Contract.Commitments(&_Packet.CallOpts, arg0)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string sourceChain, string destChain) view returns(uint64)
func (_Packet *PacketCaller) GetNextSequenceSend(opts *bind.CallOpts, sourceChain string, destChain string) (uint64, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "getNextSequenceSend", sourceChain, destChain)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string sourceChain, string destChain) view returns(uint64)
func (_Packet *PacketSession) GetNextSequenceSend(sourceChain string, destChain string) (uint64, error) {
	return _Packet.Contract.GetNextSequenceSend(&_Packet.CallOpts, sourceChain, destChain)
}

// GetNextSequenceSend is a free data retrieval call binding the contract method 0x582418b6.
//
// Solidity: function getNextSequenceSend(string sourceChain, string destChain) view returns(uint64)
func (_Packet *PacketCallerSession) GetNextSequenceSend(sourceChain string, destChain string) (uint64, error) {
	return _Packet.Contract.GetNextSequenceSend(&_Packet.CallOpts, sourceChain, destChain)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Packet *PacketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Packet *PacketSession) Owner() (common.Address, error) {
	return _Packet.Contract.Owner(&_Packet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Packet *PacketCallerSession) Owner() (common.Address, error) {
	return _Packet.Contract.Owner(&_Packet.CallOpts)
}

// Receipts is a free data retrieval call binding the contract method 0xa6992b83.
//
// Solidity: function receipts(bytes ) view returns(bool)
func (_Packet *PacketCaller) Receipts(opts *bind.CallOpts, arg0 []byte) (bool, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "receipts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Receipts is a free data retrieval call binding the contract method 0xa6992b83.
//
// Solidity: function receipts(bytes ) view returns(bool)
func (_Packet *PacketSession) Receipts(arg0 []byte) (bool, error) {
	return _Packet.Contract.Receipts(&_Packet.CallOpts, arg0)
}

// Receipts is a free data retrieval call binding the contract method 0xa6992b83.
//
// Solidity: function receipts(bytes ) view returns(bool)
func (_Packet *PacketCallerSession) Receipts(arg0 []byte) (bool, error) {
	return _Packet.Contract.Receipts(&_Packet.CallOpts, arg0)
}

// Routing is a free data retrieval call binding the contract method 0x1b77f489.
//
// Solidity: function routing() view returns(address)
func (_Packet *PacketCaller) Routing(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "routing")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Routing is a free data retrieval call binding the contract method 0x1b77f489.
//
// Solidity: function routing() view returns(address)
func (_Packet *PacketSession) Routing() (common.Address, error) {
	return _Packet.Contract.Routing(&_Packet.CallOpts)
}

// Routing is a free data retrieval call binding the contract method 0x1b77f489.
//
// Solidity: function routing() view returns(address)
func (_Packet *PacketCallerSession) Routing() (common.Address, error) {
	return _Packet.Contract.Routing(&_Packet.CallOpts)
}

// Sequences is a free data retrieval call binding the contract method 0xeeebb020.
//
// Solidity: function sequences(bytes ) view returns(uint64)
func (_Packet *PacketCaller) Sequences(opts *bind.CallOpts, arg0 []byte) (uint64, error) {
	var out []interface{}
	err := _Packet.contract.Call(opts, &out, "sequences", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Sequences is a free data retrieval call binding the contract method 0xeeebb020.
//
// Solidity: function sequences(bytes ) view returns(uint64)
func (_Packet *PacketSession) Sequences(arg0 []byte) (uint64, error) {
	return _Packet.Contract.Sequences(&_Packet.CallOpts, arg0)
}

// Sequences is a free data retrieval call binding the contract method 0xeeebb020.
//
// Solidity: function sequences(bytes ) view returns(uint64)
func (_Packet *PacketCallerSession) Sequences(arg0 []byte) (uint64, error) {
	return _Packet.Contract.Sequences(&_Packet.CallOpts, arg0)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0x07f3f612.
//
// Solidity: function acknowledgePacket((uint64,string,string,string,string,bytes) packet, bytes acknowledgement, bytes proofAcked, (uint64,uint64) height) returns()
func (_Packet *PacketTransactor) AcknowledgePacket(opts *bind.TransactOpts, packet PacketTypesPacket, acknowledgement []byte, proofAcked []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "acknowledgePacket", packet, acknowledgement, proofAcked, height)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0x07f3f612.
//
// Solidity: function acknowledgePacket((uint64,string,string,string,string,bytes) packet, bytes acknowledgement, bytes proofAcked, (uint64,uint64) height) returns()
func (_Packet *PacketSession) AcknowledgePacket(packet PacketTypesPacket, acknowledgement []byte, proofAcked []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.Contract.AcknowledgePacket(&_Packet.TransactOpts, packet, acknowledgement, proofAcked, height)
}

// AcknowledgePacket is a paid mutator transaction binding the contract method 0x07f3f612.
//
// Solidity: function acknowledgePacket((uint64,string,string,string,string,bytes) packet, bytes acknowledgement, bytes proofAcked, (uint64,uint64) height) returns()
func (_Packet *PacketTransactorSession) AcknowledgePacket(packet PacketTypesPacket, acknowledgement []byte, proofAcked []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.Contract.AcknowledgePacket(&_Packet.TransactOpts, packet, acknowledgement, proofAcked, height)
}

// CleanPacket is a paid mutator transaction binding the contract method 0xb8fa9a9c.
//
// Solidity: function cleanPacket((uint64,string,string,string) packet) returns()
func (_Packet *PacketTransactor) CleanPacket(opts *bind.TransactOpts, packet PacketTypesCleanPacket) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "cleanPacket", packet)
}

// CleanPacket is a paid mutator transaction binding the contract method 0xb8fa9a9c.
//
// Solidity: function cleanPacket((uint64,string,string,string) packet) returns()
func (_Packet *PacketSession) CleanPacket(packet PacketTypesCleanPacket) (*types.Transaction, error) {
	return _Packet.Contract.CleanPacket(&_Packet.TransactOpts, packet)
}

// CleanPacket is a paid mutator transaction binding the contract method 0xb8fa9a9c.
//
// Solidity: function cleanPacket((uint64,string,string,string) packet) returns()
func (_Packet *PacketTransactorSession) CleanPacket(packet PacketTypesCleanPacket) (*types.Transaction, error) {
	return _Packet.Contract.CleanPacket(&_Packet.TransactOpts, packet)
}

// RecvCleanPacket is a paid mutator transaction binding the contract method 0x56d889ee.
//
// Solidity: function recvCleanPacket((uint64,string,string,string) packet, bytes proof, (uint64,uint64) height) returns()
func (_Packet *PacketTransactor) RecvCleanPacket(opts *bind.TransactOpts, packet PacketTypesCleanPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "recvCleanPacket", packet, proof, height)
}

// RecvCleanPacket is a paid mutator transaction binding the contract method 0x56d889ee.
//
// Solidity: function recvCleanPacket((uint64,string,string,string) packet, bytes proof, (uint64,uint64) height) returns()
func (_Packet *PacketSession) RecvCleanPacket(packet PacketTypesCleanPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.Contract.RecvCleanPacket(&_Packet.TransactOpts, packet, proof, height)
}

// RecvCleanPacket is a paid mutator transaction binding the contract method 0x56d889ee.
//
// Solidity: function recvCleanPacket((uint64,string,string,string) packet, bytes proof, (uint64,uint64) height) returns()
func (_Packet *PacketTransactorSession) RecvCleanPacket(packet PacketTypesCleanPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.Contract.RecvCleanPacket(&_Packet.TransactOpts, packet, proof, height)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x1c7249de.
//
// Solidity: function recvPacket((uint64,string,string,string,string,bytes) packet, bytes proof, (uint64,uint64) height) returns()
func (_Packet *PacketTransactor) RecvPacket(opts *bind.TransactOpts, packet PacketTypesPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "recvPacket", packet, proof, height)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x1c7249de.
//
// Solidity: function recvPacket((uint64,string,string,string,string,bytes) packet, bytes proof, (uint64,uint64) height) returns()
func (_Packet *PacketSession) RecvPacket(packet PacketTypesPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.Contract.RecvPacket(&_Packet.TransactOpts, packet, proof, height)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x1c7249de.
//
// Solidity: function recvPacket((uint64,string,string,string,string,bytes) packet, bytes proof, (uint64,uint64) height) returns()
func (_Packet *PacketTransactorSession) RecvPacket(packet PacketTypesPacket, proof []byte, height HeightData) (*types.Transaction, error) {
	return _Packet.Contract.RecvPacket(&_Packet.TransactOpts, packet, proof, height)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Packet *PacketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Packet *PacketSession) RenounceOwnership() (*types.Transaction, error) {
	return _Packet.Contract.RenounceOwnership(&_Packet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Packet *PacketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Packet.Contract.RenounceOwnership(&_Packet.TransactOpts)
}

// SendPacket is a paid mutator transaction binding the contract method 0x7d086e38.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes) packet) returns()
func (_Packet *PacketTransactor) SendPacket(opts *bind.TransactOpts, packet PacketTypesPacket) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "sendPacket", packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x7d086e38.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes) packet) returns()
func (_Packet *PacketSession) SendPacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Packet.Contract.SendPacket(&_Packet.TransactOpts, packet)
}

// SendPacket is a paid mutator transaction binding the contract method 0x7d086e38.
//
// Solidity: function sendPacket((uint64,string,string,string,string,bytes) packet) returns()
func (_Packet *PacketTransactorSession) SendPacket(packet PacketTypesPacket) (*types.Transaction, error) {
	return _Packet.Contract.SendPacket(&_Packet.TransactOpts, packet)
}

// SetClientManager is a paid mutator transaction binding the contract method 0x8f3de42e.
//
// Solidity: function setClientManager(address _clientManager) returns()
func (_Packet *PacketTransactor) SetClientManager(opts *bind.TransactOpts, _clientManager common.Address) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "setClientManager", _clientManager)
}

// SetClientManager is a paid mutator transaction binding the contract method 0x8f3de42e.
//
// Solidity: function setClientManager(address _clientManager) returns()
func (_Packet *PacketSession) SetClientManager(_clientManager common.Address) (*types.Transaction, error) {
	return _Packet.Contract.SetClientManager(&_Packet.TransactOpts, _clientManager)
}

// SetClientManager is a paid mutator transaction binding the contract method 0x8f3de42e.
//
// Solidity: function setClientManager(address _clientManager) returns()
func (_Packet *PacketTransactorSession) SetClientManager(_clientManager common.Address) (*types.Transaction, error) {
	return _Packet.Contract.SetClientManager(&_Packet.TransactOpts, _clientManager)
}

// SetRouting is a paid mutator transaction binding the contract method 0x109f543f.
//
// Solidity: function setRouting(address _routing) returns()
func (_Packet *PacketTransactor) SetRouting(opts *bind.TransactOpts, _routing common.Address) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "setRouting", _routing)
}

// SetRouting is a paid mutator transaction binding the contract method 0x109f543f.
//
// Solidity: function setRouting(address _routing) returns()
func (_Packet *PacketSession) SetRouting(_routing common.Address) (*types.Transaction, error) {
	return _Packet.Contract.SetRouting(&_Packet.TransactOpts, _routing)
}

// SetRouting is a paid mutator transaction binding the contract method 0x109f543f.
//
// Solidity: function setRouting(address _routing) returns()
func (_Packet *PacketTransactorSession) SetRouting(_routing common.Address) (*types.Transaction, error) {
	return _Packet.Contract.SetRouting(&_Packet.TransactOpts, _routing)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Packet *PacketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Packet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Packet *PacketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Packet.Contract.TransferOwnership(&_Packet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Packet *PacketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Packet.Contract.TransferOwnership(&_Packet.TransactOpts, newOwner)
}

// PacketAckWrittenIterator is returned from FilterAckWritten and is used to iterate over the raw logs and unpacked data for AckWritten events raised by the Packet contract.
type PacketAckWrittenIterator struct {
	Event *PacketAckWritten // Event containing the contract specifics and raw log

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
func (it *PacketAckWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PacketAckWritten)
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
		it.Event = new(PacketAckWritten)
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
func (it *PacketAckWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PacketAckWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PacketAckWritten represents a AckWritten event raised by the Packet contract.
type PacketAckWritten struct {
	Packet PacketTypesPacket
	Ack    []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAckWritten is a free log retrieval operation binding the contract event 0x773ac32e677533310d84b93686be4f715d2bd399c322a619529ee6e3ce3f25e7.
//
// Solidity: event AckWritten((uint64,string,string,string,string,bytes) packet, bytes ack)
func (_Packet *PacketFilterer) FilterAckWritten(opts *bind.FilterOpts) (*PacketAckWrittenIterator, error) {

	logs, sub, err := _Packet.contract.FilterLogs(opts, "AckWritten")
	if err != nil {
		return nil, err
	}
	return &PacketAckWrittenIterator{contract: _Packet.contract, event: "AckWritten", logs: logs, sub: sub}, nil
}

// WatchAckWritten is a free log subscription operation binding the contract event 0x773ac32e677533310d84b93686be4f715d2bd399c322a619529ee6e3ce3f25e7.
//
// Solidity: event AckWritten((uint64,string,string,string,string,bytes) packet, bytes ack)
func (_Packet *PacketFilterer) WatchAckWritten(opts *bind.WatchOpts, sink chan<- *PacketAckWritten) (event.Subscription, error) {

	logs, sub, err := _Packet.contract.WatchLogs(opts, "AckWritten")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PacketAckWritten)
				if err := _Packet.contract.UnpackLog(event, "AckWritten", log); err != nil {
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

// ParseAckWritten is a log parse operation binding the contract event 0x773ac32e677533310d84b93686be4f715d2bd399c322a619529ee6e3ce3f25e7.
//
// Solidity: event AckWritten((uint64,string,string,string,string,bytes) packet, bytes ack)
func (_Packet *PacketFilterer) ParseAckWritten(log types.Log) (*PacketAckWritten, error) {
	event := new(PacketAckWritten)
	if err := _Packet.contract.UnpackLog(event, "AckWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PacketCleanPacketSentIterator is returned from FilterCleanPacketSent and is used to iterate over the raw logs and unpacked data for CleanPacketSent events raised by the Packet contract.
type PacketCleanPacketSentIterator struct {
	Event *PacketCleanPacketSent // Event containing the contract specifics and raw log

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
func (it *PacketCleanPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PacketCleanPacketSent)
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
		it.Event = new(PacketCleanPacketSent)
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
func (it *PacketCleanPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PacketCleanPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PacketCleanPacketSent represents a CleanPacketSent event raised by the Packet contract.
type PacketCleanPacketSent struct {
	Packet PacketTypesCleanPacket
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCleanPacketSent is a free log retrieval operation binding the contract event 0x755e688e418e0c7eecdc9584c4c0298bb4397d6dcc817d25aa46bb9dd5b9cafc.
//
// Solidity: event CleanPacketSent((uint64,string,string,string) packet)
func (_Packet *PacketFilterer) FilterCleanPacketSent(opts *bind.FilterOpts) (*PacketCleanPacketSentIterator, error) {

	logs, sub, err := _Packet.contract.FilterLogs(opts, "CleanPacketSent")
	if err != nil {
		return nil, err
	}
	return &PacketCleanPacketSentIterator{contract: _Packet.contract, event: "CleanPacketSent", logs: logs, sub: sub}, nil
}

// WatchCleanPacketSent is a free log subscription operation binding the contract event 0x755e688e418e0c7eecdc9584c4c0298bb4397d6dcc817d25aa46bb9dd5b9cafc.
//
// Solidity: event CleanPacketSent((uint64,string,string,string) packet)
func (_Packet *PacketFilterer) WatchCleanPacketSent(opts *bind.WatchOpts, sink chan<- *PacketCleanPacketSent) (event.Subscription, error) {

	logs, sub, err := _Packet.contract.WatchLogs(opts, "CleanPacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PacketCleanPacketSent)
				if err := _Packet.contract.UnpackLog(event, "CleanPacketSent", log); err != nil {
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

// ParseCleanPacketSent is a log parse operation binding the contract event 0x755e688e418e0c7eecdc9584c4c0298bb4397d6dcc817d25aa46bb9dd5b9cafc.
//
// Solidity: event CleanPacketSent((uint64,string,string,string) packet)
func (_Packet *PacketFilterer) ParseCleanPacketSent(log types.Log) (*PacketCleanPacketSent, error) {
	event := new(PacketCleanPacketSent)
	if err := _Packet.contract.UnpackLog(event, "CleanPacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PacketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Packet contract.
type PacketOwnershipTransferredIterator struct {
	Event *PacketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PacketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PacketOwnershipTransferred)
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
		it.Event = new(PacketOwnershipTransferred)
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
func (it *PacketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PacketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PacketOwnershipTransferred represents a OwnershipTransferred event raised by the Packet contract.
type PacketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Packet *PacketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PacketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Packet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PacketOwnershipTransferredIterator{contract: _Packet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Packet *PacketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PacketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Packet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PacketOwnershipTransferred)
				if err := _Packet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Packet *PacketFilterer) ParseOwnershipTransferred(log types.Log) (*PacketOwnershipTransferred, error) {
	event := new(PacketOwnershipTransferred)
	if err := _Packet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PacketPacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the Packet contract.
type PacketPacketSentIterator struct {
	Event *PacketPacketSent // Event containing the contract specifics and raw log

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
func (it *PacketPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PacketPacketSent)
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
		it.Event = new(PacketPacketSent)
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
func (it *PacketPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PacketPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PacketPacketSent represents a PacketSent event raised by the Packet contract.
type PacketPacketSent struct {
	Packet PacketTypesPacket
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x9f4eeeeeb04711557a8bca7f4ac577fc7db19653f604920f381f7358e8a686a0.
//
// Solidity: event PacketSent((uint64,string,string,string,string,bytes) packet)
func (_Packet *PacketFilterer) FilterPacketSent(opts *bind.FilterOpts) (*PacketPacketSentIterator, error) {

	logs, sub, err := _Packet.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &PacketPacketSentIterator{contract: _Packet.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x9f4eeeeeb04711557a8bca7f4ac577fc7db19653f604920f381f7358e8a686a0.
//
// Solidity: event PacketSent((uint64,string,string,string,string,bytes) packet)
func (_Packet *PacketFilterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *PacketPacketSent) (event.Subscription, error) {

	logs, sub, err := _Packet.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PacketPacketSent)
				if err := _Packet.contract.UnpackLog(event, "PacketSent", log); err != nil {
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

// ParsePacketSent is a log parse operation binding the contract event 0x9f4eeeeeb04711557a8bca7f4ac577fc7db19653f604920f381f7358e8a686a0.
//
// Solidity: event PacketSent((uint64,string,string,string,string,bytes) packet)
func (_Packet *PacketFilterer) ParsePacketSent(log types.Log) (*PacketPacketSent, error) {
	event := new(PacketPacketSent)
	if err := _Packet.contract.UnpackLog(event, "PacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
