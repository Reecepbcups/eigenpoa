// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake_registry

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
	_ = abi.ConvertType
)

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// Quorum is an auto generated low-level Go binding around an user-defined struct.
type Quorum struct {
	Strategies []StrategyParams
}

// StrategyParams is an auto generated low-level Go binding around an user-defined struct.
type StrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// StakeRegistryMetaData contains all meta data concerning the StakeRegistry contract.
var StakeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLastCheckpointOperatorWeight\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointThresholdWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointThresholdWeightAtBlock\",\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointTotalWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastCheckpointTotalWeightAtBlock\",\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastestOperatorSigningKey\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSigningKeyAtBlock\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeight\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeightAtBlock\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_serviceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_thresholdWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_quorum\",\"type\":\"tuple\",\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_dataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signatureData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorRegistered\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorWithSignature\",\"inputs\":[{\"name\":\"_operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_signingKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinimumWeight\",\"inputs\":[{\"name\":\"_newMinimumWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorSigningKey\",\"inputs\":[{\"name\":\"_newSigningKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperators\",\"inputs\":[{\"name\":\"_operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorsForQuorum\",\"inputs\":[{\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\",\"internalType\":\"address[][]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQuorumConfig\",\"inputs\":[{\"name\":\"_quorum\",\"type\":\"tuple\",\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"_operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateStakeThreshold\",\"inputs\":[{\"name\":\"_thresholdWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumWeightUpdated\",\"inputs\":[{\"name\":\"_old\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_new\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorWeightUpdated\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumUpdated\",\"inputs\":[{\"name\":\"_old\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"_new\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structQuorum\",\"components\":[{\"name\":\"strategies\",\"type\":\"tuple[]\",\"internalType\":\"structStrategyParams[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SigningKeyUpdate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"updateBlock\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newSigningKey\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldSigningKey\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdWeightUpdated\",\"inputs\":[{\"name\":\"_thresholdWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalWeightUpdated\",\"inputs\":[{\"name\":\"oldTotalWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newTotalWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateMinimumWeight\",\"inputs\":[{\"name\":\"oldMinimumWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinimumWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientSignedStake\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuorum\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidReferenceBlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignedWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustUpdateAllOperators\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]}]",
}

// StakeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeRegistryMetaData.ABI instead.
var StakeRegistryABI = StakeRegistryMetaData.ABI

// StakeRegistry is an auto generated Go binding around an Ethereum contract.
type StakeRegistry struct {
	StakeRegistryCaller     // Read-only binding to the contract
	StakeRegistryTransactor // Write-only binding to the contract
	StakeRegistryFilterer   // Log filterer for contract events
}

// StakeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeRegistrySession struct {
	Contract     *StakeRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeRegistryCallerSession struct {
	Contract *StakeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeRegistryTransactorSession struct {
	Contract     *StakeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRegistryRaw struct {
	Contract *StakeRegistry // Generic contract binding to access the raw methods on
}

// StakeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeRegistryCallerRaw struct {
	Contract *StakeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// StakeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeRegistryTransactorRaw struct {
	Contract *StakeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeRegistry creates a new instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistry(address common.Address, backend bind.ContractBackend) (*StakeRegistry, error) {
	contract, err := bindStakeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeRegistry{StakeRegistryCaller: StakeRegistryCaller{contract: contract}, StakeRegistryTransactor: StakeRegistryTransactor{contract: contract}, StakeRegistryFilterer: StakeRegistryFilterer{contract: contract}}, nil
}

// NewStakeRegistryCaller creates a new read-only instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistryCaller(address common.Address, caller bind.ContractCaller) (*StakeRegistryCaller, error) {
	contract, err := bindStakeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryCaller{contract: contract}, nil
}

// NewStakeRegistryTransactor creates a new write-only instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeRegistryTransactor, error) {
	contract, err := bindStakeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryTransactor{contract: contract}, nil
}

// NewStakeRegistryFilterer creates a new log filterer instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeRegistryFilterer, error) {
	contract, err := bindStakeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryFilterer{contract: contract}, nil
}

// bindStakeRegistry binds a generic wrapper to an already deployed contract.
func bindStakeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeRegistry *StakeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeRegistry.Contract.StakeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeRegistry *StakeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRegistry.Contract.StakeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeRegistry *StakeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeRegistry.Contract.StakeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeRegistry *StakeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeRegistry *StakeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeRegistry *StakeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetLastCheckpointOperatorWeight(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getLastCheckpointOperatorWeight", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetLastCheckpointOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointOperatorWeight(&_StakeRegistry.CallOpts, _operator)
}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetLastCheckpointOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointOperatorWeight(&_StakeRegistry.CallOpts, _operator)
}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetLastCheckpointThresholdWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getLastCheckpointThresholdWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetLastCheckpointThresholdWeight() (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointThresholdWeight(&_StakeRegistry.CallOpts)
}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetLastCheckpointThresholdWeight() (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointThresholdWeight(&_StakeRegistry.CallOpts)
}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetLastCheckpointThresholdWeightAtBlock(opts *bind.CallOpts, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getLastCheckpointThresholdWeightAtBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetLastCheckpointThresholdWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointThresholdWeightAtBlock(&_StakeRegistry.CallOpts, _blockNumber)
}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetLastCheckpointThresholdWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointThresholdWeightAtBlock(&_StakeRegistry.CallOpts, _blockNumber)
}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetLastCheckpointTotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getLastCheckpointTotalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetLastCheckpointTotalWeight() (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointTotalWeight(&_StakeRegistry.CallOpts)
}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetLastCheckpointTotalWeight() (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointTotalWeight(&_StakeRegistry.CallOpts)
}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetLastCheckpointTotalWeightAtBlock(opts *bind.CallOpts, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getLastCheckpointTotalWeightAtBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetLastCheckpointTotalWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointTotalWeightAtBlock(&_StakeRegistry.CallOpts, _blockNumber)
}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetLastCheckpointTotalWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetLastCheckpointTotalWeightAtBlock(&_StakeRegistry.CallOpts, _blockNumber)
}

// GetLastestOperatorSigningKey is a free data retrieval call binding the contract method 0xcdcd3581.
//
// Solidity: function getLastestOperatorSigningKey(address _operator) view returns(address)
func (_StakeRegistry *StakeRegistryCaller) GetLastestOperatorSigningKey(opts *bind.CallOpts, _operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getLastestOperatorSigningKey", _operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLastestOperatorSigningKey is a free data retrieval call binding the contract method 0xcdcd3581.
//
// Solidity: function getLastestOperatorSigningKey(address _operator) view returns(address)
func (_StakeRegistry *StakeRegistrySession) GetLastestOperatorSigningKey(_operator common.Address) (common.Address, error) {
	return _StakeRegistry.Contract.GetLastestOperatorSigningKey(&_StakeRegistry.CallOpts, _operator)
}

// GetLastestOperatorSigningKey is a free data retrieval call binding the contract method 0xcdcd3581.
//
// Solidity: function getLastestOperatorSigningKey(address _operator) view returns(address)
func (_StakeRegistry *StakeRegistryCallerSession) GetLastestOperatorSigningKey(_operator common.Address) (common.Address, error) {
	return _StakeRegistry.Contract.GetLastestOperatorSigningKey(&_StakeRegistry.CallOpts, _operator)
}

// GetOperatorSigningKeyAtBlock is a free data retrieval call binding the contract method 0x5e1042e8.
//
// Solidity: function getOperatorSigningKeyAtBlock(address _operator, uint256 _blockNumber) view returns(address)
func (_StakeRegistry *StakeRegistryCaller) GetOperatorSigningKeyAtBlock(opts *bind.CallOpts, _operator common.Address, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getOperatorSigningKeyAtBlock", _operator, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSigningKeyAtBlock is a free data retrieval call binding the contract method 0x5e1042e8.
//
// Solidity: function getOperatorSigningKeyAtBlock(address _operator, uint256 _blockNumber) view returns(address)
func (_StakeRegistry *StakeRegistrySession) GetOperatorSigningKeyAtBlock(_operator common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _StakeRegistry.Contract.GetOperatorSigningKeyAtBlock(&_StakeRegistry.CallOpts, _operator, _blockNumber)
}

// GetOperatorSigningKeyAtBlock is a free data retrieval call binding the contract method 0x5e1042e8.
//
// Solidity: function getOperatorSigningKeyAtBlock(address _operator, uint256 _blockNumber) view returns(address)
func (_StakeRegistry *StakeRegistryCallerSession) GetOperatorSigningKeyAtBlock(_operator common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _StakeRegistry.Contract.GetOperatorSigningKeyAtBlock(&_StakeRegistry.CallOpts, _operator, _blockNumber)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetOperatorWeight(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getOperatorWeight", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _StakeRegistry.Contract.GetOperatorWeight(&_StakeRegistry.CallOpts, _operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _StakeRegistry.Contract.GetOperatorWeight(&_StakeRegistry.CallOpts, _operator)
}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetOperatorWeightAtBlock(opts *bind.CallOpts, _operator common.Address, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getOperatorWeightAtBlock", _operator, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetOperatorWeightAtBlock(_operator common.Address, _blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetOperatorWeightAtBlock(&_StakeRegistry.CallOpts, _operator, _blockNumber)
}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetOperatorWeightAtBlock(_operator common.Address, _blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetOperatorWeightAtBlock(&_StakeRegistry.CallOpts, _operator, _blockNumber)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_StakeRegistry *StakeRegistryCaller) IsValidSignature(opts *bind.CallOpts, _dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "isValidSignature", _dataHash, _signatureData)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_StakeRegistry *StakeRegistrySession) IsValidSignature(_dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	return _StakeRegistry.Contract.IsValidSignature(&_StakeRegistry.CallOpts, _dataHash, _signatureData)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_StakeRegistry *StakeRegistryCallerSession) IsValidSignature(_dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	return _StakeRegistry.Contract.IsValidSignature(&_StakeRegistry.CallOpts, _dataHash, _signatureData)
}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) MinimumWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "minimumWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) MinimumWeight() (*big.Int, error) {
	return _StakeRegistry.Contract.MinimumWeight(&_StakeRegistry.CallOpts)
}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) MinimumWeight() (*big.Int, error) {
	return _StakeRegistry.Contract.MinimumWeight(&_StakeRegistry.CallOpts)
}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_StakeRegistry *StakeRegistryCaller) OperatorRegistered(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "operatorRegistered", _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_StakeRegistry *StakeRegistrySession) OperatorRegistered(_operator common.Address) (bool, error) {
	return _StakeRegistry.Contract.OperatorRegistered(&_StakeRegistry.CallOpts, _operator)
}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_StakeRegistry *StakeRegistryCallerSession) OperatorRegistered(_operator common.Address) (bool, error) {
	return _StakeRegistry.Contract.OperatorRegistered(&_StakeRegistry.CallOpts, _operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeRegistry *StakeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeRegistry *StakeRegistrySession) Owner() (common.Address, error) {
	return _StakeRegistry.Contract.Owner(&_StakeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeRegistry *StakeRegistryCallerSession) Owner() (common.Address, error) {
	return _StakeRegistry.Contract.Owner(&_StakeRegistry.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_StakeRegistry *StakeRegistryCaller) Quorum(opts *bind.CallOpts) (Quorum, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(Quorum), err
	}

	out0 := *abi.ConvertType(out[0], new(Quorum)).(*Quorum)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_StakeRegistry *StakeRegistrySession) Quorum() (Quorum, error) {
	return _StakeRegistry.Contract.Quorum(&_StakeRegistry.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_StakeRegistry *StakeRegistryCallerSession) Quorum() (Quorum, error) {
	return _StakeRegistry.Contract.Quorum(&_StakeRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_StakeRegistry *StakeRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "deregisterOperator")
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_StakeRegistry *StakeRegistrySession) DeregisterOperator() (*types.Transaction, error) {
	return _StakeRegistry.Contract.DeregisterOperator(&_StakeRegistry.TransactOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_StakeRegistry *StakeRegistryTransactorSession) DeregisterOperator() (*types.Transaction, error) {
	return _StakeRegistry.Contract.DeregisterOperator(&_StakeRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_StakeRegistry *StakeRegistryTransactor) Initialize(opts *bind.TransactOpts, _serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "initialize", _serviceManager, _thresholdWeight, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_StakeRegistry *StakeRegistrySession) Initialize(_serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _StakeRegistry.Contract.Initialize(&_StakeRegistry.TransactOpts, _serviceManager, _thresholdWeight, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) Initialize(_serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _StakeRegistry.Contract.Initialize(&_StakeRegistry.TransactOpts, _serviceManager, _thresholdWeight, _quorum)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x3d5611f6.
//
// Solidity: function registerOperatorWithSignature((bytes,bytes32,uint256) _operatorSignature, address _signingKey) returns()
func (_StakeRegistry *StakeRegistryTransactor) RegisterOperatorWithSignature(opts *bind.TransactOpts, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _signingKey common.Address) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "registerOperatorWithSignature", _operatorSignature, _signingKey)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x3d5611f6.
//
// Solidity: function registerOperatorWithSignature((bytes,bytes32,uint256) _operatorSignature, address _signingKey) returns()
func (_StakeRegistry *StakeRegistrySession) RegisterOperatorWithSignature(_operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _signingKey common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.RegisterOperatorWithSignature(&_StakeRegistry.TransactOpts, _operatorSignature, _signingKey)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x3d5611f6.
//
// Solidity: function registerOperatorWithSignature((bytes,bytes32,uint256) _operatorSignature, address _signingKey) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) RegisterOperatorWithSignature(_operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _signingKey common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.RegisterOperatorWithSignature(&_StakeRegistry.TransactOpts, _operatorSignature, _signingKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeRegistry *StakeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeRegistry *StakeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _StakeRegistry.Contract.RenounceOwnership(&_StakeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeRegistry *StakeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakeRegistry.Contract.RenounceOwnership(&_StakeRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeRegistry *StakeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeRegistry *StakeRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.TransferOwnership(&_StakeRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.TransferOwnership(&_StakeRegistry.TransactOpts, newOwner)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_StakeRegistry *StakeRegistryTransactor) UpdateMinimumWeight(opts *bind.TransactOpts, _newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateMinimumWeight", _newMinimumWeight, _operators)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_StakeRegistry *StakeRegistrySession) UpdateMinimumWeight(_newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateMinimumWeight(&_StakeRegistry.TransactOpts, _newMinimumWeight, _operators)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateMinimumWeight(_newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateMinimumWeight(&_StakeRegistry.TransactOpts, _newMinimumWeight, _operators)
}

// UpdateOperatorSigningKey is a paid mutator transaction binding the contract method 0x743c31f4.
//
// Solidity: function updateOperatorSigningKey(address _newSigningKey) returns()
func (_StakeRegistry *StakeRegistryTransactor) UpdateOperatorSigningKey(opts *bind.TransactOpts, _newSigningKey common.Address) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateOperatorSigningKey", _newSigningKey)
}

// UpdateOperatorSigningKey is a paid mutator transaction binding the contract method 0x743c31f4.
//
// Solidity: function updateOperatorSigningKey(address _newSigningKey) returns()
func (_StakeRegistry *StakeRegistrySession) UpdateOperatorSigningKey(_newSigningKey common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorSigningKey(&_StakeRegistry.TransactOpts, _newSigningKey)
}

// UpdateOperatorSigningKey is a paid mutator transaction binding the contract method 0x743c31f4.
//
// Solidity: function updateOperatorSigningKey(address _newSigningKey) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateOperatorSigningKey(_newSigningKey common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorSigningKey(&_StakeRegistry.TransactOpts, _newSigningKey)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_StakeRegistry *StakeRegistryTransactor) UpdateOperators(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateOperators", _operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_StakeRegistry *StakeRegistrySession) UpdateOperators(_operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperators(&_StakeRegistry.TransactOpts, _operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateOperators(_operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperators(&_StakeRegistry.TransactOpts, _operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_StakeRegistry *StakeRegistryTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, arg1)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_StakeRegistry *StakeRegistrySession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorsForQuorum(&_StakeRegistry.TransactOpts, operatorsPerQuorum, arg1)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorsForQuorum(&_StakeRegistry.TransactOpts, operatorsPerQuorum, arg1)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_StakeRegistry *StakeRegistryTransactor) UpdateQuorumConfig(opts *bind.TransactOpts, _quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateQuorumConfig", _quorum, _operators)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_StakeRegistry *StakeRegistrySession) UpdateQuorumConfig(_quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateQuorumConfig(&_StakeRegistry.TransactOpts, _quorum, _operators)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateQuorumConfig(_quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateQuorumConfig(&_StakeRegistry.TransactOpts, _quorum, _operators)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_StakeRegistry *StakeRegistryTransactor) UpdateStakeThreshold(opts *bind.TransactOpts, _thresholdWeight *big.Int) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateStakeThreshold", _thresholdWeight)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_StakeRegistry *StakeRegistrySession) UpdateStakeThreshold(_thresholdWeight *big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateStakeThreshold(&_StakeRegistry.TransactOpts, _thresholdWeight)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateStakeThreshold(_thresholdWeight *big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateStakeThreshold(&_StakeRegistry.TransactOpts, _thresholdWeight)
}

// StakeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakeRegistry contract.
type StakeRegistryInitializedIterator struct {
	Event *StakeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *StakeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryInitialized)
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
		it.Event = new(StakeRegistryInitialized)
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
func (it *StakeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryInitialized represents a Initialized event raised by the StakeRegistry contract.
type StakeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeRegistry *StakeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakeRegistryInitializedIterator, error) {

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakeRegistryInitializedIterator{contract: _StakeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeRegistry *StakeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryInitialized)
				if err := _StakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeRegistry *StakeRegistryFilterer) ParseInitialized(log types.Log) (*StakeRegistryInitialized, error) {
	event := new(StakeRegistryInitialized)
	if err := _StakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryMinimumWeightUpdatedIterator is returned from FilterMinimumWeightUpdated and is used to iterate over the raw logs and unpacked data for MinimumWeightUpdated events raised by the StakeRegistry contract.
type StakeRegistryMinimumWeightUpdatedIterator struct {
	Event *StakeRegistryMinimumWeightUpdated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryMinimumWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryMinimumWeightUpdated)
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
		it.Event = new(StakeRegistryMinimumWeightUpdated)
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
func (it *StakeRegistryMinimumWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryMinimumWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryMinimumWeightUpdated represents a MinimumWeightUpdated event raised by the StakeRegistry contract.
type StakeRegistryMinimumWeightUpdated struct {
	Old *big.Int
	New *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinimumWeightUpdated is a free log retrieval operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_StakeRegistry *StakeRegistryFilterer) FilterMinimumWeightUpdated(opts *bind.FilterOpts) (*StakeRegistryMinimumWeightUpdatedIterator, error) {

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "MinimumWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeRegistryMinimumWeightUpdatedIterator{contract: _StakeRegistry.contract, event: "MinimumWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumWeightUpdated is a free log subscription operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_StakeRegistry *StakeRegistryFilterer) WatchMinimumWeightUpdated(opts *bind.WatchOpts, sink chan<- *StakeRegistryMinimumWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "MinimumWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryMinimumWeightUpdated)
				if err := _StakeRegistry.contract.UnpackLog(event, "MinimumWeightUpdated", log); err != nil {
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

// ParseMinimumWeightUpdated is a log parse operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_StakeRegistry *StakeRegistryFilterer) ParseMinimumWeightUpdated(log types.Log) (*StakeRegistryMinimumWeightUpdated, error) {
	event := new(StakeRegistryMinimumWeightUpdated)
	if err := _StakeRegistry.contract.UnpackLog(event, "MinimumWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the StakeRegistry contract.
type StakeRegistryOperatorDeregisteredIterator struct {
	Event *StakeRegistryOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *StakeRegistryOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryOperatorDeregistered)
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
		it.Event = new(StakeRegistryOperatorDeregistered)
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
func (it *StakeRegistryOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryOperatorDeregistered represents a OperatorDeregistered event raised by the StakeRegistry contract.
type StakeRegistryOperatorDeregistered struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_StakeRegistry *StakeRegistryFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, _operator []common.Address, _avs []common.Address) (*StakeRegistryOperatorDeregisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "OperatorDeregistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryOperatorDeregisteredIterator{contract: _StakeRegistry.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_StakeRegistry *StakeRegistryFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *StakeRegistryOperatorDeregistered, _operator []common.Address, _avs []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "OperatorDeregistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryOperatorDeregistered)
				if err := _StakeRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_StakeRegistry *StakeRegistryFilterer) ParseOperatorDeregistered(log types.Log) (*StakeRegistryOperatorDeregistered, error) {
	event := new(StakeRegistryOperatorDeregistered)
	if err := _StakeRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the StakeRegistry contract.
type StakeRegistryOperatorRegisteredIterator struct {
	Event *StakeRegistryOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *StakeRegistryOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryOperatorRegistered)
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
		it.Event = new(StakeRegistryOperatorRegistered)
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
func (it *StakeRegistryOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryOperatorRegistered represents a OperatorRegistered event raised by the StakeRegistry contract.
type StakeRegistryOperatorRegistered struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_StakeRegistry *StakeRegistryFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, _operator []common.Address, _avs []common.Address) (*StakeRegistryOperatorRegisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "OperatorRegistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryOperatorRegisteredIterator{contract: _StakeRegistry.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_StakeRegistry *StakeRegistryFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *StakeRegistryOperatorRegistered, _operator []common.Address, _avs []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "OperatorRegistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryOperatorRegistered)
				if err := _StakeRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_StakeRegistry *StakeRegistryFilterer) ParseOperatorRegistered(log types.Log) (*StakeRegistryOperatorRegistered, error) {
	event := new(StakeRegistryOperatorRegistered)
	if err := _StakeRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryOperatorWeightUpdatedIterator is returned from FilterOperatorWeightUpdated and is used to iterate over the raw logs and unpacked data for OperatorWeightUpdated events raised by the StakeRegistry contract.
type StakeRegistryOperatorWeightUpdatedIterator struct {
	Event *StakeRegistryOperatorWeightUpdated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryOperatorWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryOperatorWeightUpdated)
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
		it.Event = new(StakeRegistryOperatorWeightUpdated)
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
func (it *StakeRegistryOperatorWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryOperatorWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryOperatorWeightUpdated represents a OperatorWeightUpdated event raised by the StakeRegistry contract.
type StakeRegistryOperatorWeightUpdated struct {
	Operator  common.Address
	OldWeight *big.Int
	NewWeight *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorWeightUpdated is a free log retrieval operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_StakeRegistry *StakeRegistryFilterer) FilterOperatorWeightUpdated(opts *bind.FilterOpts, _operator []common.Address) (*StakeRegistryOperatorWeightUpdatedIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "OperatorWeightUpdated", _operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryOperatorWeightUpdatedIterator{contract: _StakeRegistry.contract, event: "OperatorWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWeightUpdated is a free log subscription operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_StakeRegistry *StakeRegistryFilterer) WatchOperatorWeightUpdated(opts *bind.WatchOpts, sink chan<- *StakeRegistryOperatorWeightUpdated, _operator []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "OperatorWeightUpdated", _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryOperatorWeightUpdated)
				if err := _StakeRegistry.contract.UnpackLog(event, "OperatorWeightUpdated", log); err != nil {
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

// ParseOperatorWeightUpdated is a log parse operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_StakeRegistry *StakeRegistryFilterer) ParseOperatorWeightUpdated(log types.Log) (*StakeRegistryOperatorWeightUpdated, error) {
	event := new(StakeRegistryOperatorWeightUpdated)
	if err := _StakeRegistry.contract.UnpackLog(event, "OperatorWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakeRegistry contract.
type StakeRegistryOwnershipTransferredIterator struct {
	Event *StakeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryOwnershipTransferred)
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
		it.Event = new(StakeRegistryOwnershipTransferred)
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
func (it *StakeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the StakeRegistry contract.
type StakeRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeRegistry *StakeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryOwnershipTransferredIterator{contract: _StakeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeRegistry *StakeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryOwnershipTransferred)
				if err := _StakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakeRegistry *StakeRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*StakeRegistryOwnershipTransferred, error) {
	event := new(StakeRegistryOwnershipTransferred)
	if err := _StakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryQuorumUpdatedIterator is returned from FilterQuorumUpdated and is used to iterate over the raw logs and unpacked data for QuorumUpdated events raised by the StakeRegistry contract.
type StakeRegistryQuorumUpdatedIterator struct {
	Event *StakeRegistryQuorumUpdated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryQuorumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryQuorumUpdated)
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
		it.Event = new(StakeRegistryQuorumUpdated)
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
func (it *StakeRegistryQuorumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryQuorumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryQuorumUpdated represents a QuorumUpdated event raised by the StakeRegistry contract.
type StakeRegistryQuorumUpdated struct {
	Old Quorum
	New Quorum
	Raw types.Log // Blockchain specific contextual infos
}

// FilterQuorumUpdated is a free log retrieval operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_StakeRegistry *StakeRegistryFilterer) FilterQuorumUpdated(opts *bind.FilterOpts) (*StakeRegistryQuorumUpdatedIterator, error) {

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeRegistryQuorumUpdatedIterator{contract: _StakeRegistry.contract, event: "QuorumUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumUpdated is a free log subscription operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_StakeRegistry *StakeRegistryFilterer) WatchQuorumUpdated(opts *bind.WatchOpts, sink chan<- *StakeRegistryQuorumUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryQuorumUpdated)
				if err := _StakeRegistry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
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

// ParseQuorumUpdated is a log parse operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_StakeRegistry *StakeRegistryFilterer) ParseQuorumUpdated(log types.Log) (*StakeRegistryQuorumUpdated, error) {
	event := new(StakeRegistryQuorumUpdated)
	if err := _StakeRegistry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistrySigningKeyUpdateIterator is returned from FilterSigningKeyUpdate and is used to iterate over the raw logs and unpacked data for SigningKeyUpdate events raised by the StakeRegistry contract.
type StakeRegistrySigningKeyUpdateIterator struct {
	Event *StakeRegistrySigningKeyUpdate // Event containing the contract specifics and raw log

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
func (it *StakeRegistrySigningKeyUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistrySigningKeyUpdate)
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
		it.Event = new(StakeRegistrySigningKeyUpdate)
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
func (it *StakeRegistrySigningKeyUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistrySigningKeyUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistrySigningKeyUpdate represents a SigningKeyUpdate event raised by the StakeRegistry contract.
type StakeRegistrySigningKeyUpdate struct {
	Operator      common.Address
	UpdateBlock   *big.Int
	NewSigningKey common.Address
	OldSigningKey common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyUpdate is a free log retrieval operation binding the contract event 0xd061168252f441733658f09e4d8f5b2d998ed4ef24a2bbfd6ceca52ea1315002.
//
// Solidity: event SigningKeyUpdate(address indexed operator, uint256 indexed updateBlock, address indexed newSigningKey, address oldSigningKey)
func (_StakeRegistry *StakeRegistryFilterer) FilterSigningKeyUpdate(opts *bind.FilterOpts, operator []common.Address, updateBlock []*big.Int, newSigningKey []common.Address) (*StakeRegistrySigningKeyUpdateIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var updateBlockRule []interface{}
	for _, updateBlockItem := range updateBlock {
		updateBlockRule = append(updateBlockRule, updateBlockItem)
	}
	var newSigningKeyRule []interface{}
	for _, newSigningKeyItem := range newSigningKey {
		newSigningKeyRule = append(newSigningKeyRule, newSigningKeyItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "SigningKeyUpdate", operatorRule, updateBlockRule, newSigningKeyRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistrySigningKeyUpdateIterator{contract: _StakeRegistry.contract, event: "SigningKeyUpdate", logs: logs, sub: sub}, nil
}

// WatchSigningKeyUpdate is a free log subscription operation binding the contract event 0xd061168252f441733658f09e4d8f5b2d998ed4ef24a2bbfd6ceca52ea1315002.
//
// Solidity: event SigningKeyUpdate(address indexed operator, uint256 indexed updateBlock, address indexed newSigningKey, address oldSigningKey)
func (_StakeRegistry *StakeRegistryFilterer) WatchSigningKeyUpdate(opts *bind.WatchOpts, sink chan<- *StakeRegistrySigningKeyUpdate, operator []common.Address, updateBlock []*big.Int, newSigningKey []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var updateBlockRule []interface{}
	for _, updateBlockItem := range updateBlock {
		updateBlockRule = append(updateBlockRule, updateBlockItem)
	}
	var newSigningKeyRule []interface{}
	for _, newSigningKeyItem := range newSigningKey {
		newSigningKeyRule = append(newSigningKeyRule, newSigningKeyItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "SigningKeyUpdate", operatorRule, updateBlockRule, newSigningKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistrySigningKeyUpdate)
				if err := _StakeRegistry.contract.UnpackLog(event, "SigningKeyUpdate", log); err != nil {
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

// ParseSigningKeyUpdate is a log parse operation binding the contract event 0xd061168252f441733658f09e4d8f5b2d998ed4ef24a2bbfd6ceca52ea1315002.
//
// Solidity: event SigningKeyUpdate(address indexed operator, uint256 indexed updateBlock, address indexed newSigningKey, address oldSigningKey)
func (_StakeRegistry *StakeRegistryFilterer) ParseSigningKeyUpdate(log types.Log) (*StakeRegistrySigningKeyUpdate, error) {
	event := new(StakeRegistrySigningKeyUpdate)
	if err := _StakeRegistry.contract.UnpackLog(event, "SigningKeyUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryThresholdWeightUpdatedIterator is returned from FilterThresholdWeightUpdated and is used to iterate over the raw logs and unpacked data for ThresholdWeightUpdated events raised by the StakeRegistry contract.
type StakeRegistryThresholdWeightUpdatedIterator struct {
	Event *StakeRegistryThresholdWeightUpdated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryThresholdWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryThresholdWeightUpdated)
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
		it.Event = new(StakeRegistryThresholdWeightUpdated)
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
func (it *StakeRegistryThresholdWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryThresholdWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryThresholdWeightUpdated represents a ThresholdWeightUpdated event raised by the StakeRegistry contract.
type StakeRegistryThresholdWeightUpdated struct {
	ThresholdWeight *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterThresholdWeightUpdated is a free log retrieval operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_StakeRegistry *StakeRegistryFilterer) FilterThresholdWeightUpdated(opts *bind.FilterOpts) (*StakeRegistryThresholdWeightUpdatedIterator, error) {

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "ThresholdWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeRegistryThresholdWeightUpdatedIterator{contract: _StakeRegistry.contract, event: "ThresholdWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdWeightUpdated is a free log subscription operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_StakeRegistry *StakeRegistryFilterer) WatchThresholdWeightUpdated(opts *bind.WatchOpts, sink chan<- *StakeRegistryThresholdWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "ThresholdWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryThresholdWeightUpdated)
				if err := _StakeRegistry.contract.UnpackLog(event, "ThresholdWeightUpdated", log); err != nil {
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

// ParseThresholdWeightUpdated is a log parse operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_StakeRegistry *StakeRegistryFilterer) ParseThresholdWeightUpdated(log types.Log) (*StakeRegistryThresholdWeightUpdated, error) {
	event := new(StakeRegistryThresholdWeightUpdated)
	if err := _StakeRegistry.contract.UnpackLog(event, "ThresholdWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryTotalWeightUpdatedIterator is returned from FilterTotalWeightUpdated and is used to iterate over the raw logs and unpacked data for TotalWeightUpdated events raised by the StakeRegistry contract.
type StakeRegistryTotalWeightUpdatedIterator struct {
	Event *StakeRegistryTotalWeightUpdated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryTotalWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryTotalWeightUpdated)
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
		it.Event = new(StakeRegistryTotalWeightUpdated)
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
func (it *StakeRegistryTotalWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryTotalWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryTotalWeightUpdated represents a TotalWeightUpdated event raised by the StakeRegistry contract.
type StakeRegistryTotalWeightUpdated struct {
	OldTotalWeight *big.Int
	NewTotalWeight *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalWeightUpdated is a free log retrieval operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_StakeRegistry *StakeRegistryFilterer) FilterTotalWeightUpdated(opts *bind.FilterOpts) (*StakeRegistryTotalWeightUpdatedIterator, error) {

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "TotalWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeRegistryTotalWeightUpdatedIterator{contract: _StakeRegistry.contract, event: "TotalWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalWeightUpdated is a free log subscription operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_StakeRegistry *StakeRegistryFilterer) WatchTotalWeightUpdated(opts *bind.WatchOpts, sink chan<- *StakeRegistryTotalWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "TotalWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryTotalWeightUpdated)
				if err := _StakeRegistry.contract.UnpackLog(event, "TotalWeightUpdated", log); err != nil {
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

// ParseTotalWeightUpdated is a log parse operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_StakeRegistry *StakeRegistryFilterer) ParseTotalWeightUpdated(log types.Log) (*StakeRegistryTotalWeightUpdated, error) {
	event := new(StakeRegistryTotalWeightUpdated)
	if err := _StakeRegistry.contract.UnpackLog(event, "TotalWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryUpdateMinimumWeightIterator is returned from FilterUpdateMinimumWeight and is used to iterate over the raw logs and unpacked data for UpdateMinimumWeight events raised by the StakeRegistry contract.
type StakeRegistryUpdateMinimumWeightIterator struct {
	Event *StakeRegistryUpdateMinimumWeight // Event containing the contract specifics and raw log

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
func (it *StakeRegistryUpdateMinimumWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryUpdateMinimumWeight)
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
		it.Event = new(StakeRegistryUpdateMinimumWeight)
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
func (it *StakeRegistryUpdateMinimumWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryUpdateMinimumWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryUpdateMinimumWeight represents a UpdateMinimumWeight event raised by the StakeRegistry contract.
type StakeRegistryUpdateMinimumWeight struct {
	OldMinimumWeight *big.Int
	NewMinimumWeight *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumWeight is a free log retrieval operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_StakeRegistry *StakeRegistryFilterer) FilterUpdateMinimumWeight(opts *bind.FilterOpts) (*StakeRegistryUpdateMinimumWeightIterator, error) {

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "UpdateMinimumWeight")
	if err != nil {
		return nil, err
	}
	return &StakeRegistryUpdateMinimumWeightIterator{contract: _StakeRegistry.contract, event: "UpdateMinimumWeight", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumWeight is a free log subscription operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_StakeRegistry *StakeRegistryFilterer) WatchUpdateMinimumWeight(opts *bind.WatchOpts, sink chan<- *StakeRegistryUpdateMinimumWeight) (event.Subscription, error) {

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "UpdateMinimumWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryUpdateMinimumWeight)
				if err := _StakeRegistry.contract.UnpackLog(event, "UpdateMinimumWeight", log); err != nil {
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

// ParseUpdateMinimumWeight is a log parse operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_StakeRegistry *StakeRegistryFilterer) ParseUpdateMinimumWeight(log types.Log) (*StakeRegistryUpdateMinimumWeight, error) {
	event := new(StakeRegistryUpdateMinimumWeight)
	if err := _StakeRegistry.contract.UnpackLog(event, "UpdateMinimumWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
